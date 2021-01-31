package taskswrapper

import (
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// GetTasksService Creates the tasks service using the files
func GetTasksService() (*tasks.Service, error) {
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
func GetAllTasksLists(srv *tasks.Service) (*tasks.TaskLists, error) {
	userLists, err := srv.Tasklists.List().Do()
	return userLists, err
}

// GetTaskList returns TaskList object at index idx
func GetTaskList(srv *tasks.Service, idx int) (*tasks.TaskList, error) {
	var list *tasks.TaskList
	userLists, err := GetAllTasksLists(srv)
	if len(userLists.Items) > (idx - 1) {
		list = userLists.Items[idx-1]
	}
	return list, err
}

// GetFirstTaskList fetches the first Task List from the user
func GetFirstTaskList(srv *tasks.Service) (*tasks.TaskList, error) {
	var firstList *tasks.TaskList
	userLists, err := GetAllTasksLists(srv)
	if len(userLists.Items) > 0 {
		firstList = userLists.Items[0]
	}
	return firstList, err
}

// GetFirstTaskList Sends the list items
func GetAllTaskListItems(srv *tasks.Service, id string) (*tasks.Tasks, error) {
	return srv.Tasks.List(id).Do()
}
