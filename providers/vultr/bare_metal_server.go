// Copyright 2019 The Terraformer Authors.
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

package vultr

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/vultr/govultr"
)

type BareMetalServerGenerator struct {
	VultrService
}

func (g BareMetalServerGenerator) createResources(serverList []govultr.BareMetalServer) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, server := range serverList {
		resources = append(resources, terraformutils.NewSimpleResource(0, server.BareMetalServerID, server.BareMetalServerID, "vultr_bare_metal_server", "vultr", []string{}))
	}
	return resources
}

func (g *BareMetalServerGenerator) InitResources() error {
	client := g.generateClient()
	output, err := client.BareMetalServer.List(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
