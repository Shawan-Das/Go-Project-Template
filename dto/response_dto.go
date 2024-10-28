package dto

type ResponseDto struct {
	Message    string      `json:"message"`
	IsSuccess  bool        `json:"isSuccess"`
	Payload    interface{} `json:"payload,omitempty"`
	StatusCode int         `json:"statusCode"`
}

type ResponseDto_v2 struct {
	Message    string      `json:"message"`
	IsSuccess  bool        `json:"isSuccess"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"statusCode"`
}

type ResponseDtoV2 struct {
	Message    string      `json:"message"`
	IsSuccess  bool        `json:"isSuccess"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"statusCode"`
}
