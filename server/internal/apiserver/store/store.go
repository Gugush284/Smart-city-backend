package store

type Store interface {
	News() Newsrepository
	Broadcast() Broadcastrepository
	Timetable() Timetablerepository
	Teams() Teamsrepository
	Scoreboard() Scoreboardrepository
}
