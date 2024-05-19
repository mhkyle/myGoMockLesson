package db

type DB interface {
	Get(key string) (string, error)
}
