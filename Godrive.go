package main

import (
	"fmt"
	"log"

	"github.com/blinderjay/Godrive/auth"
)

func main() {

	// If modifying these scopes, delete your previously saved token.json.

	srv, err := auth.GetService()
	if err != nil {
		log.Fatalln(err)
	}
	r, err := srv.Files.List().PageSize(35).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalln(err)
	}

	var flist = "Files:\n"
	if len(r.Files) == 0 {
		fmt.Printf("No file found")
	} else {
		flist += "\n"
		for _, i := range r.Files {
			flist += i.Name + "(" + i.Id + ")\n"
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
