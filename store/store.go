package store

// Provider storage provider interface
type Provider interface {
	// OpenStore opens a store with given name space and returns the handle
	OpenStore(name string) (Store, error)
	// Close closes all stores created under this store provider
	Close() error
}

type Store interface {
	// Put stores the key and the record
	Put(k []byte, v []byte) error
	// Get fetches the record based on key
	Get(k []byte) ([]byte, error)
	//check the record exist base on key
	Has(k []byte) (bool,error)
}
