// Copyright 2022, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awss3exporter

import (
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestTranslateMetricsData(t *testing.T) {
	fmt.Println("translateMetricsData")
	md := pmetric.NewMetrics()
	md.ResourceMetrics().EnsureCapacity(2)
	rm := md.ResourceMetrics().AppendEmpty()
	ilms := rm.ScopeMetrics()
	ilms.EnsureCapacity(2)
	assert.NotNil(t, md, "failed to create metrics")
	
}
