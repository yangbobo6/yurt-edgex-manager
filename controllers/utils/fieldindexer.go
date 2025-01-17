/*
Copyright 2022 The OpenYurt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"context"
	"sync"

	"github.com/openyurtio/yurt-edgex-manager/api/v1alpha1"
	"github.com/openyurtio/yurt-edgex-manager/api/v1alpha2"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	IndexerPathForNodepool = "spec.poolname"
)

var registerOnce sync.Once

func RegisterFieldIndexers(fi client.FieldIndexer) error {
	var err error
	registerOnce.Do(func() {
		// register the fieldIndexer for device
		if err = fi.IndexField(context.TODO(), &v1alpha1.EdgeX{}, IndexerPathForNodepool, func(rawObj client.Object) []string {
			edgex, ok := rawObj.(*v1alpha1.EdgeX)
			if ok {
				return []string{edgex.Spec.PoolName}
			}
			return []string{}
		}); err != nil {
			return
		}
	})
	if err != nil {
		return err
	}
	registerOnce.Do(func() {
		// register the fieldIndexer for device
		if err = fi.IndexField(context.TODO(), &v1alpha2.EdgeX{}, IndexerPathForNodepool, func(rawObj client.Object) []string {
			edgex, ok := rawObj.(*v1alpha2.EdgeX)
			if ok {
				return []string{edgex.Spec.PoolName}
			}
			return []string{}
		}); err != nil {
			return
		}
	})
	return err
}
