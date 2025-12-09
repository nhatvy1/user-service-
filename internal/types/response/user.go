package response

type UserLoginResponse struct {
	UserId      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
}
