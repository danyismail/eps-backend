package utils

import (
	"eps-backend/model"
	"eps-backend/structs"
)

func MappingUserResponse(users []model.User) []structs.UserResponse {
	mapUsers := []structs.UserResponse{}
	for _, v := range users {
		user := structs.UserResponse{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Username:  v.Username,
			Email:     v.Email,
			Role:      v.Role,
		}
		mapUsers = append(mapUsers, user)
	}
	return mapUsers
}
