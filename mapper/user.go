package mapper

import (
	"be-pbo/dto"
	"be-pbo/models"
	"time"
)

func UserToDTO(u models.User) dto.UserResponse {
	return dto.UserResponse {
		ID:               u.ID,
		Name:             u.Name,
		Email:            u.Email,
	}
}

func UpdateUser(u *models.User, dto dto.UserRequest) {
	if dto.Name != "" {
		u.Name = dto.Name
	}
	if dto.Email != "" {
		u.Email = dto.Email
	}
	u.UpdatedAt = time.Now()
}