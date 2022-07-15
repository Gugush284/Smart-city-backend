package store

import (
	modelBroadcast "Smart-city/internal/apiserver/model/broadcast"
	modelNews "Smart-city/internal/apiserver/model/news"
	modelTeams "Smart-city/internal/apiserver/model/teams"
	modeltimetable "Smart-city/internal/apiserver/model/timetable"
)

type Newsrepository interface {
	GetNews() ([]modelNews.News, error)
}

type Broadcastrepository interface {
	GetBroadcast() ([]modelBroadcast.Broadcast, error)
}

type Timetablerepository interface {
	Create(t *modeltimetable.Timetable) error
	GetTimetable(idUser int) ([]modeltimetable.Timetable, error)
}

type Teamsrepository interface {
	GetTeams(idUser int) ([]modelTeams.Team, error)
}

type Scoreboardrepository interface {
}
