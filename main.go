package main

import (
	"encoding/json"
	"log"
	"os"
	"timewax-cli/api"
)

type Config struct {
    Client string `json:"client"`
    Username string `json:"username"`
    Password string `json:"password"`
}

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

    api.GetToken(config.Client, config.Username, config.Password)
}
