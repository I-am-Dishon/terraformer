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

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
)

var mediapackageAllowEmptyValues = []string{"tags."}

type MediaPackageGenerator struct {
	AWSService
}

func (g *MediaPackageGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := mediapackage.New(config)
	p := mediapackage.NewListChannelsPaginator(svc.ListChannelsRequest(&mediapackage.ListChannelsInput{}))
	var resources []terraformutils.Resource
	for p.Next(context.Background()) {
		for _, channel := range p.CurrentPage().Channels {
			channelID := aws.StringValue(channel.Id)
			resources = append(resources, terraformutils.NewSimpleResource(0, channelID, channelID, "aws_media_package_channel", "aws", mediapackageAllowEmptyValues))
		}
	}
	g.Resources = resources
	return p.Err()
}
