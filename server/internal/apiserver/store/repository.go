package store

import (
	modelBroadcast "Smart-city/internal/apiserver/model/broadcast"
	modelNews "Smart-city/internal/apiserver/model/news"
)

type Newsrepository interface {
	GetNews() ([]modelNews.News, error)
}

type Broadcastrepository interface {
	GetBroadcast() ([]modelBroadcast.Broadcast, error)
}
