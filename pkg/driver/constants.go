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

// constants of keys in secrets
const (
	FilerUsernameKey = "username"
	FilerPasswordKey = "password"
)

// constants of keys in volume parameters
const (
	// Hostname or IP address of the Ctera Filer
	FilerAddressKey = "fileraddress"
	// Path on the Ctera Filer to share
	PathKey = "path"

	// KmsKeyId represents key for KMS encryption key
	KmsKeyIDKey = "kmskeyid"
)

// constants for default command line flag values
const (
	DefaultCSIEndpoint = "unix://tmp/csi.sock"
)
