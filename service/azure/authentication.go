package azure_authentication

type AuthenResp struct {
	Access_token  string
	Refresh_token string
}

func Authenticate(token string) AuthenResp {

	return AuthenResp{
		Access_token:  "access_token",
		Refresh_token: "refresh_token",
	}
}
