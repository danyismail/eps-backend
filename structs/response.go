package structs

type CommonResponse struct {
	Total      int         `json:"total,omitempty"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
}
