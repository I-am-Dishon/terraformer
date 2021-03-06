// Copyright 2018 The Terraformer Authors.
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

package rabbitmq

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	RBTService
}

type User struct {
	Name string `json:"name"`
}

type Users []User

var UserAllowEmptyValues = []string{}

func (g UserGenerator) createResources(users Users) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range users {
		resources = append(resources, terraformutils.NewSimpleResource(0, user.Name, "user_"+normalizeResourceName(user.Name), "rabbitmq_user", "rabbitmq", UserAllowEmptyValues, ))
	}
	return resources
}

func (g *UserGenerator) InitResources() error {
	body, err := g.generateRequest("/api/users?columns=name")
	if err != nil {
		return err
	}
	var users Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(users)
	return nil
}
