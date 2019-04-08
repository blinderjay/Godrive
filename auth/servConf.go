package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func generateConf(scope ...string) (*oauth2.Config, error) {
	b := []byte(credential)
	// b, err := ioutil.ReadFile("credentials.json")
	// if err != nil {
	//         log.Fatalf("Unable to read client secret file: %v", err)
	// }
	return google.ConfigFromJSON(b, scope...)
	// If modifying these scopes, delete your previously saved token.json.
}
