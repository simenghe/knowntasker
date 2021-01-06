package taskswrapper

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// GetTasksService Creates the tasks service using the files
func GetTasksService() (* tasks.Service, error){
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, tasks.TasksScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}
	ctx := context.Background()
	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	return srv, err
}

// GetAllTasksLists fetches the user's lists
func GetAllTasksLists(srv * tasks.Service) {
	fmt.Println("Get Task List")
	fmt.Println(srv)
}
