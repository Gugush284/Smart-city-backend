package sqlstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // sql driver ...
)

// Store for clients ...
type SqlStore struct {
	DbURL                string
	Db                   *sql.DB
	Newsrepository       *Newsrepository
	Broadcastrepository  *Broadcastrepository
	Timetablerepository  *Timetablerepository
	Teamsrepository      *Teamsrepository
	Scoreboardrepository *Scoreboardrepository
	Eventsrepository     *Eventsrepository
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

func createTable(s *SqlStore, req string) error {
	statement, err := s.Db.Prepare(req)
	if err != nil {
		return err
	}
	statement.Exec()
	statement.Close()

	return nil
}

func (s *SqlStore) CreateTables() error {
	if err := s.Open(); err != nil {
		return err
	}

	if err := createTable(s, `create table IF NOT EXISTS broadcast (
		id              integer      not null PRIMARY KEY AUTO_INCREMENT,
		name			varchar(50)	 not null default 'Noname',
		broadURL		varchar(50),
		picURL			varchar(50)  
	)`); err != nil {
		return err
	}

	if err := createTable(s, `create table IF NOT EXISTS eventdetaildto (
		id              integer      not null PRIMARY KEY AUTO_INCREMENT,
		title			varchar(50)	 not null default 'Noname',
		description		TEXT,
		beginTime		varchar(50),
		endTime         varchar(50),
		address         varchar(100),
		money           integer,
		curCount        integer,
		targetCount     integer,
		eventType       varchar(50),
		picture         varchar(100)
	)`); err != nil {
		return err
	}

	if err := createTable(s, `create table IF NOT EXISTS news (
		id              integer      not null PRIMARY KEY AUTO_INCREMENT,
		title			varchar(50)	 not null default 'Notitle',
		author			varchar(50)  not null default 'Noauthor',
		txt				TEXT,
		time			varchar(50)  not null,
		picURL			varchar(50)  
	)`); err != nil {
		return err
	}

	if err := createTable(s, `create table IF NOT EXISTS timetable (
		id              integer      not null PRIMARY KEY AUTO_INCREMENT,
		id_user			integer	 	 not null,
		txt				TEXT,
		time			int          not null,
		title			varchar(50)	 not null default 'Notitle'
	)`); err != nil {
		return err
	}

	if err := createTable(s, `create table IF NOT EXISTS teams (
		id              integer      not null PRIMARY KEY AUTO_INCREMENT,
		IdUser			varchar(50)  not null,
		Name			varchar(50),
		sport			varchar(50),
		place			varchar(50)
	)`); err != nil {
		return err
	}

	return nil
}
