package auth

var credential string = `{"installed":{"client_id":"782204627780-ct84354d3cjplvvolfmokra4e58rte4h.apps.googleusercontent.com","project_id":"jeditor-1547883087965","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"tZC-0ta-bMAHsn6lVTgTXJH2","redirect_uris":["urn:ietf:wg:oauth:2.0:oob","http://localhost"]}}`

func setCredential(cre string) {
	credential = cre
}

// b, err := ioutil.ReadFile("credentials.json")
// if err != nil {
//         log.Fatalf("Unable to read client secret file: %v", err)
// }
