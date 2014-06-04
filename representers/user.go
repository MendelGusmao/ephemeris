package representers

type UserResponse struct {
	Id            int
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	Administrator bool   `json:"administrator,omitempty"`
}

type UserRequest struct {
	Username      string  `json:"username,omitempty"`
	Password      *string `json:"password,omitempty"`
	Administrator bool    `json:"administrator,omitempty"`
}
