// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package measurements

import "yunion.io/x/onecloud/pkg/apis/monitor"

var podCpu = SMeasurement{
	Context: []SMonitorContext{
		{
			Name:         "pod_cpu",
			DisplayName:  "Pod cpu",
			ResourceType: monitor.METRIC_RES_TYPE_CONTAINER,
			Database:     monitor.METRIC_DATABASE_TELE,
		},
	},
	Metrics: []SMetric{
		{
			Name:        "usage_rate",
			DisplayName: "Pod cpu usage rate",
			Unit:        monitor.METRIC_UNIT_PERCENT,
		},
	},
}

var podMem = SMeasurement{
	Context: []SMonitorContext{
		{
			Name:         "pod_mem",
			DisplayName:  "Pod memory",
			ResourceType: monitor.METRIC_RES_TYPE_CONTAINER,
			Database:     monitor.METRIC_DATABASE_TELE,
		},
	},
	Metrics: []SMetric{
		{
			Name:        "usage_rate",
			DisplayName: "Pod memory usage rate",
			Unit:        monitor.METRIC_UNIT_PERCENT,
		},
		{
			Name:        "working_set_bytes",
			DisplayName: "Pod memory working set bytes",
			Unit:        monitor.METRIC_UNIT_BYTE,
		},
	},
}

var podVolume = SMeasurement{
	Context: []SMonitorContext{
		{
			Name:         "pod_volume",
			DisplayName:  "Pod volume",
			ResourceType: monitor.METRIC_RES_TYPE_CONTAINER,
			Database:     monitor.METRIC_DATABASE_TELE,
		},
	},
	Metrics: []SMetric{
		{
			Name:        "free",
			DisplayName: "Pod volume free size",
			Unit:        monitor.METRIC_UNIT_BYTE,
		},
		{
			Name:        "used",
			DisplayName: "Pod volume used size",
			Unit:        monitor.METRIC_UNIT_BYTE,
		},
		{
			Name:        "used_percent",
			DisplayName: "Pod volume used percent",
			Unit:        monitor.METRIC_UNIT_PERCENT,
		},
		{
			Name:        "inodes_total",
			DisplayName: "Pod volume inodes total count",
			Unit:        monitor.METRIC_UNIT_COUNT,
		},
		{
			Name:        "inodes_free",
			DisplayName: "Pod volume inodes free count",
			Unit:        monitor.METRIC_UNIT_COUNT,
		},
		{
			Name:        "inodes_used",
			DisplayName: "Pod volume inodes used count",
			Unit:        monitor.METRIC_UNIT_COUNT,
		},
		{
			Name:        "inodes_used_percent",
			DisplayName: "Pod volume inodes used percent",
			Unit:        monitor.METRIC_UNIT_PERCENT,
		},
	},
}