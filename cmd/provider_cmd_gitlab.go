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
package cmd

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
	"log"
    gitlab_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/gitlab"
)

func newCmdGitlabImporter(options ImportOptions) *cobra.Command {
	token := ""
	url := ""
	cmd := &cobra.Command{
		Use:   "gitlab",
		Short: "Import current state to Terraform configuration from Gitlab",
		Long:  "Import current state to Terraform configuration from Gitlab",
		RunE: func(cmd *cobra.Command, args []string) error {
			// originalPathPattern := options.PathPattern
			// for _, organization := range organizations {
				provider := newGitlabProvider()
			// 	options.PathPattern = originalPathPattern
			// 	// options.PathPattern = strings.ReplaceAll(options.PathPattern, "{provider}", "{provider}/"+organization)
			log.Println(provider.GetName() + " importing gitlan")
				err := Import(provider, options, []string{url, token})
				if err != nil {
					return err
				}
			// }

			return nil
		},
	}


	cmd.AddCommand(listCmd(newGitlabProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "repository", "repository=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "YOUR_GITLAB_TOKEN or env param GITLAB_TOKEN")
	cmd.PersistentFlags().StringVarP(&url, "GITLAB_BASE_URL", "", "", "GITLAB_BASE_URL=http://localhost:8929/api/v4/")
	return cmd
}

func newGitlabProvider() terraformutils.ProviderGenerator {
	return &gitlab_terraforming.GitlabProvider{}
}
