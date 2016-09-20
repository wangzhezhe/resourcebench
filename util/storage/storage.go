package storage

type Storage interface {
	UpLoad(path string, name string, blobsum string) error

	DownLoad(path string, name string, blobsum string) error //check the specific blobsum in
}
