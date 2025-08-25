package selreceiver

import (
	"github.com/groq/opentelemetry-collector-contrib/receiver/selreceiver/internal/metadata"
	selInputOperator "github.com/groq/opentelemetry-collector-contrib/receiver/selreceiver/internal/selinputoperator"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// ReceiverType implements adapter.LogReceiverType
// to create a syslog receiver
type ReceiverType struct{}

// Type is the receiver type
func (ReceiverType) Type() component.Type {
	return metadata.Type
}

// CreateDefaultConfig creates a config with type and version
func (ReceiverType) CreateDefaultConfig() component.Config {
	return &SelConfig{
		BaseConfig: adapter.BaseConfig{
			Operators: []operator.Config{},
		},
		InputConfig: *selInputOperator.NewDefaultConfig(),
	}
}

// BaseConfig gets the base config from config, for now
func (ReceiverType) BaseConfig(cfg component.Config) adapter.BaseConfig {
	return cfg.(*SelConfig).BaseConfig
}

// SelConfig defines configuration for the sel receiver
type SelConfig struct {
	InputConfig        selInputOperator.Config `mapstructure:",squash"`
	adapter.BaseConfig `mapstructure:",squash"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// InputConfig unmarshals the input operator
func (ReceiverType) InputConfig(cfg component.Config) operator.Config {
	return operator.NewConfig(&cfg.(*SelConfig).InputConfig)
}

func (cfg *SelConfig) Unmarshal(componentParser *confmap.Conf) error {
	if componentParser == nil {
		// Nothing to do if there is no config given.
		return nil
	}

	return componentParser.Unmarshal(cfg)
}
