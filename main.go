package main

import (
	"encoding/json"
    "encoding/xml"
	"fmt"
	"log"
	"os"
	"timewax-cli/api"
    "time"
)

type Config struct {
    Client string `json:"client"`
    Username string `json:"username"`
    Password string `json:"password"`
    Name string `json:"name"` 
}

type tokenframe struct {
    XMLName    xml.Name `xml:"response"`
    Valid      string   `xml:"valid"`
    Token      string   `xml:"token"`
    ValidUntil string   `xml:"validUntil"`
}

var token string

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

    tokenFile, err := os.Open(homeDir+"/.config/timewax-cli/token.xml")
    if err == nil {

        var tokenData tokenframe
        err = xml.NewDecoder(tokenFile).Decode(&tokenData)
        if err != nil {
            log.Fatal(err)
        } 
        
        currentTime := time.Now()
        validTime, err := time.Parse("20060102T150405", tokenData.ValidUntil)
        if err != nil {
            log.Fatal(err)
        }

        if currentTime.Before(validTime) {
            token = tokenData.Token
        } else {
            token, err =api.GetToken(config.Client, config.Username, config.Password, homeDir)
            if err != nil {
                log.Fatal(err)
            }
        }

    } else {

        token, err =api.GetToken(config.Client, config.Username, config.Password, homeDir)
        if err != nil {
            log.Fatal(err)
        }
    }

    api.ListTimeEntries(token, config.Name)
}
