package response

type UserLoginResponse struct {
	AccessToken string `json:"access_token"`
}

type UserRegisterResponse struct {
	VerificationId string `json:"verificationId"`
}

type UserVerifyOTPResponse struct {
	AccessToken string `json:"access_token"`
}

type OTPData struct {
	UserId int32  `json:"userId"`
	OTP    string `json:"otp"`
}
