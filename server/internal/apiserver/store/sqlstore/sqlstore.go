package sqlstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // sql driver ...
)

// Store for clients ...
type SqlStore struct {
	DbURL               string
	Db                  *sql.DB
	Newsrepository      *Newsrepository
	Broadcastrepository *Broadcastrepository
}

// New Store ...
func New(URL string) *SqlStore {
	return &SqlStore{
		DbURL: URL,
	}
}

// Open db ...
func (s *SqlStore) Open() error {
	db, err := sql.Open("mysql", s.DbURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Db = db

	return nil
}

// Close connection ...
func (s *SqlStore) Close() {
	s.Db.Close()
}

func (s *SqlStore) CreateTables() error {
	if err := s.Open(); err != nil {
		return err
	}

	statement, err := s.Db.Prepare(`create table IF NOT EXISTS news (
		id        integer      not null PRIMARY KEY AUTO_INCREMENT,
		title           varchar(150) not null default '',
		author      	varchar(50)  not null default '',
		txt				Text		 not null,
		time			DATETIME	 not null,
		picURL			varchar(50)  
	)`)
	if err != nil {
		return err
	}
	statement.Exec()
	statement.Close()

	statement, err = s.Db.Prepare(`create table IF NOT EXISTS broadcast (
		id              integer      not null PRIMARY KEY AUTO_INCREMENT,
		name			varchar(50)	 not null default 'Noname',
		broadURL		varchar(50),
		picURL			varchar(50)  
	)`)
	if err != nil {
		return err
	}
	statement.Exec()
	statement.Close()

	return nil
}
