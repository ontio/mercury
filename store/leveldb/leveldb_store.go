package store

import (
	"git.ont.io/ontid/otf/store"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

// Provider leveldb implementation of storage.Provider interface
type Provider struct {
	dbPath  string
	dbStore *levelDBStore
}

func NewProvider(dbPath string) *Provider {
	return &Provider{
		dbPath:  dbPath,
		dbStore: &levelDBStore{},
	}
}

type levelDBStore struct {
	db    *leveldb.DB
	batch *leveldb.Batch
}

func (p *Provider) OpenStore(path string) (store.Store, error) {
	return p.newLevelDBStore(path)
}

func (p *Provider) newLevelDBStore(path string) (*levelDBStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(path, nil)
	}
	if err != nil {
		return nil, err
	}
	return &levelDBStore{
		db:    db,
		batch: nil,
	}, nil
}

func (p *Provider) Close() error {
	return p.dbStore.Close()
}

//Put a key-value pair to leveldb
func (self *levelDBStore) Put(key []byte, value []byte) error {
	return self.db.Put(key, value, nil)
}

//Get the value of a key from leveldb
func (self *levelDBStore) Get(key []byte) ([]byte, error) {
	dat, err := self.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

//Has return whether the key is exist in leveldb
func (self *levelDBStore) Has(key []byte) (bool, error) {
	return self.db.Has(key, nil)
}

//Delete the the in leveldb
func (self *levelDBStore) Delete(key []byte) error {
	return self.db.Delete(key, nil)
}

//Close leveldb
func (self *levelDBStore) Close() error {
	err := self.db.Close()
	return err
}

//NewBatch start commit batch
func (self *levelDBStore) NewBatch() {
	self.batch = new(leveldb.Batch)
}

//BatchPut put a key-value pair to leveldb batch
func (self *levelDBStore) BatchPut(key []byte, value []byte) {
	self.batch.Put(key, value)
}

//BatchDelete delete a key to leveldb batch
func (self *levelDBStore) BatchDelete(key []byte) {
	self.batch.Delete(key)
}

//BatchCommit commit batch to leveldb
func (self *levelDBStore) BatchCommit() error {
	err := self.db.Write(self.batch, nil)
	if err != nil {
		return err
	}
	self.batch = nil
	return nil
}
