package client

import (
	"context"
	"strings"

	"github.com/amazeeio/lagoon-client-go/schema"
)

// ProjectByName queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectByName(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest(
		map[string]interface{}{
			"name": name,
		},
		Request{
			query:    "_lgraphql/projectByName.graphql",
			fragment: "_lgraphql/fragments/fragmentProjectByName.graphql",
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// Me queries the Lagoon API for me, and
// unmarshals the response into project.
func (c *Client) Me(
	ctx context.Context, user *schema.User) error {

	req, err := c.newRequest(nil, Request{
		query:    "_lgraphql/me.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"me"`
	}{
		Response: user,
	})
}

// EnvironmentByName queries the Lagoon API for an environment by its name and
// parent projectID, and unmarshals the response into environment.
func (c *Client) EnvironmentByName(ctx context.Context, name string,
	projectID uint, environment *schema.Environment) error {

	req, err := c.newRequest(
		map[string]interface{}{
			"name":    name,
			"project": projectID,
		},
		Request{
			query:    "_lgraphql/environmentByName.graphql",
			fragment: "",
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// LagoonAPIVersion queries the Lagoon API for its version, and
// unmarshals the response.
func (c *Client) LagoonAPIVersion(
	ctx context.Context, lagoonAPIVersion *schema.LagoonVersion) error {

	req, err := c.newRequest(nil, Request{
		query:    "_lgraphql/lagoonVersion.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &lagoonAPIVersion)
}

// LagoonSchema queries the Lagoon API for its schema, and
// unmarshals the response.
func (c *Client) LagoonSchema(
	ctx context.Context, lagoonSchema *schema.LagoonSchema) error {

	req, err := c.newRequest(nil, Request{
		query:    "_lgraphql/lagoonSchema.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.LagoonSchema `json:"__schema"`
	}{
		Response: lagoonSchema,
	})
}

// MinimalProjectByName queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) MinimalProjectByName(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest(
		map[string]interface{}{
			"name": name,
		},
		Request{
			query:    "_lgraphql/projectByName.graphql",
			fragment: "_lgraphql/fragments/fragmentMinimalProjectByName.graphql",
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// ProjectsByMetadata queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectsByMetadata(
	ctx context.Context, key string, value string, projects *[]schema.ProjectMetadata) error {

	req, err := c.newRequest(
		map[string]interface{}{
			"key":   key,
			"value": value,
		},
		Request{
			query:    "_lgraphql/projectsByMetadata.graphql",
			fragment: "",
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.ProjectMetadata `json:"projectsByMetadata"`
	}{
		Response: projects,
	})
}

// AllOpenshifts queries the Lagoon API for allopenshifts, and
// unmarshals the response.
func (c *Client) AllOpenshifts(
	ctx context.Context, allOpenshifts *[]schema.Openshift) error {

	req, err := c.newRequest(nil, Request{
		query:    "_lgraphql/allOpenshifts.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Openshift `json:"allOpenshifts"`
	}{
		Response: allOpenshifts,
	})
}

// AllEnvironments queries the Lagoon API for all environments, and
// unmarshals the response.
func (c *Client) AllEnvironments(
	ctx context.Context, envType string, deleted bool, allEnvironments *[]schema.AllEnvironments, request Request) error {
	gqlVars := map[string]interface{}{}
	if envType != "" {
		gqlVars = map[string]interface{}{
			"type": strings.ToUpper(envType),
		}
	}
	request.query = "_lgraphql/allEnvironments.graphql"
	request.fragment = "_lgraphql/fragments/fragmentMinimalAllEnvironments.graphql"
	req, err := c.newRequest(gqlVars, request)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *[]schema.AllEnvironments `json:"allEnvironments"`
	}{
		Response: allEnvironments,
	})
}

// AllEnvironmentsByProject queries the Lagoon API for all environments by project,
// and unmarshals the response.
func (c *Client) AllEnvironmentsByProject(
	ctx context.Context, projectName string, envType string, deleted bool, allEnvironments *schema.ProjectEnvironments) error {
	gqlVars := map[string]interface{}{}
	if envType != "" {
		gqlVars = map[string]interface{}{
			"type": strings.ToUpper(envType),
		}
	}
	gqlVars["name"] = projectName
	gqlVars["deleted"] = deleted
	req, err := c.newRequest(gqlVars, Request{
		query:    "_lgraphql/projectByName.graphql",
		fragment: "_lgraphql/fragments/fragmentMinimalProjectByName.graphql",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectEnvironments `json:"projectByName"`
	}{
		Response: allEnvironments,
	})
}

// EnvironmentByOpenshiftProjectName queries the Lagoon API for environments by openshift projectname,
// and unmarshals the response.
func (c *Client) EnvironmentByOpenshiftProjectName(
	ctx context.Context, openshiftProjectName string, environment *schema.AllEnvironments, request Request) error {
	request.query = "_lgraphql/environmentByOpenshiftProjectName.graphql"
	request.fragment = "_lgraphql/fragments/fragmentMinimalEnvironmentByOpenshiftProjectname.graphql"
	req, err := c.newRequest(
		map[string]interface{}{
			"openshiftProjectName": openshiftProjectName,
		},
		request,
	)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.AllEnvironments `json:"environmentByOpenshiftProjectName"`
	}{
		Response: environment,
	})
}
