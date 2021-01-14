// Copyright 2020 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gitlab

import (
	"log"
	
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gitlab "github.com/xanzy/go-gitlab"
	
)

type UserGenerator struct {
	GitlabService
}


func (u UserGenerator) createResources(userList []*gitlab.User) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range userList {
		resources = append(resources, terraformutils.NewSimpleResource(
			user.Username,
			user.Username,
			user.Name,
			user.State,
			[]string{},
			// "gitlab_user",
		))
			
	}
	return resources
}
// {
//     "id": 1,
//     "username": "john_smith",
//     "name": "John Smith",
//     "state": "active",
//     "avatar_url": "http://localhost:3000/uploads/user/avatar/1/cd8.jpeg",
//     "web_url": "http://localhost:3000/john_smith"
//   }

func (u *UserGenerator) InitResources() error {

	git, err := gitlab.NewClient(u.Args["token"].(string), gitlab.WithBaseURL(u.Args["url"].(string)))
	if err != nil {
	log.Fatalf("Failed to create client: %v", err)
	}
	opt := &gitlab.ListUsersOptions{
		ListOptions: gitlab.ListOptions{PerPage: 100},
	}
	//users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	users, _, err := git.Users.ListUsers(opt)
	u.Resources = u.createResources(users)
	return nil	
}