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

package v2

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"reflect"
	"testing"

	"github.com/awnumar/memguard"
	"github.com/stretchr/testify/assert"

	containerv1 "github.com/zntrio/harp/v2/api/gen/go/harp/container/v1"
	"github.com/zntrio/harp/v2/pkg/sdk/security/crypto/deterministicecdsa"
)

func Test_deriveSharedKeyFromRecipient(t *testing.T) {
	key1, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00001-deterministic-buffer-for-tests-26FBE7DED9E992BC36C06C988C1AC8A1E672B4B5959EF60672A983EFA7C8EE0F")))
	assert.NoError(t, err)
	assert.NotNil(t, key1)

	key2, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00002-deterministic-buffer-for-tests-37ACB0DD3A3CE5A0960CCE0F6A0D7E663DFFD221FBE8EEB03B20D3AD91BCDD55")))
	assert.NoError(t, err)
	assert.NotNil(t, key2)

	dk1, err := deriveSharedKeyFromRecipient(&key1.PublicKey, key2, nil)
	assert.NoError(t, err)
	assert.Equal(t, &[32]byte{
		0xfa, 0x88, 0x52, 0x30, 0x55, 0xe8, 0xd6, 0x8a,
		0xa8, 0x11, 0xa9, 0xf7, 0x92, 0x79, 0x2a, 0xe6,
		0x10, 0x12, 0xbd, 0x9d, 0xee, 0x98, 0x54, 0x9e,
		0x50, 0x25, 0xb3, 0xaa, 0x79, 0x77, 0xce, 0xd3,
	}, dk1)

	dk2, err := deriveSharedKeyFromRecipient(&key2.PublicKey, key1, nil)
	assert.NoError(t, err)
	assert.Equal(t, dk1, dk2)
}

func Test_deriveSharedKeyFromRecipient_WithPSK(t *testing.T) {
	var psk [preSharedKeySize]byte
	memguard.WipeBytes(psk[:]) // Ensure zeroed psk

	key1, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00001-deterministic-buffer-for-tests-26FBE7DED9E992BC36C06C988C1AC8A1E672B4B5959EF60672A983EFA7C8EE0F")))
	assert.NoError(t, err)
	assert.NotNil(t, key1)

	key2, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00002-deterministic-buffer-for-tests-37ACB0DD3A3CE5A0960CCE0F6A0D7E663DFFD221FBE8EEB03B20D3AD91BCDD55")))
	assert.NoError(t, err)
	assert.NotNil(t, key2)

	dk1, err := deriveSharedKeyFromRecipient(&key1.PublicKey, key2, &psk)
	assert.NoError(t, err)
	assert.Equal(t, &[32]byte{
		0x95, 0xd9, 0xca, 0xce, 0x6d, 0x6c, 0x83, 0x93,
		0xa0, 0x4d, 0x84, 0xe5, 0x36, 0x81, 0x62, 0xc3,
		0xa5, 0x86, 0x72, 0xba, 0xe1, 0xac, 0x2f, 0x6a,
		0xa2, 0xae, 0x50, 0x3, 0xe3, 0xd1, 0xb5, 0x1a,
	}, dk1)

	dk2, err := deriveSharedKeyFromRecipient(&key2.PublicKey, key1, &psk)
	assert.NoError(t, err)
	assert.Equal(t, dk1, dk2)
}

func Test_keyIdentifierFromDerivedKey(t *testing.T) {
	dk := &[32]byte{
		0x9f, 0x6c, 0xb8, 0x33, 0xf4, 0x7a, 0x4, 0xb2,
		0xba, 0x65, 0x30, 0xf2, 0xc, 0x7c, 0xb1, 0x30,
		0x22, 0xa0, 0x6a, 0x15, 0x57, 0x73, 0xc1, 0xa9,
		0xc7, 0x21, 0x48, 0xdd, 0x3c, 0xc8, 0x36, 0xc7,
	}

	id, err := keyIdentifierFromDerivedKey(dk, nil)
	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0xe0, 0xe9, 0xd5, 0xc5, 0x7a, 0x9e, 0x1c, 0x3,
		0x9d, 0x4b, 0xc0, 0x21, 0x6e, 0x72, 0x1a, 0xda,
		0xac, 0xd0, 0x57, 0xb8, 0x21, 0x16, 0x48, 0xac,
		0x46, 0x7c, 0x64, 0xf9, 0x4d, 0xe5, 0x86, 0x23,
	}, id)
}

func Test_keyIdentifierFromDerivedKey_WithPSK(t *testing.T) {
	var psk [preSharedKeySize]byte
	memguard.WipeBytes(psk[:]) // Ensure zeroed psk

	dk := &[32]byte{
		0x9f, 0x6c, 0xb8, 0x33, 0xf4, 0x7a, 0x4, 0xb2,
		0xba, 0x65, 0x30, 0xf2, 0xc, 0x7c, 0xb1, 0x30,
		0x22, 0xa0, 0x6a, 0x15, 0x57, 0x73, 0xc1, 0xa9,
		0xc7, 0x21, 0x48, 0xdd, 0x3c, 0xc8, 0x36, 0xc7,
	}

	id, err := keyIdentifierFromDerivedKey(dk, &psk)
	assert.NoError(t, err)
	assert.Equal(t, []byte{
		0x70, 0x5e, 0xe6, 0x3e, 0x5f, 0xf7, 0x78, 0x22,
		0xda, 0x79, 0x1c, 0xd6, 0x92, 0x53, 0xd8, 0x66,
		0xbb, 0xb, 0xcf, 0xf1, 0x86, 0x7b, 0x9e, 0x9b,
		0x7b, 0x1d, 0x9f, 0xaf, 0x65, 0x6d, 0xa5, 0x7f,
	}, id)
}

