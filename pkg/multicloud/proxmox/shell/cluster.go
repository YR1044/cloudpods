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

package shell

import (
	"yunion.io/x/onecloud/pkg/multicloud/proxmox"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type ClusterListOptions struct {
	}
	shellutils.R(&ClusterListOptions{}, "cluster-list", "list clusters", func(cli *proxmox.SRegion, args *ClusterListOptions) error {
		clusters, err := cli.GetClusterAllResources()
		if err != nil {
			return err
		}
		printList(clusters, 0, 0, 0, []string{})
		return nil
	})

	shellutils.R(&ClusterListOptions{}, "cluster-node-list", "list nodes", func(cli *proxmox.SRegion, args *ClusterListOptions) error {
		clusters, err := cli.GetClusterNodeResources()
		if err != nil {
			return err
		}
		printList(clusters, 0, 0, 0, []string{})
		return nil
	})

	shellutils.R(&ClusterListOptions{}, "cluster-storge-list", "list storges", func(cli *proxmox.SRegion, args *ClusterListOptions) error {
		clusters, err := cli.GetClusterStoragesResources()
		if err != nil {
			return err
		}
		printList(clusters, 0, 0, 0, []string{})
		return nil
	})

}