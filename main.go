package main

import (
	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the mutator plugin config.
type Config struct {
	sensu.PluginConfig
}

var (
	mutatorConfig = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-registration-mutator",
			Short:    "Mutates agent registration events to include details of the host.",
			Keyspace: "sensu.io/plugins/sensu-registration-mutator/config",
		},
	}
)

func main() {
	mutator := sensu.NewGoMutator(&mutatorConfig.PluginConfig, nil, checkArgs, executeMutator)
	mutator.Execute()
}

func checkArgs(_ *types.Event) error {
	return nil
}

func executeMutator(event *types.Event) (*types.Event, error) {
	
	if(event.Check.Name == "registration"){
	        event.Check.Status = 0
		event.Check.Output = "New host registration."
	} else if (event.Check.Name == "deregistration"){
                event.Check.Status = 0
                event.Check.Output = "Host automatically removed."
	}

	return event, nil
}
