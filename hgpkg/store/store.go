package store

type Data interface {
}

type Store interface {
	Write(*Data)
	Load() Data
}
