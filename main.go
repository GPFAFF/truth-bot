package main

import (
	"fmt"
	credentials "github.com/gpfaff/truth-bot"
	"os"
)

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	creds := credentials.Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	// Print out the pointer to our client
	// for now so it doesn't throw errors
	fmt.Printf("%+v\n", client)
}
