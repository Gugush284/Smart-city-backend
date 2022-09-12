package sqlstore

import (
	modelEvents "Smart-city/internal/apiserver/model/event"
	"Smart-city/internal/apiserver/store"
)

type Eventsrepository struct {
	store *SqlStore
}

func (s *SqlStore) Event() store.Eventsrepository {
	if s.Eventsrepository != nil {
		return s.Eventsrepository
	}

	s.Eventsrepository = &Eventsrepository{
		store: s,
	}

	return s.Eventsrepository
}

func (r *Eventsrepository) GetEvents() ([]modelEvents.Event, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	return nil, nil
}
