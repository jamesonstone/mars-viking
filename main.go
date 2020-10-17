package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jstone28/mars-command/pkg/server"
	"github.com/jstone28/mars-command/pkg/utils"
)

func main() {
	readOpenAPI()
	server.Start()
}

// OpenAPISpec the struct representation of the openapi.yml file
type OpenAPISpec struct {
	Openapi string `json:"openapi"`
	Info    struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Version     string `json:"version"`
	} `json:"info"`
	Servers []struct {
		URL         string `json:"url"`
		Description string `json:"description"`
	} `json:"servers"`
	Paths struct {
		Testjson struct {
			Get struct {
				Summary     string `json:"summary"`
				Description string `json:"description"`
				Responses   struct {
					Num200 struct {
						Description string `json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Type  string `json:"type"`
									Items struct {
										Type string `json:"type"`
									} `json:"items"`
								} `json:"schema"`
							} `json:"application/json"`
						} `json:"content"`
					} `json:"200"`
				} `json:"responses"`
			} `json:"get"`
		} `json:"/testjson"`
	} `json:"paths"`
}

func readOpenAPI() {
	oas := OpenAPISpec{}

	openAPISpec, err := ioutil.ReadFile("openapi.json")
	utils.CheckError(err)

	err = json.Unmarshal(openAPISpec, &oas)
	utils.CheckError(err)

	fmt.Println(string(openAPISpec))
}

func postOpenAPISpec() {

}
