package dto

type ValidateAuthorizationOutput struct {
	Message   string      `json:"message"`
	IsSuccess bool        `json:"isSuccess"`
	Payload   interface{} `json:"errorMessage,omitempty"`
}
