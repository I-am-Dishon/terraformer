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
	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

var qldbAllowEmptyValues = []string{"tags."}

type QLDBGenerator struct {
	AWSService
}

func (g *QLDBGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := qldb.New(config)
	p := qldb.NewListLedgersPaginator(svc.ListLedgersRequest(&qldb.ListLedgersInput{}))
	var resources []terraformutils.Resource
	for p.Next(context.Background()) {
		for _, ledger := range p.CurrentPage().Ledgers {
			ledgerName := aws.StringValue(ledger.Name)
			resources = append(resources, terraformutils.NewSimpleResource(0, ledgerName, ledgerName, "aws_qldb_ledger", "aws", qldbAllowEmptyValues))
		}
	}
	g.Resources = resources
	return p.Err()
}
