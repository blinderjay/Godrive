package godrive

import (
	"C"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/blinderjay/Godrive/env"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	home, err := Home()
	tokFile := home + "/.config/token.json"
	// tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

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

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

//export goauth
func goauth() *C.char {
	var credential string = `{"installed":{"client_id":"782204627780-ct84354d3cjplvvolfmokra4e58rte4h.apps.googleusercontent.com","project_id":"jeditor-1547883087965","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"tZC-0ta-bMAHsn6lVTgTXJH2","redirect_uris":["urn:ietf:wg:oauth:2.0:oob","http://localhost"]}}`
	var b []byte = []byte(credential)
	// b, err := ioutil.ReadFile("credentials.json")
	// if err != nil {
	//         log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
	if err != nil {
		return C.CString("Unable to parse client secret file to config")
	}
	client := getClient(config)

	srv, err := drive.New(client)
	if err != nil {
		return C.CString("Unable to retrieve Drive client")
	}

	r, err := srv.Files.List().PageSize(35).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		return C.CString("Unable to retrieve files")
	}

	var flist = "Files:"
	if len(r.Files) == 0 {
		return C.CString(flist + "No files found.")
	} else {
		flist += "\n"
		for _, i := range r.Files {
			flist += i.Name + "(" + i.Id + ")\n"
			// fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
		return C.CString(flist)
	}
}

// func main() {}
