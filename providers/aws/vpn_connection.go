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

var VpnConnectionAllowEmptyValues = []string{"tags."}

type VpnConnectionGenerator struct {
	AWSService
}

func (VpnConnectionGenerator) createResources(vpncs *ec2.DescribeVpnConnectionsResponse) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, vpnc := range vpncs.VpnConnections {
		resources = append(resources, terraformutils.NewSimpleResource(0, aws.StringValue(vpnc.ConnectionId), aws.StringValue(vpnc.ConnectionId), "aws_vpn_connection", "aws", VpnConnectionAllowEmptyValues, ))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each vpn connection create 1 TerraformResource.
// Need VpnConnectionId as ID for terraform resource
func (g *VpnConnectionGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.New(config)
	vpncs, err := svc.DescribeVpnConnectionsRequest(&ec2.DescribeVpnConnectionsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	g.Resources = g.createResources(vpncs)
	return nil
}
