package config

import(
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"log"
)

// RESPONSE
func GetConection(URL string) []byte {
	response, err := http.Get(URL)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}