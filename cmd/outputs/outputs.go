package outputs

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dannielss/go-github-cli/cmd/model"
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
	-u, --user get all repositories, to export as csv use -u [USER]).
	-h, --help print all cli options.
	-e, --export export as csv
	-w, --number of workers to export concurrently
		`)
}

func GetRepositoriesInfo(user string) {
	var repositories []model.Repository

	response, err := services.GetAllRepositories(user)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(response, &repositories)

	for _, repo := range repositories {
		fmt.Printf("Repository name: %s\n", repo.Name)
		fmt.Printf("Repository url: %s\n\n", repo.Html_url)
	}

	fmt.Printf("Total of repositories: %v\n", len(repositories))
}

func ExportAsCSV(user string) {
	start := time.Now()

	var repositories []model.Repository

	response, err := services.GetAllRepositories(user)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(response, &repositories)

	records := [][]string{
		{"name", "repository_url"},
	}

	for i := 0; i < len(repositories); i++ {
		records = append(records, []string{
			repositories[i].Name,
			repositories[i].Html_url,
		})
	}

	f, err := os.Create("repositories.csv")

	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing repository to file", err)
		}
	}
	fmt.Println(time.Since(start))
}

var wg sync.WaitGroup

func ExportAsCSVConcurrently(user string, numberOfWorkers int) {
	start := time.Now()

	var repositories []model.Repository
	var divided [][]model.Repository

	response, err := services.GetAllRepositories(user)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(response, &repositories)

	divided = divideInChunks(repositories, numberOfWorkers)

	for i := 0; i < len(divided); i += 1 {
		wg.Add(1)
		go writeFile(divided[i], i)
	}

	wg.Wait()

	fmt.Println(time.Since(start))
}

func writeFile(r []model.Repository, i int) {
	fmt.Printf("Comecei a executar %v\n", i)

	records := [][]string{
		{"name", "repository_url"},
	}

	for i := 0; i < len(r); i++ {
		records = append(records, []string{
			r[i].Name,
			r[i].Html_url,
		})
	}

	name := "repositories" + fmt.Sprintf("%v", i) + ".csv"

	f, err := os.Create(name)

	if err != nil {
		fmt.Println("Failed to open file", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing repository to file", err)
		}
	}

	wg.Done()
}

func divideInChunks(r []model.Repository, numberOfWorkers int) [][]model.Repository {
	var chunks [][]model.Repository

	chunkSize := (len(r) + numberOfWorkers - 1) / numberOfWorkers

	for i := 0; i < len(r); i += chunkSize {
		end := i + chunkSize

		if end > len(r) {
			end = len(r)
		}

		chunks = append(chunks, r[i:end])
	}

	return chunks
}
