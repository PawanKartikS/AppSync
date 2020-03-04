// github.go
// Contains the code responsible for communicating with the GitHub API and returning
// the matching gist URLs specified through appName via Init()

package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserInstance struct {
	// Do not expose the members to non-packaged code
	githubUserName string
	appName        string
}

func Init(githubUserName string, appName string) *UserInstance {
	return &UserInstance{githubUserName: githubUserName, appName: appName}
}

// Reads the API response (pointed by endpoint) and returns it in the string format
func readEndpointResponse(endpoint string) string {
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

// Lists all the gists that the user has made public
func listUserGists(instance *UserInstance) interface{} {
	endpoint := fmt.Sprintf("https://api.github.com/users/%s/gists", instance.githubUserName)
	response := readEndpointResponse(endpoint)

	var parsedData []map[string]interface{}
	err := json.Unmarshal([]byte(response), &parsedData)
	if err != nil {
		panic(err)
	}

	return parsedData
}

// Gets all the gists that the user has requested through Init()
func GetMatchingGistsUrls(instance *UserInstance) []string {
	urls := make([]string, 1)

	gists := listUserGists(instance).([]map[string]interface{})
	for _, gist := range gists {
		files := gist["files"].(map[string]interface{})
		for _, file := range files {
			parsedData := file.(map[string]interface{})
			filename := parsedData["filename"].(string)

			// it is mandatory that the filename contain the appName
			// ABC_settings.json - matched with the program name ABC (appName)
			if strings.Contains(filename, instance.appName) {
				urls = append(urls, parsedData["raw_url"].(string))
			}
		}
	}

	return urls
}
