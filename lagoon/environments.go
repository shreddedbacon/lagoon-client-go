// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-client-go/lagoon/client"
	"github.com/amazeeio/lagoon-client-go/schema"
)

// Environments interface contains methods for interacting with environments in lagoon.
type Environments interface {
	AllEnvironments(ctx context.Context, envType string, deleted bool, result *[]schema.AllEnvironments, request client.Request) error
	AllEnvironmentsByProject(ctx context.Context, projectName string, envType string, deleted bool, result *schema.ProjectEnvironments) error
	EnvironmentByOpenshiftProjectName(ctx context.Context, openshiftProjectName string, result *schema.AllEnvironments, request client.Request) error
}

// GetAllEnvironments deploys the latest environment.
func GetAllEnvironments(ctx context.Context, envType string, deleted bool, request client.Request, e Environments) (*[]schema.AllEnvironments, error) {
	result := []schema.AllEnvironments{}
	return &result, e.AllEnvironments(ctx, envType, deleted, &result, request)
}

// GetAllEnvironmentsByProject deploys the latest environment.
func GetAllEnvironmentsByProject(ctx context.Context, projectName string, envType string, deleted bool, e Environments) (*schema.ProjectEnvironments, error) {
	result := schema.ProjectEnvironments{}
	return &result, e.AllEnvironmentsByProject(ctx, projectName, envType, deleted, &result)
}

// GetEnvironmentByOpenshiftProjectName deploys the latest environment.
func GetEnvironmentByOpenshiftProjectName(ctx context.Context, openshiftProjectName string, request client.Request, e Environments) (*schema.AllEnvironments, error) {
	result := schema.AllEnvironments{}
	return &result, e.EnvironmentByOpenshiftProjectName(ctx, openshiftProjectName, &result, request)
}

// Request is a way to provide the custom query or fragments for a request
type Request struct {
	CustomQuery      []byte
	CustomFragment   []byte
	queryLocation    string
	fragmentLocation string
}
