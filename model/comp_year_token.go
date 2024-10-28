package model

type Comp_year_token struct {
	Username string `json:"username"`
	Geninfo  string `json:"geninfo"`
	Id       int    `json:"id"`
}

type Comp_year_token_dto struct {
	Username string `json:"username"`
	Geninfo  string `json:"geninfo"`
}

type Comp_year_token_input_dto struct {
	User_id int    `json:"user_id"`
	Geninfo string `json:"geninfo"`
}
