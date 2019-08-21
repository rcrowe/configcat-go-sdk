package main

import (
	"fmt"
	"gopkg.in/configcat/go-sdk.v1"
)

func main() {
	client := configcat.NewClient("PKDVCLf-Hq-h-kCzMp-L7Q/psuH7BGHoUmdONrzzUOY7A")

	// create a user object to identify your user (optional)
	user := configcat.NewUser("##SOME-USER-IDENTIFICATION##")

	// get individual config values identified by a key and a user object
	value := client.GetValueForUser("keySampleText", "", user)

	fmt.Println("keySampleText: ", value)
}