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

package heroku

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	heroku "github.com/heroku/heroku-go/v5"
)

type PipelineCouplingGenerator struct {
	HerokuService
}

func (g PipelineCouplingGenerator) createResources(pipelineCouplingList []heroku.PipelineCoupling) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, pipelineCoupling := range pipelineCouplingList {
		resources = append(resources, terraformutils.NewSimpleResource(0, pipelineCoupling.ID, pipelineCoupling.ID, "heroku_pipeline_coupling", "heroku", []string{}))
	}
	return resources
}

func (g *PipelineCouplingGenerator) InitResources() error {
	svc := g.generateService()
	output, err := svc.PipelineCouplingList(context.TODO(), &heroku.ListRange{Field: "id"})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
