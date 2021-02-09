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

type PipelineGenerator struct {
	HerokuService
}

func (g PipelineGenerator) createResources(pipelineList []heroku.Pipeline) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, pipeline := range pipelineList {
		resources = append(resources, terraformutils.NewSimpleResource(0, pipeline.ID, pipeline.Name, "heroku_pipeline", "heroku", []string{}))
	}
	return resources
}

func (g *PipelineGenerator) InitResources() error {
	svc := g.generateService()
	output, err := svc.PipelineList(context.TODO(), &heroku.ListRange{Field: "id"})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
