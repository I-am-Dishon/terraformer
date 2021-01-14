package gitlab

import (
	"errors"
	"os"
	"net/url"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type GitlabProvider struct { //nolint
	terraformutils.Provider
	token string
	url string
}

func (p *GitlabProvider) Init(args []string) error {
	if os.Getenv("GITLAB_TOKEN") == "" {
		return errors.New("set GITLAB_TOKEN env var")
	}
	p.token = os.Getenv("GITLAB_TOKEN")

	gitlabUrl, err := url.Parse(args[1])
	
	if err != nil || gitlabUrl.Host == "" || gitlabUrl.Scheme == "" {
		return errors.New("give a valid URL")
	}
	p.url = gitlabUrl.String()

	return nil
}

func (p *GitlabProvider) GetName() string {
	return "gitlab"
}

func (p *GitlabProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
		},
	}
}


func (p *GitlabProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *GitlabProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"labels":                &LabelsGenerator{},
		"projects":        &ProjectsGenerator{},
		"groups":   &GroupsGenerator{},
		"gitlab_deploy":             &GitlabDeployGenerator{},
		"user":            &UserGenerator{},
		"pipelines":   &PipelinesGenerator{},
		"gitlab_service":           &GitlabServiceGenerator{},
		"protection":        &ProtectionGenerator{},
		"instance": &InstanceGenerator{},
	}
}

func (p *GitlabProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("gitlab: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"token": p.token,
		"url": p.url,
	})
	return nil
}
