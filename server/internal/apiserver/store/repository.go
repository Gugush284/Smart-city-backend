package store

import modelNews "Smart-city/internal/apiserver/model/news"

type Newsrepository interface {
	GetNews() ([]modelNews.News, error)
}
