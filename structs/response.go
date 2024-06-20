package structs

import (
	"time"
)

type CommonResponse struct {
	Total       int64       `json:"total"`
	ResultCount int64       `json:"resultCount"`
	Success     int64       `json:"success"`
	Failed      int64       `json:"failed"`
	StatusCode  int         `json:"statusCode"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

type SimpleCommonResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}
