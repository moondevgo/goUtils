package guCloud

import (
	// "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	basic "github.com/moondevgo/goUtils/guBasic"
)

const (
	TokenPath = "token.json"
)

var Scopes = map[string][]string{
	"sheets": {"https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/drive", "https://www.googleapis.com/auth/calendar"},
	"script": {"https://www.googleapis.com/auth/script", "https://www.googleapis.com/auth/script.projects"},
	// "script": {"https://www.googleapis.com/auth/script.projects"},
}

// * google 접속 설정 .json 경로
func GetGoogleJsonPath(nick, authType string) string {
	root := basic.GetRootFolder("DEV_CFG_ROOT")
	paths := basic.GetConfigMap(basic.SetFilePath(root, "paths.yaml"), "CLOUD")
	return basic.SetFilePath(root, paths["CFG_SUBROOT"].(string), paths["CFG_GOOGLE"].(string), authType+"_"+nick+".json")
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := TokenPath // TODO: 함수화
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
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
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

// api string *http.Client
func ApiClient(api_name, bot_nick, user_nick string) *http.Client {
	nick := bot_nick
	authType := "bot"
	if user_nick != "" {
		nick = user_nick
		authType = "user"
	}
	path := GetGoogleJsonPath(nick, authType)

	b := basic.BytesFromFile(path)

	var client *http.Client
	if authType == "bot" {
		config, err := google.JWTConfigFromJSON(b, Scopes[api_name]...)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		// client = config.Client(oauth2.NoContext)
		client = config.Client(context.Background())
	} else if authType == "user" {
		config, err := google.ConfigFromJSON(b, Scopes[api_name]...)
		fmt.Printf("\n*****config: %T\n", config)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		client = getClient(config)
	}

	return client
}
