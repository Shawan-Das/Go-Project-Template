package model

type Openbalance struct {
    Accid       int     `json:"accid"`
    Openbalance float32 `json:"openbalance"`
    Headname    string  `json:"headname"`
    Parent      int     `json:"parent"`
    Compid      int     `json:"compid"`
    Yearid      int     `json:"yearid"`
}

type OpenbalanceArchive struct {
	Openbalance
	ChangeUser string `json:"change_user"`
	ChangeDate string `json:"change_date"`
	ChangeTime string `json:"change_time"`
	TrackId    int    `json:"track_id"`
}