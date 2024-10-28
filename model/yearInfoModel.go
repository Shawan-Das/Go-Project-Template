package model

type Yearinfo struct {
	Yearcode     int    `json:"yearcode"`
	Openingfield string `json:"openingfield"`
	Startdate    string `json:"startdate"`
	Enddate      string `json:"enddate"`
	Status       string `json:"status"`
	Tdstatus     int    `json:"tdstatus"`
}
