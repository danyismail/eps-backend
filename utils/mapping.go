package utils

import (
	"eps-backend/model"
	"eps-backend/structs"
	"strings"
)

func MappingUserResponse(users []model.User) []structs.UserResponse {
	mapUsers := []structs.UserResponse{}
	for _, v := range users {
		user := structs.UserResponse{
			ID:        strings.ToUpper(v.ID.String()),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Username:  v.Username,
			Email:     v.Email,
			Role:      v.Role.Name,
		}
		mapUsers = append(mapUsers, user)
	}
	return mapUsers
}
