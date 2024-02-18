package main

import (
	"context"
	"fmt"

	"github.com/lieberlois/oci/cmd/shared"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
)

func main() {
	artifactSources := map[string][]string{
		"ocitestpush":  {"demo.txt", "juche.txt", "demo"},
		"ocitestpush2": {"index.json"},
	}

	for sourceDir, sourceFiles := range artifactSources {
		fmt.Println("Packaging", sourceDir)

		// 0. Create a file store
		fs, err := file.New(sourceDir)
		if err != nil {
			panic(err)
		}
		defer fs.Close()
		ctx := context.Background()

		// 1. Add files to the file store
		mediaType := "application/vnd.test.file"
		fileNames := sourceFiles
		fileDescriptors := make([]v1.Descriptor, 0, len(fileNames))
		for _, name := range fileNames {
			fileDescriptor, err := fs.Add(ctx, name, mediaType, "")
			if err != nil {
				panic(err)
			}

			fileDescriptors = append(fileDescriptors, fileDescriptor)
			fmt.Printf("file descriptor for %s: %v\n", name, fileDescriptor)
		}

		// 2. Pack the files and tag the packed manifest
		artifactType := "application/vnd.test.artifact"
		opts := oras.PackManifestOptions{
			Layers: fileDescriptors,
		}

		manifestDescriptor, err := oras.PackManifest(
			ctx,
			fs,
			oras.PackManifestVersion1_1_RC4,
			artifactType,
			opts,
		)

		if err != nil {
			panic(err)
		}
		fmt.Println("manifest descriptor:", manifestDescriptor)

		tag := shared.Version
		if err = fs.Tag(ctx, manifestDescriptor, tag); err != nil {
			panic(err)
		}

		// 3. Connect to a remote repository
		reg := shared.Remote
		repo, err := remote.NewRepository(reg + "/" + sourceDir)
		if err != nil {
			panic(err)
		}

		repo.PlainHTTP = true

		// 4. Copy from the file store to the remote repository
		_, err = oras.Copy(ctx, fs, tag, repo, tag, oras.DefaultCopyOptions)
		if err != nil {
			panic(err)
		}
	}
}
