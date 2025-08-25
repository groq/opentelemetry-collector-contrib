// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package selInputOperator // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/stdin"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/bougou/go-ipmi"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "sel_input"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig("") })
}

func NewDefaultConfig() *Config {
	return &Config{
		InputConfig: helper.NewInputConfig("sel_input", operatorType),
	}
}

// NewConfig creates a new stdin input config with default values
func NewConfig(operatorID string) *Config {
	return &Config{
		InputConfig: helper.NewInputConfig(operatorID, operatorType),
	}
}

// Config is the configuration of a stdin input operator.
type Config struct {
	helper.InputConfig `mapstructure:",squash"`

	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// Build will build a stdin input operator.
func (c *Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	inputOperator, err := c.InputConfig.Build(set)
	if err != nil {
		return nil, err
	}

	ipmiClient, err := ipmi.NewClient(c.Host, 80, c.User, c.Password)
	if err != nil {
		return nil, err
	}

	return &Input{
		InputOperator: inputOperator,
		ipmiClient:    ipmiClient,
	}, nil
}
