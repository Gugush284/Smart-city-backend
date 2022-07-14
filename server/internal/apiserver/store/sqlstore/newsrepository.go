package sqlstore

import (
	modelNews "Smart-city/internal/apiserver/model/news"
	"Smart-city/internal/apiserver/store"
	"time"
)

type Newsrepository struct {
	store *SqlStore
}

func (s *SqlStore) News() store.Newsrepository {
	if s.Newsrepository != nil {
		return s.Newsrepository
	}

	s.Newsrepository = &Newsrepository{
		store: s,
	}

	return s.Newsrepository
}

func (r *Newsrepository) GetNews() ([]modelNews.News, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	data := []modelNews.News{}

	rows, err := r.store.Db.Query("select * from news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tstr string

	for rows.Next() {
		n := &modelNews.News{}
		err := rows.Scan(
			&n.Id,
			&n.Title,
			&n.Author,
			&n.Txt,
			&tstr,
			&n.PicURL,
		)
		if err != nil {
			return nil, err
		}

		t, err := time.Parse("2006-01-02T15:04:05", tstr)
		if err != nil {
			return nil, err
		}

		n.Time = t.Unix()

		data = append(data, *n)
	}

	return data, nil
}
