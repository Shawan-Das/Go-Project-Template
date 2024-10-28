package dto

type CreateTokenOutput struct {
	Message      string      `json:"message"`
	IsSuccess    bool        `json:"isSuccess"`
	Token        string      `json:"token,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	Username     interface{} `json:"username,omitempty"`
}
