package blockchainstore

import "github.com/gouniverse/dataobject"

type Blockchain struct {
	dataobject.DataObject
}

func NewBlockchain() *Blockchain {
	return &Blockchain{}
}
