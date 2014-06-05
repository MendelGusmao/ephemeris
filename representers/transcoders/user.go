package transcoders

import (
	"ephemeris/models"
	"ephemeris/representers"
)

func UserFromRequest(from *representers.UserRequest, to *models.User) {
	to.Username = from.Username

	if from.Password != nil {
		to.Password = *from.Password
	}

	to.Administrator = from.Administrator
}

func UserToResponse(from *models.User) representers.UserResponse {
	return representers.UserResponse{
		Id:            from.Id,
		Username:      from.Username,
		Administrator: from.Administrator,
	}
}

func UserToResponsePrivate(from *models.User) representers.UserResponse {
	return representers.UserResponse{
		Username: from.Username,
	}
}
