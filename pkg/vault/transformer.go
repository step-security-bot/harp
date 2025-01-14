// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package vault

import (
	"fmt"
	"path"
	"strings"

	"github.com/zntrio/harp/v2/pkg/sdk/value"
	"github.com/zntrio/harp/v2/pkg/sdk/value/encryption"
	"github.com/zntrio/harp/v2/pkg/sdk/value/encryption/aead"
	"github.com/zntrio/harp/v2/pkg/sdk/value/encryption/envelope"
	"github.com/zntrio/harp/v2/pkg/sdk/value/encryption/secretbox"
	vaultpath "github.com/zntrio/harp/v2/pkg/vault/path"
)

type DataEncryption string

var (
	AESGCM           DataEncryption = "aesgcm"
	Chacha20Poly1305 DataEncryption = "chacha20poly1305"
	Secretbox        DataEncryption = "secretbox"
)

func init() {
	encryption.Register("vault", FromKey)
}

// Vault returns an envelope encryption using a remote transit backend for key
// encryption.
// vault:<path>:<data encryption>
func FromKey(key string) (value.Transformer, error) {
	// Remove the prefix
	key = strings.TrimPrefix(key, "vault:")

	// Split path / encryption
	parts := strings.SplitN(key, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("key format error, invalid part count")
	}

	// Split transit backend path
	mountPath, keyName := path.Split(parts[0])

	// Delegate to transformer
	return Transformer(mountPath, keyName, DataEncryption(parts[1]))
}

func TransformerKey(mountPath, keyName string, dataEncryption DataEncryption) string {
	return fmt.Sprintf("vault:%s/%s:%s", vaultpath.SanitizePath(mountPath), strings.TrimPrefix(keyName, "/"), dataEncryption)
}

// Transformer returns an envelope encryption using a remote transit backend for key
// encryption.
func Transformer(mountPath, keyName string, dataEncryption DataEncryption) (value.Transformer, error) {
	// Create default vault client
	client, err := DefaultClient()
	if err != nil {
		return nil, fmt.Errorf("unable to initialize vault client: %w", err)
	}

	// Create transit backend service
	backend, err := client.Transit(vaultpath.SanitizePath(mountPath), keyName)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize vault transit backend service: %w", err)
	}

	// Prepare data encryption
	var dataEncryptionFunc encryption.TransformerFactoryFunc
	switch dataEncryption {
	case AESGCM:
		dataEncryptionFunc = aead.AESGCM
	case Chacha20Poly1305:
		dataEncryptionFunc = aead.Chacha20Poly1305
	case Secretbox:
		dataEncryptionFunc = secretbox.Transformer
	default:
		return nil, fmt.Errorf("unsupported data encryption %q for envelope transformer", dataEncryption)
	}

	// Wrap the transformer with envelope
	return envelope.Transformer(backend, dataEncryptionFunc)
}
