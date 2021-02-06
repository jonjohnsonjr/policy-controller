/*
Copyright The Cosign Authors.

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

package pkg

import (
	"encoding/json"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

func Payload(img v1.Descriptor) ([]byte, error) {

	simpleSigning := SimpleSigning{
		Critical: Critical{
			Image: Image{
				DockerManifestDigest: img.Digest.Hex,
			},
			Type: "cosign container signature",
		},
	}

	b, err := json.Marshal(simpleSigning)
	if err != nil {
		return nil, err
	}
	return b, err
}