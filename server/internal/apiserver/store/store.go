package store

type Store interface {
	News() Newsrepository
	Broadcast() Broadcastrepository
}
