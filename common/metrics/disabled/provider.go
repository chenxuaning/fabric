/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package disabled

import "github.com/hyperledger/fabric/common/metrics"

type Provider struct{}

func (p *Provider) NewCounter(opts metrics.CounterOpts) metrics.Counter { return &Counter{} }

func (p *Provider) NewGauge(opts metrics.GaugeOpts) metrics.Gauge { return &Gauge{} }

func (p *Provider) NewHistogram(opts metrics.HistogramOpts) metrics.Histogram { return &Histogram{} }

type Counter struct{}

func (c *Counter) With(labelValues ...string) metrics.Counter { return c }

func (c *Counter) Add(delta float64) {}

type Gauge struct{}

func (g *Gauge) With(labelValues ...string) metrics.Gauge { return g }

func (g *Gauge) Set(value float64) {}

type Histogram struct{}

func (h *Histogram) With(labelValues ...string) metrics.Histogram { return h }

func (h *Histogram) Observe(value float64) {}
