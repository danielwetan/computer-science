package store

type Store interface {
	// General
	DBType() string
	DBVersion() string
}
