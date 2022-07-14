package sqlstore

import (
	modeltimetable "Smart-city/internal/apiserver/model/timetable"
	"Smart-city/internal/apiserver/store"
)

type Timetablerepository struct {
	store *SqlStore
}

func (s *SqlStore) Timetable() store.Timetablerepository {
	if s.Timetablerepository != nil {
		return s.Timetablerepository
	}

	s.Timetablerepository = &Timetablerepository{
		store: s,
	}

	return s.Timetablerepository
}

func (r *Timetablerepository) Create(t *modeltimetable.Timetable) error {
	if err := t.Validate(); err != nil {
		return err
	}

	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return err
		}
	}

	statement, err := r.store.Db.Exec(
		"INSERT INTO timetable (id_user, txt, time, title) VALUES (?, ?, ?, ?)",
		t.IdUser,
		t.Txt,
		t.Time,
		t.Title,
	)
	if err != nil {
		return err
	}

	id, err := statement.LastInsertId()
	if err != nil {
		return err
	}
	if id == 0 {
		return nil
	}

	t.Id = int(id)

	return nil
}

func (r *Timetablerepository) GetTimetable(idUser int) ([]modeltimetable.Timetable, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	data := []modeltimetable.Timetable{}

	rows, err := r.store.Db.Query("select * from timetable where id_user = ?", idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := &modeltimetable.Timetable{}
		err := rows.Scan(
			&t.Id,
			&t.IdUser,
			&t.Txt,
			&t.Time,
			&t.Title,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, *t)
	}

	return data, nil
}
