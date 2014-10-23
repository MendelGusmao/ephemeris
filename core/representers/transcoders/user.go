package transcoders

import (
	"ephemeris/core/models"
	"ephemeris/core/representers"
)

func UserFromRequest(from *representers.UserRequest, to *models.User) {
	to.Username = from.Username

	if from.Password != nil {
		to.Password = *from.Password
	}

	to.Role = models.UserRole(from.Role)
}

func UserToResponse(from *models.User) representers.UserResponse {
	return representers.UserResponse{
		Id:       from.Id,
		Username: from.Username,
		Role:     int(from.Role),
	}
}

func UserToResponsePrivate(from *models.User) representers.UserResponse {
	return representers.UserResponse{
		Username: from.Username,
	}
}
