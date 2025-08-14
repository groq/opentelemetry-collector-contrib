//go:build !linux

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package edacscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/edacscraper"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"
)

// newEDACMetricsScraper creates an EDAC Scraper (non-Linux stub)
func newEDACMetricsScraper(_ context.Context, _ scraper.Settings, _ *Config) *edacScraper {
	return &edacScraper{}
}

type edacScraper struct{}

func (s *edacScraper) start(_ context.Context, _ component.Host) error {
	return errors.New("EDAC scraper is only supported on Linux")
}

func (s *edacScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	return pmetric.NewMetrics(), errors.New("EDAC scraper is only supported on Linux")
}