func Test_packRecipient(t *testing.T) {
	payloadKey := &[32]byte{}

	key1, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00001-deterministic-buffer-for-tests-26FBE7DED9E992BC36C06C988C1AC8A1E672B4B5959EF60672A983EFA7C8EE0F")))
	assert.NoError(t, err)
	assert.NotNil(t, key1)

	key2, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00002-deterministic-buffer-for-tests-37ACB0DD3A3CE5A0960CCE0F6A0D7E663DFFD221FBE8EEB03B20D3AD91BCDD55")))
	assert.NoError(t, err)
	assert.NotNil(t, key2)

	recipient, err := packRecipient(rand.Reader, payloadKey, key1, &key2.PublicKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, recipient)
	assert.Equal(t, []byte{
		0xaa, 0xc5, 0x2b, 0x2e, 0xdf, 0x44, 0x9e, 0x87,
		0xc3, 0xc9, 0x9a, 0x98, 0xb1, 0xad, 0x7a, 0xcd,
		0x70, 0xe9, 0xa1, 0x56, 0xf6, 0xd5, 0x87, 0xb8,
		0x25, 0x94, 0x18, 0x3f, 0xf7, 0x8e, 0xdc, 0x46,
	}, recipient.Identifier)
	assert.Equal(t, seedSize+encryptionKeySize+macSize, len(recipient.Key))
}

func Test_packRecipient_WithPSK(t *testing.T) {
	var psk [preSharedKeySize]byte
	memguard.WipeBytes(psk[:]) // Ensure zeroed psk

	payloadKey := &[32]byte{}

	key1, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00001-deterministic-buffer-for-tests-26FBE7DED9E992BC36C06C988C1AC8A1E672B4B5959EF60672A983EFA7C8EE0F")))
	assert.NoError(t, err)
	assert.NotNil(t, key1)

	key2, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00002-deterministic-buffer-for-tests-37ACB0DD3A3CE5A0960CCE0F6A0D7E663DFFD221FBE8EEB03B20D3AD91BCDD55")))
	assert.NoError(t, err)
	assert.NotNil(t, key2)

	recipient, err := packRecipient(rand.Reader, payloadKey, key1, &key2.PublicKey, &psk)
	assert.NoError(t, err)
	assert.NotNil(t, recipient)
	assert.Equal(t, []byte{
		0xe0, 0x83, 0xf2, 0x13, 0xea, 0x5f, 0x48, 0x2b,
		0xa1, 0x4e, 0x77, 0x55, 0xa9, 0x5b, 0x58, 0xdc,
		0x2c, 0xb3, 0x11, 0xcb, 0xbc, 0xf8, 0xfd, 0x2c,
		0xca, 0xb7, 0x9, 0x41, 0xc9, 0x28, 0xb1, 0x29,
	}, recipient.Identifier)
	assert.Equal(t, seedSize+encryptionKeySize+macSize, len(recipient.Key))
}

func Test_tryRecipientKeys(t *testing.T) {
	payloadKey := &[32]byte{}

	key1, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00001-deterministic-buffer-for-tests-26FBE7DED9E992BC36C06C988C1AC8A1E672B4B5959EF60672A983EFA7C8EE0F")))
	assert.NoError(t, err)
	assert.NotNil(t, key1)

	key2, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00002-deterministic-buffer-for-tests-37ACB0DD3A3CE5A0960CCE0F6A0D7E663DFFD221FBE8EEB03B20D3AD91BCDD55")))
	assert.NoError(t, err)
	assert.NotNil(t, key2)

	recipient, err := packRecipient(rand.Reader, payloadKey, key1, &key2.PublicKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, recipient)
	assert.Equal(t, []byte{
		0xaa, 0xc5, 0x2b, 0x2e, 0xdf, 0x44, 0x9e, 0x87,
		0xc3, 0xc9, 0x9a, 0x98, 0xb1, 0xad, 0x7a, 0xcd,
		0x70, 0xe9, 0xa1, 0x56, 0xf6, 0xd5, 0x87, 0xb8,
		0x25, 0x94, 0x18, 0x3f, 0xf7, 0x8e, 0xdc, 0x46,
	}, recipient.Identifier)
	assert.Equal(t, seedSize+encryptionKeySize+macSize, len(recipient.Key))

	// -------------------------------------------------------------------------
	dk, err := deriveSharedKeyFromRecipient(&key1.PublicKey, key2, nil)
	assert.NoError(t, err)
	assert.Equal(t, &[32]byte{
		0xfa, 0x88, 0x52, 0x30, 0x55, 0xe8, 0xd6, 0x8a,
		0xa8, 0x11, 0xa9, 0xf7, 0x92, 0x79, 0x2a, 0xe6,
		0x10, 0x12, 0xbd, 0x9d, 0xee, 0x98, 0x54, 0x9e,
		0x50, 0x25, 0xb3, 0xaa, 0x79, 0x77, 0xce, 0xd3,
	}, dk)

	expectedID := []byte{
		0xaa, 0xc5, 0x2b, 0x2e, 0xdf, 0x44, 0x9e, 0x87,
		0xc3, 0xc9, 0x9a, 0x98, 0xb1, 0xad, 0x7a, 0xcd,
		0x70, 0xe9, 0xa1, 0x56, 0xf6, 0xd5, 0x87, 0xb8,
		0x25, 0x94, 0x18, 0x3f, 0xf7, 0x8e, 0xdc, 0x46,
	}
	id, err := keyIdentifierFromDerivedKey(dk, nil)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, id)
	assert.Equal(t, expectedID, recipient.Identifier)

	decodedPayloadKey, err := tryRecipientKeys(dk, []*containerv1.Recipient{
		recipient,
	}, nil)
	assert.NoError(t, err)
	assert.Equal(t, payloadKey, decodedPayloadKey)
}

