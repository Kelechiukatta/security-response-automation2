package closebucket

// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"context"

	"github.com/googlecloudplatform/security-response-automation/services"
	"github.com/pkg/errors"
)

// publicUsers contains a slice of public users we want to remove.
var publicUsers = []string{"allUsers", "allAuthenticatedUsers"}

// Values contains the required values needed for this function.
type Values struct {
	BucketName string
	ProjectID  string
	DryRun     bool
}

// Services contains the services needed for this function.
type Services struct {
	Resource *services.Resource
	Logger   *services.Logger
}

// Execute will remove any public users from buckets found within the provided folders.
func Execute(ctx context.Context, values *Values, services *Services) error {
	if values.DryRun {
		services.Logger.Info("dry_run on, would have removed public members from bucket %q in project %q", values.BucketName, values.ProjectID)
		return nil
	}
	if err := services.Resource.RemoveMembersFromBucket(ctx, values.BucketName, publicUsers); err != nil {
		return errors.Wrapf(err, "failed while performing Remove Members Fro mBucket on %+v", values)
	}
	services.Logger.Info("removed public members from bucket %q in project %q", values.BucketName, values.ProjectID)
	return nil
}
