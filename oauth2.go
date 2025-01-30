package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"golang.org/x/oauth2"
)

func oauth2_access() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "9ac83c8f5c7e5eefd986cb3ea2aeac0a15b432dc",
		ClientSecret: "aa7b113b1a6354bd1d4498dea18fc27639d3beae",
		RedirectURL:  "https://example.com/oauth",
		Endpoint: oauth2.Endpoint{
			// AuthURL: "https://launchpad.37signals.com/authorization/new",
			AuthURL: "https://launchpad.37signals.com/authorization/new?type=user_agent",
			// AuthURL:  "https://launchpad.37signals.com/authorization/new?type=web_server",
			// TokenURL: "https://launchpad.37signals.com/authorization/token",
			TokenURL: "https://launchpad.37signals.com/authorization/token?type=user_agent",
			// TokenURL: "https://launchpad.37signals.com/authorization/token?type=web_server",
		},
	}

	verifier := oauth2.GenerateVerifier()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the url for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	log.Println(client)

	resp, err := client.Get("https://launchpad.37signals.com/authorization.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string([]byte(body)))
	log.Printf("%+v\n", resp)
}
