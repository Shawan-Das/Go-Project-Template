package dto

type ValidateTokenOutput struct {
	Message      string `json:"message"`
	IsSuccess    bool   `json:"isSuccess"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
