package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/RTradeLtd/Temporal/mini"
	"github.com/RTradeLtd/kaas"
	"github.com/RTradeLtd/rtfs"
	log "github.com/sirupsen/logrus"

	"github.com/RTradeLtd/config"

	"github.com/RTradeLtd/database/models"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"

	pb "github.com/RTradeLtd/grpc/krab"
	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

// ProcessIPFSKeyCreation is used to create IPFS keys
func (qm *Manager) ProcessIPFSKeyCreation(ctx context.Context, wg *sync.WaitGroup, msgs <-chan amqp.Delivery, db *gorm.DB, cfg *config.TemporalConfig) error {
	kb, err := kaas.NewClient(cfg.Endpoints)
	if err != nil {
		return err
	}
	userManager := models.NewUserManager(db)
	qm.LogInfo("processing ipfs key creation requests")
	for {
		select {
		case d := <-msgs:
			wg.Add(1)
			go func(d amqp.Delivery) {
				defer wg.Done()
				qm.LogInfo("new message detected")
				key := IPFSKeyCreation{}
				if err = json.Unmarshal(d.Body, &key); err != nil {
					qm.LogError(err, "failed to unmarshal message")
					d.Ack(false)
					return
				}
				var (
					keyTypeInt int
					bitsInt    int
				)
				// validate the key parameters used for creation
				switch key.Type {
				case "rsa":
					keyTypeInt = ci.RSA
					// ensure the provided key size is within a valid range, otherwise default to 2048
					if key.Size > 4096 || key.Size < 2048 {
						bitsInt = 2048
					} else {
						bitsInt = key.Size
					}
				case "ed25519":
					keyTypeInt = ci.Ed25519
					// ed25519 keys use 256 bits, so regardless of what the user provides for bit size, hard set 256
					bitsInt = 256
				default:
					qm.refundCredits(key.UserName, "key", key.CreditCost, db)
					qm.LogError(err, "invalid key type must be ed25519 or rsa")
					d.Ack(false)
					return
				}
				// generate the appropriate keypair
				pk, _, err := ci.GenerateKeyPair(keyTypeInt, bitsInt)
				if err != nil {
					qm.refundCredits(key.UserName, "key", key.CreditCost, db)
					qm.LogError(err, "failed to create key")
					d.Ack(false)
					return
				}
				// retrieve a human friendly format, also verifying the key is a valid ipfs key
				id, err := peer.IDFromPrivateKey(pk)
				if err != nil {
					qm.refundCredits(key.UserName, "key", key.CreditCost, db)
					qm.LogError(err, "failed to get id from private key")
					d.Ack(false)
					return
				}
				// convert the key to bytes to send to krab for processing
				pkBytes, err := pk.Bytes()
				if err != nil {
					qm.refundCredits(key.UserName, "key", key.CreditCost, db)
					qm.LogError(err, "failed to create key")
					d.Ack(false)
					return
				}

				// to prevent key name collision, we prefix the keyname by a hyphen, and the username
				// key name of key2 for user testuser becomes testuser-key2
				keyName := fmt.Sprintf("%s-%s", key.UserName, key.Name)
				// store the key in krab
				if _, err := kb.PutPrivateKey(context.Background(), &pb.KeyPut{Name: keyName, PrivateKey: pkBytes}); err != nil {
					qm.refundCredits(key.UserName, "key", key.CreditCost, db)
					qm.LogError(err, "failed to create key")
					d.Ack(false)
					return
				}
				// doesn't need a refund, key was generated and stored in our keystore, but information not saved to db
				if err := userManager.AddIPFSKeyForUser(key.UserName, keyName, id.Pretty()); err != nil {
					qm.LogError(err, "failed to add ipfs key to database")
					d.Ack(false)
					return
				}
				qm.LogInfo("successfully processed ipfs key creation request")
				d.Ack(false)
				return // we must return here in order to trigger the wg.Done() defer
			}(d)
		case <-ctx.Done():
			qm.Close()
			wg.Done()
			return nil
		}
	}
}

