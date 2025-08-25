package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var user = "me"


func getApplications(srv *gmail.Service, user string) error {
	pageToken := ""
	query := `after:2024/8/01 ("thank you for applying" OR "we received your application" OR "application submitted" OR "thank you for your interest" OR "we have received your application" OR "your application has been received" OR "we appreciate your interest" OR "application confirmation" OR "we have received your resume" OR "your application is being reviewed" OR "we are reviewing your application" OR "we have your application" OR "your application has been submitted" OR "thank you for your application" OR "we have received your job application") -in:chats -is:sent`
	includeSpamTrash := false
	count := 0

	for {
		req := srv.Users.Messages.List(user).Q(query).IncludeSpamTrash(includeSpamTrash)
		if pageToken != "" {
			req.PageToken(pageToken)
		}

		res, err := req.Do()
		if err != nil {
			return fmt.Errorf("list messages: %w", err)
		}

		for _, msg := range res.Messages {
			msgDetail, err := srv.Users.Messages.Get(user, msg.Id).Format("full").Do()
			if err != nil {
				fmt.Printf("⚠️ Could not fetch message %s: %v\n", msg.Id, err)
				continue
			}

			var subject, from, date string
			for _, header := range msgDetail.Payload.Headers {
				switch header.Name {
				case "Subject":
					subject = header.Value
				case "From":
					from = header.Value
				case "Date":
					date = header.Value
				}
			}

			fmt.Printf("-----\nSubject: %s\nFrom: %s\nDate: %s\n", subject, from, date)

			// Print snippet (short preview of the message)
			fmt.Printf("Snippet: %s\n", msgDetail.Snippet)
			count++
		}

		if res.NextPageToken == "" {
			break
		}
		pageToken = res.NextPageToken
	}
	fmt.Printf("📬 Total job-related emails since Dec 1, 2024: %d\n", count)
	return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}               //creating a struct to hold token data
	err = json.NewDecoder(f).Decode(tok) //decoding the json data in the file and putting it in the tok struct
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	ctx := context.Background() //many go functions require context.context arg but this one doesnt cancel api calls or anything it just chills
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope) //builds the config for oauth
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config) //connects to gmail without manually handling tokens and headers (client is the middleman between api and i)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client)) //srv is now gmail service instant allowing us to call apis
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	getApplications(srv, user)
}

