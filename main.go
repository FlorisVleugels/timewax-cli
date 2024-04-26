package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
    Client string `json:"client"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type request struct {
	Client   string   `xml:"client"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

const baseURL = "https://api.timewax.com/"

func main() {

    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }

    configFile, err := os.Open(homeDir+"/.config/timewax-cli/config.json")
    if err != nil {
        log.Fatal(err)
    }

    var config Config 
    err = json.NewDecoder(configFile).Decode(&config)
    if err != nil {
        log.Fatal(err)
    }

    getToken := &request{config.Client, config.Username, config.Password}

    xmlData, err := xml.MarshalIndent(getToken, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling XML: %v\n", err)
		return
	}

    resp, err := http.Post(baseURL+"authentication/token/get/", "text/xml", strings.NewReader(string(xmlData)))
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        bodyString := string(bodyBytes)
        fmt.Println(bodyString)
    } else {
        fmt.Println("Error:", resp.StatusCode)
    }
}
