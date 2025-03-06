package model

type FailedResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}
