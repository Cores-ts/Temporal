package v2

import (
	"net/http"
	"strconv"

	"github.com/RTradeLtd/Temporal/eh"
	"github.com/RTradeLtd/Temporal/utils"

	"github.com/RTradeLtd/Temporal/queue"
	"github.com/RTradeLtd/Temporal/rtfscluster"
	"github.com/gin-gonic/gin"
	gocid "github.com/ipfs/go-cid"
)

// PinHashToCluster is used to trigger a cluster pin of a particular CID
func (api *API) pinHashToCluster(c *gin.Context) {
	username, err := GetAuthenticatedUserFromContext(c)
	if err != nil {
		api.LogError(err, eh.NoAPITokenError)(c, http.StatusBadRequest)
		return
	}
	hash := c.Param("hash")
	if _, err := gocid.Decode(hash); err != nil {
		Fail(c, err)
		return
	}
	forms := api.extractPostForms(c, "hold_time")
	if len(forms) == 0 {
		return
	}
	holdTimeInt, err := strconv.ParseInt(forms["hold_time"], 10, 64)
	if err != nil {
		Fail(c, err)
		return
	}
	cost, err := utils.CalculatePinCost(hash, holdTimeInt, api.ipfs, false)
	if err != nil {
		api.LogError(err, eh.CallCostCalculationError)(c, http.StatusBadRequest)
		return
	}
	if err := api.validateUserCredits(username, cost); err != nil {
		api.LogError(err, eh.InvalidBalanceError)(c, http.StatusPaymentRequired)
		return
	}
	ipfsClusterPin := queue.IPFSClusterPin{
		CID:              hash,
		NetworkName:      "public",
		UserName:         username,
		HoldTimeInMonths: holdTimeInt,
		CreditCost:       cost,
	}
	if err = api.queues.cluster.PublishMessage(ipfsClusterPin); err != nil {
		api.LogError(err, eh.QueuePublishError)(c, http.StatusBadRequest)
		api.refundUserCredits(username, "cluster-pin", cost)
		return
	}
	api.LogWithUser(username).Info("cluster pin request sent to backend")
	Respond(c, http.StatusOK, gin.H{"response": "cluster pin request sent to backend"})
}

// SyncClusterErrorsLocally is used to parse through the local cluster state and sync any errors that are detected.
func (api *API) syncClusterErrorsLocally(c *gin.Context) {
	username, err := GetAuthenticatedUserFromContext(c)
	if err != nil {
		api.LogError(err, eh.NoAPITokenError)(c, http.StatusBadRequest)
		return
	}
	if err := api.validateAdminRequest(username); err != nil {
		FailNotAuthorized(c, eh.UnAuthorizedAdminAccess)
		return
	}
	// initialize a connection to the cluster
	manager, err := rtfscluster.Initialize("", "")
	if err != nil {
		api.LogError(err, eh.IPFSConnectionError)(c)
		return
	}
	// parse the local cluster status, and sync any errors, retunring the cids that were in an error state
	syncedCids, err := manager.ParseLocalStatusAllAndSync()
	if err != nil {
		api.LogError(err, eh.IPFSClusterStatusError)(c)
		return
	}

	api.LogWithUser(username).Info("local cluster errors parsed")
	Respond(c, http.StatusOK, gin.H{"response": syncedCids})
}

// GetLocalStatusForClusterPin is used to get the localnode's cluster status for a particular pin
func (api *API) getLocalStatusForClusterPin(c *gin.Context) {
	username, err := GetAuthenticatedUserFromContext(c)
	if err != nil {
		api.LogError(err, eh.NoAPITokenError)(c, http.StatusBadRequest)
		return
	}
	if err := api.validateAdminRequest(username); err != nil {
		FailNotAuthorized(c, eh.UnAuthorizedAdminAccess)
		return
	}
	hash := c.Param("hash")
	if _, err := gocid.Decode(hash); err != nil {
		Fail(c, err)
		return
	}
	// initialize a connection to the cluster
	manager, err := rtfscluster.Initialize("", "")
	if err != nil {
		api.LogError(err, eh.IPFSClusterConnectionError)(c)
		return
	}
	// get the cluster status for the cid only asking the local cluster node
	status, err := manager.GetStatusForCidLocally(hash)
	if err != nil {
		api.LogError(err, eh.IPFSClusterStatusError)(c)
		return
	}

	api.LogWithUser(username).Info("local cluster status for pin requested")

	Respond(c, http.StatusOK, gin.H{"response": status})
}

// GetGlobalStatusForClusterPin is used to get the global cluster status for a particular pin
func (api *API) getGlobalStatusForClusterPin(c *gin.Context) {
	username, err := GetAuthenticatedUserFromContext(c)
	if err != nil {
		api.LogError(err, eh.NoAPITokenError)(c, http.StatusBadRequest)
		return
	}
	if err := api.validateAdminRequest(username); err != nil {
		FailNotAuthorized(c, eh.UnAuthorizedAdminAccess)
		return
	}
	hash := c.Param("hash")
	if _, err := gocid.Decode(hash); err != nil {
		Fail(c, err)
		return
	}
	// initialize a connection to the cluster
	manager, err := rtfscluster.Initialize("", "")
	if err != nil {
		api.LogError(err, eh.IPFSClusterConnectionError)(c)
		return
	}
	// get the cluster wide status for this particular pin
	status, err := manager.GetStatusForCidGlobally(hash)
	if err != nil {
		api.LogError(err, eh.IPFSClusterStatusError)(c)
		return
	}

	api.LogWithUser(username).Info("global cluster status for pin requested")

	Respond(c, http.StatusOK, gin.H{"response": status})
}

// FetchLocalClusterStatus is used to fetch the status of the localhost's cluster state, and not the rest of the cluster
func (api *API) fetchLocalClusterStatus(c *gin.Context) {
	username, err := GetAuthenticatedUserFromContext(c)
	if err != nil {
		api.LogError(err, eh.NoAPITokenError)(c, http.StatusBadRequest)
		return
	}
	if err := api.validateAdminRequest(username); err != nil {
		FailNotAuthorized(c, eh.UnAuthorizedAdminAccess)
		return
	}
	// this will hold all the retrieved content hashes
	var cids []gocid.Cid
	// this will hold all the statuses of the content hashes
	var statuses []string
	// initialize a connection to the cluster
	manager, err := rtfscluster.Initialize("", "")
	if err != nil {
		api.LogError(err, eh.IPFSClusterConnectionError)(c)
		return
	}
	// fetch a map of all the statuses
	maps, err := manager.FetchLocalStatus()
	if err != nil {
		api.LogError(err, eh.IPFSClusterStatusError)(c)
		return
	}
	// parse the maps
	for k, v := range maps {
		cids = append(cids, k)
		statuses = append(statuses, v)
	}

	api.LogWithUser(username).Info("local cluster state fetched")
	Respond(c, http.StatusOK, gin.H{"response": gin.H{"cids": cids, "statuses": statuses}})
}