func Test_tryRecipientKeys_WithPSK(t *testing.T) {
	var psk [preSharedKeySize]byte
	memguard.WipeBytes(psk[:]) // Ensure zeroed psk

	payloadKey := &[32]byte{}

	key1, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00001-deterministic-buffer-for-tests-26FBE7DED9E992BC36C06C988C1AC8A1E672B4B5959EF60672A983EFA7C8EE0F")))
	assert.NoError(t, err)
	assert.NotNil(t, key1)

	key2, err := deterministicecdsa.GenerateKey(elliptic.P384(), bytes.NewReader([]byte("00002-deterministic-buffer-for-tests-37ACB0DD3A3CE5A0960CCE0F6A0D7E663DFFD221FBE8EEB03B20D3AD91BCDD55")))
	assert.NoError(t, err)
	assert.NotNil(t, key2)

	recipient, err := packRecipient(rand.Reader, payloadKey, key1, &key2.PublicKey, &psk)
	assert.NoError(t, err)
	assert.NotNil(t, recipient)
	assert.Equal(t, []byte{
		0xe0, 0x83, 0xf2, 0x13, 0xea, 0x5f, 0x48, 0x2b,
		0xa1, 0x4e, 0x77, 0x55, 0xa9, 0x5b, 0x58, 0xdc,
		0x2c, 0xb3, 0x11, 0xcb, 0xbc, 0xf8, 0xfd, 0x2c,
		0xca, 0xb7, 0x9, 0x41, 0xc9, 0x28, 0xb1, 0x29,
	}, recipient.Identifier)
	assert.Equal(t, seedSize+encryptionKeySize+macSize, len(recipient.Key))

	// -------------------------------------------------------------------------
	dk, err := deriveSharedKeyFromRecipient(&key1.PublicKey, key2, &psk)
	assert.NoError(t, err)
	assert.Equal(t, &[32]byte{
		0x95, 0xd9, 0xca, 0xce, 0x6d, 0x6c, 0x83, 0x93,
		0xa0, 0x4d, 0x84, 0xe5, 0x36, 0x81, 0x62, 0xc3,
		0xa5, 0x86, 0x72, 0xba, 0xe1, 0xac, 0x2f, 0x6a,
		0xa2, 0xae, 0x50, 0x3, 0xe3, 0xd1, 0xb5, 0x1a,
	}, dk)

	expectedID := []byte{
		0xe0, 0x83, 0xf2, 0x13, 0xea, 0x5f, 0x48, 0x2b,
		0xa1, 0x4e, 0x77, 0x55, 0xa9, 0x5b, 0x58, 0xdc,
		0x2c, 0xb3, 0x11, 0xcb, 0xbc, 0xf8, 0xfd, 0x2c,
		0xca, 0xb7, 0x9, 0x41, 0xc9, 0x28, 0xb1, 0x29,
	}
	id, err := keyIdentifierFromDerivedKey(dk, &psk)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, id)
	assert.Equal(t, expectedID, recipient.Identifier)

	decodedPayloadKey, err := tryRecipientKeys(dk, []*containerv1.Recipient{
		recipient,
	}, &psk)
	assert.NoError(t, err)
	assert.Equal(t, payloadKey, decodedPayloadKey)
}

func TestPreAuthenticationEncoding(t *testing.T) {
	type args struct {
		pieces [][]byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				pieces: nil,
			},
			wantErr: false,
			want:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			name: "one",
			args: args{
				pieces: [][]byte{
					[]byte("test"),
				},
			},
			wantErr: false,
			want: []byte{
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Count
				0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Length
				't', 'e', 's', 't',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pae(tt.args.pieces...)
			if (err != nil) != tt.wantErr {
				t.Errorf("pae() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pae() = %v, want %v", got, tt.want)
			}
		})
	}
}
