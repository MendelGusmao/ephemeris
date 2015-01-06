package models

func UserFromRequest(from *UserRequest, to *User) {
	to.Username = from.Username

	if from.Password != nil {
		to.Password = *from.Password
	}

	to.Role = UserRole(from.Role)
}

func UserToResponse(from *User) UserResponse {
	return UserResponse{
		Id:       from.Id,
		Username: from.Username,
		Role:     int(from.Role),
	}
}

func UserToResponsePrivate(from *User) UserResponse {
	return UserResponse{
		Username: from.Username,
	}
}
