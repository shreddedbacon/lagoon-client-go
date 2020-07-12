package schema

// ProjectAvailability determines the number of pods used to run a project.
type ProjectAvailability string

// High tells Lagoon to load balance across multiple pods.
// Standard tells Lagoon to use a single pod for the site.
const (
	High     ProjectAvailability = "HIGH"
	Standard ProjectAvailability = "STANDARD"
)

// Currency for billing purposes.
type Currency string

// These are the Currency units supported by Lagoon.
const (
	AUD Currency = "AUD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	USD Currency = "USD"
	CHF Currency = "CHF"
	ZAR Currency = "ZAR"
)

// SSHKeyType are the different ssh key types supported in Lagoon.
type SSHKeyType string

// These are the different ssh key types supported in Lagoon.
const (
	SSHRsa     SSHKeyType = "SSH_RSA"
	SSHEd25519 SSHKeyType = "SSH_ED25519"
)

// DeployType are the different deploy types supported in Lagoon.
type DeployType string

// These are the different deploy types supported in Lagoon.
const (
	Branch      DeployType = "BRANCH"
	PullRequest DeployType = "PULLREQUEST"
	Promote     DeployType = "PROMOTE"
)

// EnvType are the different environment types supported in Lagoon.
type EnvType string

// These are the different environment types supported in Lagoon.
const (
	ProductionEnv  EnvType = "PRODUCTION"
	DevelopmentEnv EnvType = "DEVELOPMENT"
)

// NotificationType are the different notification types supported in Lagoon.
type NotificationType string

// These are the different notification types supported in Lagoon.
const (
	SlackNotification          NotificationType = "SLACK"
	RocketChatNotification     NotificationType = "ROCKETCHAT"
	EmailNotification          NotificationType = "EMAIL"
	MicrosoftTeamsNotification NotificationType = "MICROSOFTTEAMS"
)

// DeploymentStatusType are the different deployment status types supported in Lagoon.
type DeploymentStatusType string

// These are the different deployment status types supported in Lagoon.
const (
	NewDeploy       DeploymentStatusType = "NEW"
	PendingDeploy   DeploymentStatusType = "PENDING"
	RunningDeploy   DeploymentStatusType = "RUNNING"
	CancelledDeploy DeploymentStatusType = "CANCELLED"
	ErrorDeploy     DeploymentStatusType = "ERROR"
	FailedDeploy    DeploymentStatusType = "FAILED"
	CompleteDeploy  DeploymentStatusType = "COMPLETE"
)

// EnvVariableType are the different deployment status types supported in Lagoon.
type EnvVariableType string

// These are the different deployment status types supported in Lagoon.
const (
	ProjectVar     EnvVariableType = "PROJECT"
	EnvironmentVar EnvVariableType = "ENVIRONMENT"
)

// EnvVariableScope are the different deployment status types supported in Lagoon.
type EnvVariableScope string

// These are the different deployment status types supported in Lagoon.
const (
	BuildVar                     EnvVariableScope = "BUILD"
	RuntimeVar                   EnvVariableScope = "RUNTIME"
	GlobalVar                    EnvVariableScope = "GLOBAL"
	InternalContainerRegistryVar EnvVariableScope = "INTERNAL_CONTAINER_REGISTRY"
	ContainerRegistryVar         EnvVariableScope = "CONTAINER_REGISTRY"
)

// TaskStatusType are the different deployment status types supported in Lagoon.
type TaskStatusType string

// These are the different deployment status types supported in Lagoon.
const (
	ActiveTask    TaskStatusType = "ACTIVE"
	SucceededTask TaskStatusType = "SUCCEEDED"
	FailedTask    TaskStatusType = "FAILED"
)

// RestoreStatusType are the different deployment status types supported in Lagoon.
type RestoreStatusType string

// These are the different deployment status types supported in Lagoon.
const (
	PendingRestore    RestoreStatusType = "PENDING"
	SuccessfulRestore RestoreStatusType = "SUCCESSFUL"
	FailedRestore     RestoreStatusType = "FAILED"
)

// EnvOrderType are the different deployment status types supported in Lagoon.
type EnvOrderType string

// These are the different deployment status types supported in Lagoon.
const (
	NameEnvOrder    EnvOrderType = "NAME"
	UpdatedEnvOrder EnvOrderType = "UPDATED"
)

// ProjectOrderType are the different deployment status types supported in Lagoon.
type ProjectOrderType string

// These are the different deployment status types supported in Lagoon.
const (
	NameProjectOrder    ProjectOrderType = "NAME"
	CreatedProjectOrder ProjectOrderType = "UPCREATEDDATED"
)

// GroupRole are the different deployment status types supported in Lagoon.
type GroupRole string

// These are the different deployment status types supported in Lagoon.
const (
	GuestRole      GroupRole = "GUEST"
	ReporterRole   GroupRole = "REPORTER"
	DeveloperRole  GroupRole = "DEVELOPER"
	MaintainerRole GroupRole = "MAINTAINER"
	OwnerRole      GroupRole = "OWNER"
)
