package client

import (
	"context"
	"fmt"
	"regexp"

	"github.com/amazeeio/lagoon-client-go/schema"
)

var duplicate = regexp.MustCompile("^graphql: Duplicate entry ")

// wrapErr wraps a response error with a ErrExist type if the response
// is due to an object already existing
func wrapErr(err error) error {
	if err != nil && duplicate.MatchString(err.Error()) {
		return fmt.Errorf("couldn't create object: %w: %v", ErrExist, err)
	}
	return err
}

// AddGroup adds a group.
func (c *Client) AddGroup(
	ctx context.Context, in *schema.AddGroupInput, out *schema.Group) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addGroup.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addGroup"`
	}{
		Response: out,
	})
}

// AddUser adds a user.
func (c *Client) AddUser(
	ctx context.Context, in *schema.AddUserInput, out *schema.User) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addUser.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"addUser"`
	}{
		Response: out,
	})
}

// AddSSHKey adds an SSH key to a user.
func (c *Client) AddSSHKey(
	ctx context.Context, in *schema.AddSSHKeyInput, out *schema.SSHKey) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addSshKey.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.SSHKey `json:"addSshKey"`
	}{
		Response: out,
	})
}

// AddUserToGroup adds a user to a group.
func (c *Client) AddUserToGroup(
	ctx context.Context, in *schema.UserGroupRoleInput, out *schema.Group) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addUserToGroup.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addUserToGroup"`
	}{
		Response: out,
	})
}

// AddNotificationSlack defines a Slack notification.
func (c *Client) AddNotificationSlack(ctx context.Context,
	in *schema.AddNotificationSlackInput, out *schema.NotificationSlack) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addNotificationSlack.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationSlack `json:"addNotificationSlack"`
	}{
		Response: out,
	})
}

// AddNotificationRocketChat defines a RocketChat notification.
func (c *Client) AddNotificationRocketChat(ctx context.Context,
	in *schema.AddNotificationRocketChatInput,
	out *schema.NotificationRocketChat) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addNotificationRocketChat.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationRocketChat `json:"addNotificationRocketChat"`
	}{
		Response: out,
	})
}

// AddNotificationEmail defines an Email notification.
func (c *Client) AddNotificationEmail(ctx context.Context,
	in *schema.AddNotificationEmailInput,
	out *schema.NotificationEmail) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addNotificationEmail.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationEmail `json:"addNotificationEmail"`
	}{
		Response: out,
	})
}

// AddNotificationMicrosoftTeams defines a MicrosoftTeams notification.
func (c *Client) AddNotificationMicrosoftTeams(ctx context.Context,
	in *schema.AddNotificationMicrosoftTeamsInput,
	out *schema.NotificationMicrosoftTeams) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addNotificationMicrosoftTeams.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationMicrosoftTeams `json:"addNotificationMicrosoftTeams"`
	}{
		Response: out,
	})
}

// AddProject adds a project.
func (c *Client) AddProject(
	ctx context.Context, in *schema.AddProjectInput, out *schema.Project) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addProject.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addProject"`
	}{
		Response: out,
	}))
}

// AddEnvVariable adds an EnvVariable to an Environment or Project.
func (c *Client) AddEnvVariable(ctx context.Context,
	in *schema.EnvVariableInput, out *schema.EnvKeyValue) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addEnvVariable.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.EnvKeyValue `json:"addEnvVariable"`
	}{
		Response: out,
	})
}

// AddOrUpdateEnvironment adds or updates a Project Environment.
func (c *Client) AddOrUpdateEnvironment(ctx context.Context,
	in *schema.AddEnvironmentInput, out *schema.Environment) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addOrUpdateEnvironment.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"addOrUpdateEnvironment"`
	}{
		Response: out,
	}))
}

// AddGroupsToProject adds Groups to a Project.
func (c *Client) AddGroupsToProject(ctx context.Context,
	in *schema.ProjectGroupsInput, out *schema.Project) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addGroupsToProject.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addGroupsToProject"`
	}{
		Response: out,
	})
}

