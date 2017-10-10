// Copyright (c) 2017 Uber Technologies, Inc.
//
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

package dbmodel

import (
	"testing"

	"github.com/uber/jaeger/model"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestFilterLogTags(t *testing.T) {
	expectedTags := model.KeyValues{
		model.String(someStringTagKey, someStringTagValue),
		model.Bool(someBoolTagKey, someBoolTagValue),
		model.Int64(someLongTagKey, someLongTagValue),
	}
	testSpan := getTestJaegerSpan()
	testSpan.Tags = expectedTags
	testSpan.Process.Tags = model.KeyValues{}
	testSpan.Logs = []model.Log{
		{
			Timestamp: someLogTimestamp,
			Fields: model.KeyValues{
				model.Float64(someDoubleTagKey, someDoubleTagValue),
			},
		},
	}
	uniqueTags := FilterLogTags()(testSpan)
	if !assert.EqualValues(t, expectedTags, uniqueTags) {
		for _, diff := range pretty.Diff(expectedTags, uniqueTags) {
			t.Log(diff)
		}
	}
}
