package response

type UserResponse struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ManyUsersResponse struct {
	Users []UserResponse `json:"users"`
}

type UserAuthToken struct {
	AccessToken string `json:"access_token"`
}
