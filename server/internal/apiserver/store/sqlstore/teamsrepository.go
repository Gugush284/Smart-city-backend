package sqlstore

import (
	modelTeams "Smart-city/internal/apiserver/model/teams"
	"Smart-city/internal/apiserver/store"
)

type Teamsrepository struct {
	store *SqlStore
}

func (s *SqlStore) Teams() store.Teamsrepository {
	if s.Teamsrepository != nil {
		return s.Teamsrepository
	}

	s.Teamsrepository = &Teamsrepository{
		store: s,
	}

	return s.Teamsrepository
}

func (r *Teamsrepository) GetTeams(idUser int) ([]modelTeams.Team, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	data := []modelTeams.Team{}

	rows, err := r.store.Db.Query("select * from teams where IdUser = ?", idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := &modelTeams.Team{}
		err := rows.Scan(
			&t.Id,
			&t.IdUser,
			&t.Name,
			&t.Sport,
			&t.Place,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, *t)
	}

	return data, nil
}
