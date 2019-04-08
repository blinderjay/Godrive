package auth

import (
	"log"

	"google.golang.org/api/drive/v3"
)

func GetService() (*drive.Service, error) {
	config, err := generateConf(drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Fatalln(err)
	}
	client := getClient(config)

	return drive.New(client)

}
