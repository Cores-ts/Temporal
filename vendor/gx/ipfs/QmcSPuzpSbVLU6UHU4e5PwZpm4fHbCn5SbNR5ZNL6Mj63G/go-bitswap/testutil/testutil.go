package testutil

import (
	"math/rand"

	cid "gx/ipfs/QmTbxNB1NwDesLmKTscr4udL2tVP7MaxvXnD1D9yX7g3PN/go-cid"
	peer "gx/ipfs/QmYVXrKrKHDC9FobgmcmshCDyWwdrfwfanNQN4oxJ9Fk3h/go-libp2p-peer"
	"gx/ipfs/QmYYLnAzR28nAQ4U5MFniLprnktu6eTFKibeNt96V21EZK/go-block-format"
	bsmsg "gx/ipfs/QmcSPuzpSbVLU6UHU4e5PwZpm4fHbCn5SbNR5ZNL6Mj63G/go-bitswap/message"
	"gx/ipfs/QmcSPuzpSbVLU6UHU4e5PwZpm4fHbCn5SbNR5ZNL6Mj63G/go-bitswap/wantlist"
	blocksutil "gx/ipfs/QmcbQ44AgbtV1rnxZz5RsRVduxdgNYLANxnsuW5wvvu4ts/go-ipfs-blocksutil"
)

var blockGenerator = blocksutil.NewBlockGenerator()
var prioritySeq int

// GenerateBlocksOfSize generates a series of blocks of the given byte size
func GenerateBlocksOfSize(n int, size int64) []blocks.Block {
	generatedBlocks := make([]blocks.Block, 0, n)
	buf := make([]byte, size)
	for i := 0; i < n; i++ {
		// rand.Read never errors
		rand.Read(buf)
		b := blocks.NewBlock(buf)
		generatedBlocks = append(generatedBlocks, b)

	}
	return generatedBlocks
}

// GenerateCids produces n content identifiers.
func GenerateCids(n int) []cid.Cid {
	cids := make([]cid.Cid, 0, n)
	for i := 0; i < n; i++ {
		c := blockGenerator.Next().Cid()
		cids = append(cids, c)
	}
	return cids
}

// GenerateWantlist makes a populated wantlist.
func GenerateWantlist(n int, ses uint64) *wantlist.SessionTrackedWantlist {
	wl := wantlist.NewSessionTrackedWantlist()
	for i := 0; i < n; i++ {
		prioritySeq++
		entry := wantlist.NewRefEntry(blockGenerator.Next().Cid(), prioritySeq)
		wl.AddEntry(entry, ses)
	}
	return wl
}

// GenerateMessageEntries makes fake bitswap message entries.
func GenerateMessageEntries(n int, isCancel bool) []bsmsg.Entry {
	bsmsgs := make([]bsmsg.Entry, 0, n)
	for i := 0; i < n; i++ {
		prioritySeq++
		msg := bsmsg.Entry{
			Entry:  wantlist.NewRefEntry(blockGenerator.Next().Cid(), prioritySeq),
			Cancel: isCancel,
		}
		bsmsgs = append(bsmsgs, msg)
	}
	return bsmsgs
}

var peerSeq int

// GeneratePeers creates n peer ids.
func GeneratePeers(n int) []peer.ID {
	peerIds := make([]peer.ID, 0, n)
	for i := 0; i < n; i++ {
		peerSeq++
		p := peer.ID(peerSeq)
		peerIds = append(peerIds, p)
	}
	return peerIds
}

var nextSession uint64

// GenerateSessionID make a unit session identifier.
func GenerateSessionID() uint64 {
	nextSession++
	return uint64(nextSession)
}

// ContainsPeer returns true if a peer is found n a list of peers.
func ContainsPeer(peers []peer.ID, p peer.ID) bool {
	for _, n := range peers {
		if p == n {
			return true
		}
	}
	return false
}

// IndexOf returns the index of a given cid in an array of blocks
func IndexOf(blks []blocks.Block, c cid.Cid) int {
	for i, n := range blks {
		if n.Cid() == c {
			return i
		}
	}
	return -1
}

// ContainsBlock returns true if a block is found n a list of blocks
func ContainsBlock(blks []blocks.Block, block blocks.Block) bool {
	return IndexOf(blks, block.Cid()) != -1
}
