package utils

type successResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type failedResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func Failed(msgError string) failedResponse {
	failed := failedResponse{
		Status: false,
		Error:  msgError,
	}
	return failed

}

func Success(data interface{}) successResponse {
	payload := successResponse{
		Status: true,
		Data:   data,
	}
	return payload
}
