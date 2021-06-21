package store

// Store 存储接口
type Store interface {
	Set(k, v []byte) error
	Get(k []byte) ([]byte, error)
	Delete(k []byte) error
	Has(k []byte) (bool, error)
	ForEach(fn func(k, v []byte) error) error
	Close() error
	WALName() string
}

func OpenStore(path string) (Store, error) {
	return OpenLeveldb(path)
}
