package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	decoder := json.NewDecoder(os.Stdin)

	for {
		var jsonMap map[string]interface{}

		if err := decoder.Decode(&jsonMap); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error decoding JSON:", err)
			continue
		}

		urlUrl := jsonMap["url"]
		urlTitle := jsonMap["title"]
		urlWebserver := jsonMap["webserver"]
		urlContentType := jsonMap["content_type"]
		urlContentLength := jsonMap["content_length"]
		urlServerResponse := jsonMap["body"]

		urlUrlString := fmt.Sprintf("%v", urlUrl)
		urlTitleString := fmt.Sprintf("%v", urlTitle)
		urlWebserverString := fmt.Sprintf("%v", urlWebserver)
		urlContentTypeString := fmt.Sprintf("%v", urlContentType)
		urlContentLengthString := fmt.Sprintf("%v", urlContentLength)
		urlServerResponseString := fmt.Sprintf("%v", urlServerResponse)

		if urlUrlString != "" {
			fullCommand := "bbrf url add '" + urlUrlString + " " + urlContentLengthString + "' -t 'title:" + urlTitleString + "'" + " -t 'webserver:" + urlWebserverString + "'" + " -t 'contenttype:" + urlContentTypeString + "'" + " -t 'contentlength:" + urlContentLengthString + "'" + " -t 'serverresponse:" + urlServerResponseString + "'" + " -p @INFER"
			out, _ := exec.Command("sh", "-c", fullCommand).Output()
			fmt.Printf("%s", out)
		} else {
			break
		}
	}
}
