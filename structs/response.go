package structs

type CommonResponse struct {
	Total       int64       `json:"total"`
	ResultCount int64       `json:"resultCount"`
	Success     int64       `json:"success"`
	Failed      int64       `json:"failed"`
	Data        interface{} `json:"data"`
	StatusCode  int         `json:"statusCode"`
	Message     string      `json:"message"`
}
