//go:generate mockgen -source=projects.go -destination ../mock/mock_projects.go -package mock
package lagoon_test

import (
	"context"
	"testing"

	"github.com/amazeeio/lagoon-client-go/mock"
	"github.com/amazeeio/lagoon-client-go/schema"
	"github.com/golang/mock/gomock"
)

type projectCalls struct {
	Project         []schema.Project
	ProjectMetadata []schema.ProjectMetadata
}

func TestGetMinimalProjectByName(t *testing.T) {
	var testCases = map[string]struct {
		expect *projectCalls
	}{
		"simple": {expect: &projectCalls{
			Project:         []schema.Project{},
			ProjectMetadata: []schema.ProjectMetadata{},
		}},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			ctx := context.Background()
			// set up the mock project
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()
			project := mock.NewMockProjects(ctrl)
			for i := range tc.expect.Project {
				project.EXPECT().MinimalProjectByName(ctx, "", &tc.expect.Project[i])
				project.EXPECT().NotificationsForProject(ctx, "", &tc.expect.Project[i], nil)
			}
			for i := range tc.expect.ProjectMetadata {
				project.EXPECT().ProjectsByMetadata(ctx, "", "", &tc.expect.ProjectMetadata[i])
				project.EXPECT().UpdateProjectMetadata(ctx, 0, "", "", &tc.expect.ProjectMetadata[i])
				project.EXPECT().RemoveProjectMetadataByKey(ctx, 0, "", &tc.expect.ProjectMetadata[i])
			}
		})
	}
}
