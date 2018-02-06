package model

type TokenResponse struct {
	Token        string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token"`
	TokenType    string            `json:"token_type"` // bearer
	ExperesIn    int64             `json:"expires_in"` // secs
	Properties   map[string]string `json:"properties"`
}

type OAuthParams struct {
	ClientId     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	GrantType    string `url:"grant_type"`
}
