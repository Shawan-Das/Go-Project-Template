package model

type Compinfo struct {
	Id        int    `json:"id"`
	Compcode  int    `json:"compcode"`
	Compname  string `json:"compname"`
	Location  string `json:"location"`
	Contactno string `json:"contactno"`
	Barcode   int    `json:"barcode"`
}