// ProccessIPFSPins is used to process IPFS pin requests
func (qm *Manager) ProccessIPFSPins(ctx context.Context, wg *sync.WaitGroup, msgs <-chan amqp.Delivery, db *gorm.DB, cfg *config.TemporalConfig) error {
	userManager := models.NewUserManager(db)
	networkManager := models.NewHostedIPFSNetworkManager(db)
	uploadManager := models.NewUploadManager(db)
	// initialize a connection to the cluster pin queue so we can trigger pinning of this content to our cluster
	qmCluster, err := Initialize(IpfsClusterPinQueue, cfg.RabbitMQ.URL, true, false)
	if err != nil {
		qm.LogError(err, "failed to initialize cluster pin queue connection")
		return err
	}
	qm.LogInfo("processing ipfs pins")
	for {
		select {
		case d := <-msgs:
			wg.Add(1)
			go func(d amqp.Delivery) {
				defer wg.Done()
				qm.LogInfo("new message detected")
				pin := &IPFSPin{}
				if err := json.Unmarshal(d.Body, pin); err != nil {
					qm.LogError(err, "failed to unmarshal message")
					d.Ack(false)
					return
				}
				// setup the default api connection
				apiURL := cfg.IPFS.APIConnection.Host + ":" + cfg.IPFS.APIConnection.Port
				// check whether or not this pin is for a private network
				// if it is, verify whether the user has acess to the network, and retrieve the api url
				if pin.NetworkName != "public" {
					canAccess, err := userManager.CheckIfUserHasAccessToNetwork(pin.UserName, pin.NetworkName)
					if err != nil {
						qm.refundCredits(pin.UserName, "pin", pin.CreditCost, db)
						qm.LogError(err, "failed to lookup private network in database")
						d.Ack(false)
						return
					}
					if !canAccess {
						qm.refundCredits(pin.UserName, "pin", pin.CreditCost, db)
						qm.LogError(errors.New("user does not have access to private network"), "invalid private network access")
						d.Ack(false)
						return
					}
					apiURL, err = networkManager.GetAPIURLByName(pin.NetworkName)
					if err != nil {
						qm.refundCredits(pin.UserName, "pin", pin.CreditCost, db)
						qm.LogError(err, "failed to search for api url")
						d.Ack(false)
						return
					}
				}
				qm.LogInfo("initializing connection to ipfs")
				// connect to ipfs
				ipfsManager, err := rtfs.NewManager(apiURL, nil, time.Minute*10)
				if err != nil {
					qm.refundCredits(pin.UserName, "pin", pin.CreditCost, db)
					qm.LogError(err, "failed to initialize connection to ipfs")
					d.Ack(false)
					return
				}
				qm.LogInfo("pinning hash to ipfs")
				// pin the content
				if err = ipfsManager.Pin(pin.CID); err != nil {
					qm.refundCredits(pin.UserName, "pin", pin.CreditCost, db)
					qm.LogError(err, "failed to pin hash to ipfs")
					d.Ack(false)
					return
				}
				// cluster support for private networks isn't available yet
				// as such, skip additional processing for cluster pins
				if pin.NetworkName != "public" {
					qm.LogInfo("successfully proccessed private network pin")
					d.Ack(false)
					return
				}
				qm.LogInfo("successfully pinned hash to ipfs")
				clusterAddMsg := IPFSClusterPin{
					CID:              pin.CID,
					NetworkName:      pin.NetworkName,
					HoldTimeInMonths: pin.HoldTimeInMonths,
					UserName:         pin.UserName,
				}
				// do not perform credit handling, as the content is already pinned
				clusterAddMsg.CreditCost = 0
				if err = qmCluster.PublishMessage(clusterAddMsg); err != nil {
					qm.LogError(err, "failed to publish cluster pin message to rabbitmq")
					d.Ack(false)
					return
				}
				upload, err := uploadManager.FindUploadByHashAndNetwork(pin.CID, pin.NetworkName)
				if err != nil && err != gorm.ErrRecordNotFound {
					qm.LogError(err, "failed to find upload in database")
					d.Ack(false)
					return
				}
				// check whether or not we have seen this content hash before to determine how database needs to be updated
				if upload == nil {
					_, err = uploadManager.NewUpload(pin.CID, "pin", models.UploadOptions{
						NetworkName:      pin.NetworkName,
						Username:         pin.UserName,
						HoldTimeInMonths: pin.HoldTimeInMonths})
				} else {
					// the record already exists so we will update
					_, err = uploadManager.UpdateUpload(pin.HoldTimeInMonths, pin.UserName, pin.CID, pin.NetworkName)
				}
				// validate whether or not the database was updated properly
				if err != nil {
					qm.LogError(err, "failed to update upload in data, but cluster pin successfuly sent")
				} else {
					qm.LogInfo("successfully processed pin request")
				}
				d.Ack(false)
				return // we must return here in order to trigger the wg.Done() defer
			}(d)
		case <-ctx.Done():
			qm.Close()
			wg.Done()
			return nil
		}
	}
}

