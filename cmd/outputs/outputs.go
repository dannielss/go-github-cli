package outputs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dannielss/go-github-cli/cmd/services"
)

func GetHeader() {
	fmt.Println(`
                                              
	_____ _____ _____ _____ _____ _____    _____ __    _____ 
	|   __|     |_   _|  |  |  |  | __  |  |     |  |  |   |
	|  |  |-   -| | | |     |  |  | __ -|  |   --|  |__|- -|
	|_____|_____| |_| |__|__|_____|_____|  |_____|_____|___|
																														
	
	`)
}

func ShowHelp() {
	fmt.Println(`Usage: CLI Template [OPTIONS]

	Options:
		-s, --string print string input.
		-h, --help print all cli options.
		`)
}

type Repository struct {
	Name     string
	Html_url string
}

func GetRepositoriesInfo(user string) {
	var repositories []Repository

	response, err := services.GetAllRepositories(user)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(response, &repositories)

	for _, repo := range repositories {
		fmt.Printf("Repository name: %s\n", repo.Name)
		fmt.Printf("Repository url: %s\n\n", repo.Html_url)
	}
}
