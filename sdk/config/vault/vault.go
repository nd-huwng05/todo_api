package vault

import (
	"github.com/hashicorp/vault/api"
	"todo_api/sdk/config/yml"
)

type Vault struct {
	host         string
	port         string
	username     string
	password     string
	root_path    string
	service_name string
	client       *api.Client

	//Unit test
	yml        *yml.YamlConfig
	isTestMode bool
}

func (v *Vault) Initial(service_name string, arggs ...string) {

}
