package main

import (
	"context"
	"fmt"

	"github.com/lieberlois/oci/cmd/shared"
	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
)

func main() {
	// 0. Create a file store
	fs, err := file.New("./ocitestpull")
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	// 1. Connect to a remote repository
	ctx := context.Background()

	reg := shared.Remote
	repo, err := remote.NewRepository(reg + "/hello-artifact")
	if err != nil {
		panic(err)
	}

	repo.PlainHTTP = true

	// 2. Copy from the remote repository to the file store
	tag := shared.Version
	manifestDescriptor, err := oras.Copy(ctx, repo, tag, fs, tag, oras.DefaultCopyOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("manifest descriptor:", manifestDescriptor)
}
