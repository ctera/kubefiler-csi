/*
Copyright 2021, CTERA Networks.

Portions Copyright 2019 The Kubernetes Authors.

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

package driver

import (
	"encoding/json"
	"errors"

	"github.com/pborman/uuid"
)

type KubeFilerVolumeID struct {
	ID                  string `json:"id"`
	Namespace           string `json:"namespace"`
	KubeFilerExportName string `json:"kubefiler_export_name"`
}

func NewKubeFilerVolumeID(namespace, kubeFilerExportName string) *KubeFilerVolumeID {
	ret := KubeFilerVolumeID{
		ID:                  uuid.NewUUID().String(),
		Namespace:           namespace,
		KubeFilerExportName: kubeFilerExportName,
	}
	return &ret
}

func (c *KubeFilerVolumeID) ToVolumeID() (*string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	ret := string(bytes)
	return &ret, nil
}

func getKubeFilerVolumeIDFromVolumeID(volumeID string) (*KubeFilerVolumeID, error) {
	if len(volumeID) == 0 {
		return nil, errors.New("volume ID missing in request")
	}

	var kubeFilerVolumeID KubeFilerVolumeID
	err := json.Unmarshal([]byte(volumeID), &kubeFilerVolumeID)
	if err != nil {
		return nil, err
	}
	return &kubeFilerVolumeID, nil
}
