// A generated module for BsfWithDagger functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
)

type BsfWithDagger struct{}

// Returns a container that echoes whatever string argument is provided
func (m *BsfWithDagger) ContainerEcho(stringArg string) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *BsfWithDagger) GrepDir(ctx context.Context, image string, dir *Directory) (string, error) {
	return dag.Container().
		From(image).
		WithMountedDirectory("/mnt", dir).
		WithWorkdir("/mnt").
		WithExec([]string{"golangci-lint run"}).
		Stdout(ctx)
}
