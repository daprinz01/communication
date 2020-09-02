package models

// ErrorResponse is used to send error messages
type ErrorResponse struct {
	Errorcode    string `json:"responseCode"`
	ErrorMessage string `json:"responseDescription"`
}

// SuccessResponse is used to send success messages
type SuccessResponse struct {
	ResponseCode        string      `json:"responseCode"`
	ResponseDescription string      `json:"responseDescription"`
	ResponseMessage     interface{} `json:"responseDetails"`
}
