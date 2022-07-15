package sqlstore

import (
	"Smart-city/internal/apiserver/store"
)

type Scoreboardrepository struct {
	store *SqlStore
}

func (s *SqlStore) Scoreboard() store.Scoreboardrepository {
	if s.Scoreboardrepository != nil {
		return s.Scoreboardrepository
	}

	s.Scoreboardrepository = &Scoreboardrepository{
		store: s,
	}

	return s.Scoreboardrepository
}
