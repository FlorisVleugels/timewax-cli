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

func GetToken(client string, username string, password string) {

    xmlData := &request{client, username, password}

    xmlBody, err := xml.MarshalIndent(xmlData, "", "  ")
    if err != nil {
        fmt.Printf("Error marshalling XML: %v\n", err)
        return
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
            fmt.Println("Error unmarshalling XML:", err)
            return
        }
        fmt.Println("Token:", resp.Token)
    } else {
        fmt.Println("Error:", resp.StatusCode)
    }

}
