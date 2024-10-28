package dto

import "github.com/tools/iservice/model"

type PartyCode struct {
	PartyCode int `json:"party_code"`
}

type UpdateParty struct {
	model.ServiceParty
	ChangeUser string `json:"change_user"`
}

type DeleteParty struct {
	PartyCode  int    `json:"party_code"`
	ChangeUser string `json:"change_user"`
}
