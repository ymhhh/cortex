// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package promql

import (
	"context"
	"testing"
	"time"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/stretchr/testify/assert"
)

func TestDeriv(t *testing.T) {
	// https://github.com/prometheus/prometheus/issues/2674#issuecomment-315439393
	// This requires more precision than the usual test system offers,
	// so we test it by hand.
	storage := NewStorage(t)
	defer storage.Close()
	engine := NewEngine(nil, nil, 10, 10*time.Second)

	a := storage.Appender(context.Background())

	metric := labels.FromStrings("__name__", "foo")
	_, err := a.Append(0, metric, 1493712816939, 1.0)
	assert.NoError(t, err)

	_, err = a.Append(0, metric, 1493712846939, 1.0)
	assert.NoError(t, err)

	err = a.Commit()
	assert.NoError(t, err)

	query, err := engine.NewInstantQuery(storage, "deriv(foo[30m])", timestamp.Time(1493712846939))
	assert.NoError(t, err)

	result := query.Exec(context.Background())
	assert.NoError(t, result.Err)

	vec, _ := result.Vector()
	assert.True(t, len(vec) == 1, "Expected 1 result, got %d", len(vec))
	assert.True(t, vec[0].V == 0.0, "Expected 0.0 as value, got %f", vec[0].V)
}
