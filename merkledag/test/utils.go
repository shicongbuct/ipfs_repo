package mdutils

import (
	bsrv "github.com/ipfs/go-ipfs/blockservice"
	dag "github.com/ipfs/go-ipfs/merkledag"

	blockstore "gx/ipfs/QmRatnbGjPcoyzVjfixMZnuT1xQbjM7FgnL6FX4CKJeDE2/go-ipfs-blockstore"
	offline "gx/ipfs/QmShbyKV9P7QuFecDHXsgrQ4rxxm71MUkGVpwedT4VQ8Bf/go-ipfs-exchange-offline"
	ipld "github.com/ipfs/go-ipld-format"
	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
)

// Mock returns a new thread-safe, mock DAGService.
func Mock() ipld.DAGService {
	return dag.NewDAGService(Bserv())
}

// Bserv returns a new, thread-safe, mock BlockService.
func Bserv() bsrv.BlockService {
	bstore := blockstore.NewBlockstore(dssync.MutexWrap(ds.NewMapDatastore()))
	return bsrv.New(bstore, offline.Exchange(bstore))
}
