package sqlstore

import (
	modelBroadcast "Smart-city/internal/apiserver/model/broadcast"
	"Smart-city/internal/apiserver/store"
)

type Broadcastrepository struct {
	store *SqlStore
}

func (s *SqlStore) Broadcast() store.Broadcastrepository {
	if s.Broadcastrepository != nil {
		return s.Broadcastrepository
	}

	s.Broadcastrepository = &Broadcastrepository{
		store: s,
	}

	return s.Broadcastrepository
}

func (r *Broadcastrepository) GetBroadcast() ([]modelBroadcast.Broadcast, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	data := []modelBroadcast.Broadcast{}

	rows, err := r.store.Db.Query("select * from broadcast")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := &modelBroadcast.Broadcast{}
		err := rows.Scan(
			&n.Id,
			&n.Name,
			&n.BroadURL,
			&n.PicURL,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, *n)
	}

	return data, nil
}
