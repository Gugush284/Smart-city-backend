package sqlstore

import (
	modelEvents "Smart-city/internal/apiserver/model/event"
	"Smart-city/internal/apiserver/store"
	"encoding/json"
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

	data := []modelEvents.Event{}

	rows, err := r.store.Db.Query("select * from eventdetaildto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := &modelEvents.Event{}
		err := rows.Scan(
			&n.Id,
			&n.Title,
			&n.Description,
			&n.BeginTime,
			&n.EndTime,
			&n.Address,
			&n.Money,
			&n.CurParticCount,
			&n.TrgtParticCount,
			&n.EventType,
			&n.Picture,
			&n.Players,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, *n)
	}

	return data, nil
}

func (r *Eventsrepository) GetEvent(id string) (*modelEvents.Event, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	row := r.store.Db.QueryRow("select * from eventdetaildto where id = ?", id)

	var data modelEvents.Event

	if err := row.Scan(
		&data.Id,
		&data.Title,
		&data.Description,
		&data.BeginTime,
		&data.EndTime,
		&data.Address,
		&data.Money,
		&data.CurParticCount,
		&data.TrgtParticCount,
		&data.EventType,
		&data.Picture,
		&data.Players,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Eventsrepository) RegEvent(regevent *modelEvents.EventRegistratePLayers) (*modelEvents.Event, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	row := r.store.Db.QueryRow("select curCount from eventdetaildto where id = ?", regevent.Idevent)

	var old_cur int

	if err := row.Scan(&old_cur); err != nil {
		return nil, err
	}

	l := len(regevent.ChosenPlayers) + old_cur

	_, err := r.store.Db.Exec(
		"UPDATE eventdetaildto SET curCount = ? WHERE id = ?",
		l,
		regevent.Idevent,
	)
	if err != nil {
		return nil, err
	}

	row = r.store.Db.QueryRow("select * from eventdetaildto where id = ?", regevent.Idevent)

	var data modelEvents.Event
	var jsondata []byte

	if err := row.Scan(
		&data.Id,
		&data.Title,
		&data.Description,
		&data.BeginTime,
		&data.EndTime,
		&data.Address,
		&data.Money,
		&data.CurParticCount,
		&data.TrgtParticCount,
		&data.EventType,
		&data.Picture,
		&jsondata,
	); err != nil {
		return nil, err
	}

	jerr := json.Unmarshal(jsondata, &data.Players)
	if jerr != nil {
		return nil, err
	}

	data.Players = append(data.Players, regevent.ChosenPlayers...)

	jsondata, jerr = json.Marshal(data.Players)
	if jerr != nil {
		return nil, err
	}

	_, err = r.store.Db.Exec(
		"UPDATE eventdetaildto SET players = ? WHERE id = ?",
		jsondata,
		regevent.Idevent,
	)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
