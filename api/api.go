package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type request struct {
	Client   string   `xml:"client"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

type Response struct {
    XMLName    xml.Name `xml:"response"`
    Valid      string   `xml:"valid"`
    Token      string   `xml:"token"`
    ValidUntil string   `xml:"validUntil"`
}

const baseURL = "https://api.timewax.com/"

func GetToken(client string, username string, password string) (string, error) {

    xmlData := &request{client, username, password}

    xmlBody, err := xml.MarshalIndent(xmlData, "", "  ")
    if err != nil {
        return "", fmt.Errorf("Error marshalling XML: %v\n", err)
    }

    resp, err := http.Post(baseURL+"authentication/token/get/", "text/xml", strings.NewReader(string(xmlBody)))
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }

        bodyString := string(bodyBytes)

        var resp Response
        if err := xml.Unmarshal([]byte(bodyString), &resp); err != nil {
            return "", fmt.Errorf("Error unmarshalling XML: %v", err)
        }

        return resp.Token, nil
    } else {
        return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
    }
}
