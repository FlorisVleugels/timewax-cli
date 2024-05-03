package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type tokenrequest struct {
    XMLName xml.Name `xml:"request"`
	Client   string   `xml:"client"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

type tokenframe struct {
    XMLName    xml.Name `xml:"response"`
    Valid      string   `xml:"valid"`
    Token      string   `xml:"token"`
    ValidUntil string   `xml:"validUntil"`
}

type timelist struct {
    XMLName xml.Name `xml:"request"`
    Token string `xml:"token"`
	DateFrom string   `xml:"dateFrom"`
	DateTo string   `xml:"dateTo"`
    Resource string `xml:"resource"` 
    Entries string `xml:"onlyApprovedEntries"` 
}

const baseURL = "https://api.timewax.com/"

func GetToken(client string, username string, password string) (string, error) {

    xmlData := &tokenrequest{Client: client, Username: username, Password: password}

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

        var resp tokenframe
        if err := xml.Unmarshal([]byte(bodyString), &resp); err != nil {
            return "", fmt.Errorf("Error unmarshalling XML: %v", err)
        }

        return resp.Token, nil
    } else {
        return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
    }
}

func ListTimeEntries(token string, name string) (string, error) {
    
    xmlData := &timelist{Token: token, DateFrom: "20240201", DateTo: "20240202", Entries: "No", Resource: name}

    xmlBody, err := xml.MarshalIndent(xmlData, "", "  ")
    if err != nil {
        return "", fmt.Errorf("Error marshalling XML: %v\n", err)
    }

    resp, err := http.Post(baseURL+"time/entries/list/", "text/xml", strings.NewReader(string(xmlBody)))
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
    }
    return "", nil
}
