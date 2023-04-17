package main

import (
	"fmt"
	"log"
    compute "cloud.google.com/go/compute/apiv1"
    "context"
    "google.golang.org/api/iterator"
    computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	
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

}

func network() {
    ctx := context.Background()
    c, err := compute.NewNetworksRESTClient(ctx)
    if err != nil {
            // TODO: Handle error.
    }
    defer c.Close()

    req := &computepb.ListNetworksRequest{
            // TODO: Fill request struct fields.
            // See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/compute/v1#ListNetworksRequest.
    }
    it := c.List(ctx, req)
    for {
            resp, err := it.Next()
            if err == iterator.Done {
                    break
            }
            if err != nil {
                    // TODO: Handle error.
            }
            // TODO: Use resp.
            _ = resp
    }
}