// AddNotificationToProject adds a Notification to a Project.
func (c *Client) AddNotificationToProject(ctx context.Context,
	in *schema.AddNotificationToProjectInput, out *schema.Project) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addNotificationToProject.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addNotificationToProject"`
	}{
		Response: out,
	})
}

// AddBillingGroup adds a Billing Group.
func (c *Client) AddBillingGroup(ctx context.Context,
	in *schema.AddBillingGroupInput, out *schema.BillingGroup) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addBillingGroup.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.BillingGroup `json:"addBillingGroup"`
	}{
		Response: out,
	})
}

// AddProjectToBillingGroup adds a Project to a Billing Group.
func (c *Client) AddProjectToBillingGroup(ctx context.Context,
	in *schema.ProjectBillingGroupInput, out *schema.Project) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addProjectToBillingGroup.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addProjectToBillingGroup"`
	}{
		Response: out,
	})
}

// DeployEnvironmentLatest deploys a latest environment.
func (c *Client) DeployEnvironmentLatest(ctx context.Context,
	in *schema.DeployEnvironmentLatestInput, out *schema.DeployEnvironmentLatest) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/deployEnvironmentLatest.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPullrequest deploys a pullreguest.
func (c *Client) DeployEnvironmentPullrequest(ctx context.Context,
	in *schema.DeployEnvironmentPullrequestInput, out *schema.DeployEnvironmentPullrequest) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/deployEnvironmentPullrequest.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPromote promotes one environment into a new environment.
func (c *Client) DeployEnvironmentPromote(ctx context.Context,
	in *schema.DeployEnvironmentPromoteInput, out *schema.DeployEnvironmentPromote) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/deployEnvironmentPromote.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentBranch deploys a branch.
func (c *Client) DeployEnvironmentBranch(ctx context.Context,
	in *schema.DeployEnvironmentBranchInput, out *schema.DeployEnvironmentBranch) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/deployEnvironmentBranch.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// UpdateProjectMetadata updates a projects metadata.
func (c *Client) UpdateProjectMetadata(
	ctx context.Context, id int, key string, value string, projects *schema.ProjectMetadata) error {

	req, err := c.newRequest(
		map[string]interface{}{
			"id":    id,
			"key":   key,
			"value": value,
		},
		Request{
			query:    "_lgraphql/updateProjectMetadata.graphql",
			fragment: "",
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"updateProjectMetadata"`
	}{
		Response: projects,
	})
}

// RemoveProjectMetadataByKey removes metadata from a project for given key.
func (c *Client) RemoveProjectMetadataByKey(
	ctx context.Context, id int, key string, projects *schema.ProjectMetadata) error {

	req, err := c.newRequest(
		map[string]interface{}{
			"id":  id,
			"key": key,
		},
		Request{
			query:    "_lgraphql/removeProjectMetadataByKey.graphql",
			fragment: "",
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"removeProjectMetadataByKey"`
	}{
		Response: projects,
	})
}

// AddOpenshift adds an Openshift.
func (c *Client) AddOpenshift(
	ctx context.Context, in *schema.AddOpenshiftInput, out *schema.Openshift) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/addOpenshift.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Openshift `json:"addOpenshift"`
	}{
		Response: out,
	})
}

// UpdateOpenshift updates an Openshift.
func (c *Client) UpdateOpenshift(
	ctx context.Context, in *schema.UpdateOpenshiftInput, out *schema.Openshift) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/updateOpenshift.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Openshift `json:"updateOpenshift"`
	}{
		Response: out,
	})
}

// DeleteOpenshift deletes an Openshift.
func (c *Client) DeleteOpenshift(
	ctx context.Context, in *schema.DeleteOpenshiftInput, out *schema.DeleteOpenshift) error {
	req, err := c.newRequest(in, Request{
		query:    "_lgraphql/deleteOpenshift.graphql",
		fragment: "",
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}
