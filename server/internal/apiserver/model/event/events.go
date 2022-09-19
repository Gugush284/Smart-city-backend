package modelEvents

import modelPlayerentity "Smart-city/internal/apiserver/model/PlayerEntity"

type Event struct {
	Id              int                              `json:"id"`
	Title           string                           `json:"title"`
	Description     string                           `json:"description"`
	BeginTime       string                           `json:"beginTime"`
	EndTime         string                           `json:"endTime"`
	Address         string                           `json:"address"`
	Money           int                              `json:"money"`
	CurParticCount  int                              `json:"currentParticipationCount"`
	TrgtParticCount int                              `json:"targetParticipationCount"`
	EventType       string                           `json:"eventType"`
	Picture         string                           `json:"picture"`
	Players         []modelPlayerentity.Playerentity `json:"Players"`
}

type EventRegistratePLayers struct {
	Idevent       int `json:"idevent"`
	ChosenPlayers []modelPlayerentity.Playerentity
}
