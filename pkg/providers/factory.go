package providers

import (
	"tentacle/pkg/utils"
	"tentacle/pkg/providers/cyberark"
	"tentacle/pkg/providers/thycotic"
	"tentacle/pkg/providers/os/darwin/keychain"
	"log"
)

func Create(alias string, config interface{}) (Interface, error) {
	//stringify config keys for all provider config entries.
	config = utils.StringifyYAMLMapKeys(config)

	//begin switch for providers.
	switch providerType := config.(map[string]interface{})["type"].(string); providerType {

	//os specific providers
	case "keychain":
		provider := new(keychain.Provider)
		provider.Init(alias, config.(map[string]interface{}))
		return provider, nil


	//alphabetical list of common providers
	case "cyberark":
		provider := new(cyberark.Provider)
		provider.Init(alias, config.(map[string]interface{}))
		return provider, nil
	case "thycotic":
		provider := new(thycotic.Provider)
		provider.Init(alias, config.(map[string]interface{}))
		return provider, nil

	//fall back error message
	default:
		log.Fatalf("%v type is not supported by tentacle", providerType)
		return nil, nil
	}
}