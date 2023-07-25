package mongo

import (
	. "fibergo-kickstarter/types"

	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadPreferences() *Preferences {
	preferences := new(Preferences)

	file, err := ioutil.ReadFile("lib.json")

	if err != nil {
		log.Fatal("error reading preferences.json: ", err)
	}

	err = json.Unmarshal(file, &preferences)

	if err != nil {
		log.Fatal("error unmarshalling preferences.json: ", err)
	}

	return preferences

}
