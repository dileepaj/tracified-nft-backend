package responseDtos

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

type ResultResponse struct {
	Status   int `json:"Status"`
	Response any `json:"Response"`
}