// ProccessIPFSFiles is used to process messages sent to rabbitmq to upload files to IPFS.
// This queue is invoked by advanced file upload requests
func (qm *Manager) ProccessIPFSFiles(ctx context.Context, wg *sync.WaitGroup, msgs <-chan amqp.Delivery, cfg *config.TemporalConfig, db *gorm.DB) error {
	ipfsManager, err := rtfs.NewManager(cfg.IPFS.APIConnection.Host+":"+cfg.IPFS.APIConnection.Port, nil, time.Minute*10)
	if err != nil {
		qm.LogError(err, "failed to initialize connection to ipfs")
		return err
	}
	// initialize connection to pin queues
	qmPin, err := Initialize(IpfsPinQueue, cfg.RabbitMQ.URL, true, false)
	if err != nil {
		qm.LogError(err, "failed to initialize pin queue connection")
		return err
	}
	ue := models.NewEncryptedUploadManager(db)
	userManager := models.NewUserManager(db)
	networkManager := models.NewHostedIPFSNetworkManager(db)
	qm.LogInfo("processing ipfs files")
	for {
		select {
		case d := <-msgs:
			wg.Add(1)
			go func(d amqp.Delivery) {
				defer wg.Done()
				qm.LogInfo("new message detected")
				ipfsFile := IPFSFile{}
				// unmarshal the messagee
				err = json.Unmarshal(d.Body, &ipfsFile)
				if err != nil {
					qm.LogError(err, "failed to unmarshal message")
					d.Ack(false)
					return
				}
				// get the connection url for the minio host temporarily storing this file
				endpoint := fmt.Sprintf("%s:%s", ipfsFile.MinioHostIP, cfg.MINIO.Connection.Port)
				// grab our credentials for minio
				accessKey := cfg.MINIO.AccessKey
				secretKey := cfg.MINIO.SecretKey
				// setup our connection to minio
				minioManager, err := mini.NewMinioManager(endpoint, accessKey, secretKey, false)
				if err != nil {
					qm.LogError(err, "failed to initialize connection to minio")
					qm.refundCredits(ipfsFile.UserName, "file", ipfsFile.CreditCost, db)
					d.Ack(false)
					return
				}
				// determine whether or not the request is for a private network
				if ipfsFile.NetworkName != "public" {
					canAccess, err := userManager.CheckIfUserHasAccessToNetwork(ipfsFile.UserName, ipfsFile.NetworkName)
					if err != nil {
						qm.LogError(err, "failed to check database for user network access", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
						qm.refundCredits(ipfsFile.UserName, "file", ipfsFile.CreditCost, db)
						d.Ack(false)
						return
					}
					if !canAccess {
						qm.LogError(err, "unauthorized access to private network", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
						qm.refundCredits(ipfsFile.UserName, "file", ipfsFile.CreditCost, db)
						d.Ack(false)
						return
					}
					apiURLName, err := networkManager.GetAPIURLByName(ipfsFile.NetworkName)
					if err != nil {
						qm.LogError(err, "failed to look for api url by name", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
						qm.refundCredits(ipfsFile.UserName, "file", ipfsFile.CreditCost, db)
						d.Ack(false)
						return
					}
					ipfsManager, err = rtfs.NewManager(apiURLName, nil, time.Minute*10)
					if err != nil {
						qm.LogError(err, "failed to initialize connection to private ifps network", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
						qm.refundCredits(ipfsFile.UserName, "file", ipfsFile.CreditCost, db)
						d.Ack(false)
						return
					}
				}
				qm.LogInfo("retrieving object from minio")
				obj, err := minioManager.GetObject(ipfsFile.ObjectName, mini.GetObjectOptions{
					Bucket: ipfsFile.BucketName,
				})
				if err != nil {
					qm.LogError(err, "failed to get object from minio", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
					qm.refundCredits(ipfsFile.UserName, "file", ipfsFile.CreditCost, db)
					d.Ack(false)
					return
				}
				qm.LogInfo("successfully retrieved object from minio, adding file to ipfs")
				resp, err := ipfsManager.Add(obj)
				if err != nil {
					qm.LogError(err, "failed to add file to ipfs", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
					qm.refundCredits(ipfsFile.UserName, "file", ipfsFile.CreditCost, db)
					d.Ack(false)
					return
				}
				qm.Logger.WithFields(log.Fields{
					"service": qm.Service,
					"user":    ipfsFile.UserName,
					"network": ipfsFile.NetworkName,
				}).Info("file added to ipfs")
				holdTimeInt, err := strconv.ParseInt(ipfsFile.HoldTimeInMonths, 10, 64)
				if err != nil {
					qm.LogError(err, "failed to parse string to int, using default of 1 month")
					holdTimeInt = 1
				}
				// we don't need to do any credit handling, as it has been done already
				pin := IPFSPin{
					CID:              resp,
					NetworkName:      ipfsFile.NetworkName,
					UserName:         ipfsFile.UserName,
					HoldTimeInMonths: holdTimeInt,
					CreditCost:       0,
				}
				if err = qmPin.PublishMessageWithExchange(pin, PinExchange); err != nil {
					qm.LogError(err, "failed to publish pin request to queue", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
					d.Ack(false)
					return
				}
				// if encrypted upload, do some special processing
				if ipfsFile.Encrypted {
					if _, err = ue.NewUpload(
						ipfsFile.UserName,
						ipfsFile.FileName,
						ipfsFile.NetworkName,
						resp,
					); err != nil {
						qm.LogError(err, "failed to update database with encrypted upload", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
						d.Ack(false)
						return
					}
				}
				qm.LogInfo("removing object from minio")
				err = minioManager.RemoveObject(ipfsFile.BucketName, ipfsFile.ObjectName)
				if err != nil {
					qm.LogError(err, "failed to remove object from minio", []interface{}{"user", ipfsFile.UserName, "network", ipfsFile.NetworkName})
					d.Ack(false)
					return
				}
				qm.LogInfo("successfully processed file upload")
				d.Ack(false)
				return // we must return here in order to trigger the wg.Done() defer
			}(d)
		case <-ctx.Done():
			qm.Close()
			wg.Done()
			return nil
		}
	}
}
