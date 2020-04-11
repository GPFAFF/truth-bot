package credentials

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey, ConsumerSecret, AccessToken, AccessTokenSecret string
}

func GetClient(creds *Credentials) (*twitter.Client, error) {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecrect)

	token := oauth1.NewConfig(creds.AccessToken, creds.AccessTokenSecrect)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil

}
