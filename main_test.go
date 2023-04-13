package main_test

import (
    "context"
    //"fmt"
    "log"
	"testing"

    "google.golang.org/api/cloudresourcemanager/v1"
    "google.golang.org/api/option"
)

func TestProjectExist(t *testing.T) {
	expectedProjectName := ""
    // Set up the Google Cloud SDK credentials & project ID.
    ctx := context.Background()

    // Initialize a new Cloud Resource Manager client.
    client, err := cloudresourcemanager.NewService(ctx, option.WithCredentialsFile("path/to/credentials.json"))
    if err != nil {
        log.Fatalf("Failed to create Cloud Resource Manager client: %v", err)
    }

    // Fetch the list of projects.
    projectsList, err := client.Projects.List().Do()
    if err != nil {
        log.Fatalf("Failed to fetch list of projects: %v", err)
    }

    // Print the project information.
    for _, project := range projectsList.Projects {
		if project.ProjectId == expectedProjectName {
			t.Logf("Project '%s' exist.", project.ProjectId)
			return // Success!
		}
	}
    //if we got here , the project does not exist
	t.Errorf("Project does not exist")
}
