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

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var VpcAllowEmptyValues = []string{"tags."}

type VpcGenerator struct {
	AWSService
}

func (VpcGenerator) createResources(vpcs *ec2.DescribeVpcsOutput) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, vpc := range vpcs.Vpcs {
		resources = append(resources, terraformutils.NewSimpleResource(0, aws.StringValue(vpc.VpcId), aws.StringValue(vpc.VpcId), "aws_vpc", "aws", VpcAllowEmptyValues, ))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each vpc create 1 TerraformResource.
// Need VpcId as ID for terraform resource
func (g *VpcGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.New(config)
	p := ec2.NewDescribeVpcsPaginator(svc.DescribeVpcsRequest(&ec2.DescribeVpcsInput{}))
	for p.Next(context.Background()) {
		g.Resources = append(g.Resources, g.createResources(p.CurrentPage())...)
	}
	return p.Err()
}
