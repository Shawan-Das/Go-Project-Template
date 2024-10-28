package model

type ServiceParty struct {
	Id            int    `json:"id"`
	PartyCode     int    `json:"party_code"`
	PartyName     string `json:"party_name"`
	PartyNid      string `json:"party_nid"`
	PartyPhone    string `json:"party_phone"`
	PartyEmail    string `json:"party_email"`
	PartyAddress  string `json:"party_address"`
}

type ServicePartyArchive struct {
	ServiceParty
	ChangeEvent string `json:"change_event"`
	ChangeUser  string `json:"change_user"`
	ChangeDate  string `json:"change_date"`
	ChangeTime  string `json:"change_time"`
	TrackId     int    `json:"track_id"`
}
