package client

import (
	"log"

	"github.com/mailup/content-provider-go/model"
	"github.com/maxzerbini/sling"
)

func Authorize(clientid, clientsecret string) (token *model.TokenResponse, err error) {
	params := &model.OAuthParams{
		ClientId:     clientid,
		ClientSecret: clientsecret,
		GrantType:    "client_credentials",
	}
	base := sling.New().Base("https://mapi-auth.nweb.info").Post("/oauth/token").BodyForm(params)
	//base := sling.New().Base("http://localhost:4200").Post("/oauth/token").BodyForm(params)
	token = &model.TokenResponse{Properties: make(map[string]string)}
	_, err = base.ReceiveSuccess(token)
	if err != nil {
		log.Printf("Authorization failed: %v", err)
		return nil, err
	}
	return token, nil
}

func SendMessage(msg string, token model.TokenResponse) (err error) {
	return nil
}
