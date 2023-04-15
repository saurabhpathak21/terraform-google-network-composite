package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/compute"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

func main() {
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
        fmt.Printf("- Project ID: %s, Name: %s, Number: %d ", project.ProjectId, project.Name, project.ProjectNumber)
    }
    //Network list
    
    client, err := compute.NewClient(ctx)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    defer client.Close()
    
    networks, err := client.ListNetworks(ctx, "")
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    for _, network := range networks {
        fmt.Println(network.Name)
    }

}
