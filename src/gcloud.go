package main

import (
	"context"

	"cloud.google.com/go/compute/metadata"
)

func getProjectID() string {
	ctx := context.Background()

	projectID, err := metadata.ProjectIDWithContext(ctx)
	if err != nil {
		return "local"
	}

	return projectID
}
