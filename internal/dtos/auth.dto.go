package dtos

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	DisplayName       string `json:"displayName"`
	Bio               string `json:"bio"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	Role              string `json:"role"`
}
