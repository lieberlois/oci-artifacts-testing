package main

import (
	"fmt"
	"log"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/lieberlois/oci/cmd/shared"
)

func main() {
	// Define the images to include in the index
	images := []string{
		"ocitestpush",
		"ocitestpush2",
	}

	// Fetch the images and create the index
	var adds []mutate.IndexAddendum
	for _, img := range images {
		ref, err := name.NewTag(fmt.Sprintf("%s/%s:%s", shared.Remote, img, shared.Version), name.Insecure)
		if err != nil {
			log.Fatalf("Failed to parse reference: %v", err)
		}

		img, err := remote.Image(ref)
		if err != nil {
			log.Fatalf("Failed to get image: %v", err)
		}

		adds = append(adds, mutate.IndexAddendum{
			Add: img,
		})
	}

	// Create the index
	idx := mutate.AppendManifests(empty.Index, adds...)

	// Push the index to the repository
	repo, err := name.NewTag(fmt.Sprintf("%s/myindex:latest", shared.Remote), name.Insecure)
	if err != nil {
		panic(err)
	}

	if err := remote.WriteIndex(repo, idx); err != nil {
		log.Fatalf("Failed to push index to repository: %v", err)
	}

	fmt.Println("Pushed index to repository")
}
