// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package selInputOperator // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/stdin"

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/bougou/go-ipmi"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Input is an operator that reads input from stdin
type Input struct {
	helper.InputOperator
	wg                 sync.WaitGroup
	cancel             context.CancelFunc
	ipmiClient         *ipmi.Client
	collectionInterval time.Duration
	start              time.Time
	end                time.Time
}

// Start will start generating log entries.
func (i *Input) Start(_ operator.Persister) error {
	ctx, cancel := context.WithCancel(context.Background())
	i.cancel = cancel

	i.wg.Add(1)
	go func() {
		defer i.wg.Done()
		i.start = time.Now().Add(-i.collectionInterval)
		i.end = time.Now()
		for {
			selList, err := i.ipmiClient.GetSELEntries(ctx, 0)
			if err != nil {
				i.Logger().Error("failed to get SEL list", zap.Error(err))
				return
			}
			selListStr := ipmi.FormatSELs(selList, nil)

			e := entry.New()
			e.Body = selListStr
			err = i.Write(ctx, e)
			if err != nil {
				i.Logger().Error("failed to write entry", zap.Error(err))
				return
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(i.collectionInterval):
				i.start = i.end
				i.end = time.Now()
			}
		}
	}()

	return nil
}

// Stop will stop generating logs.
func (i *Input) Stop() error {
	i.cancel()
	i.wg.Wait()
	return nil
}
