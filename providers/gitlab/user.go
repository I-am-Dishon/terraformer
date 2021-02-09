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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/xanzy/go-gitlab"
	"log"
)

type UserGenerator struct {
	GitlabService
}

var (
	// UserAllowEmptyValues ...
	UserAllowEmptyValues = []string{}
)

func (u UserGenerator) createResources(userList []*gitlab.User) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range userList {
		log.Println("Userid", user.ID, "---N", user.Name,"---u", user.Username,"---A", "---S", user.State,
			user.AvatarURL,"---WE", user.WebURL)

		resources = append(resources, terraformutils.NewSimpleResource(
			0,
			string(user),
			user.Username,
			"gitlab_user",
			"gitlab",
			UserAllowEmptyValues,
		))
			
	}
	return resources
}

func (u *UserGenerator) InitResources() error {

	git, err := gitlab.NewClient(u.Args["token"].(string), gitlab.WithBaseURL(u.Args["url"].(string)))
	if err != nil {
	log.Fatalf("Failed to create client: %v", err)
	}
	opt := &gitlab.ListUsersOptions{
		ListOptions: gitlab.ListOptions{PerPage: 5},
	}
	//users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	users, _, err := git.Users.ListUsers(opt)
	//err = json.Unmarshal(users, *gitlab.User)
	//if err != nil {
	//	return err
	//}
	u.Resources = u.createResources(users)
	return nil	
}

