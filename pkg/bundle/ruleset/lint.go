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

package ruleset

import (
	"errors"
	"fmt"
	"io"

	"github.com/xeipuuv/gojsonschema"

	"github.com/zntrio/harp/api/jsonschema"
	"github.com/zntrio/harp/pkg/sdk/convert"
	"github.com/zntrio/harp/pkg/sdk/types"
)

// JSONSchema returns the used json schema for validation.
func JSONSchema() []byte {
	return jsonschema.BundleV1RuleSetSchema()
}

// Lint to input reader content with Bundle jsonschema.
func Lint(r io.Reader) ([]gojsonschema.ResultError, error) {
	// Check arguments
	if types.IsNil(r) {
		return nil, fmt.Errorf("reader is nil")
	}

	// Drain the reader
	jsonReader, err := convert.YAMLtoJSON(r)
	if err != nil {
		return nil, fmt.Errorf("unable to parse input as YAML: %w", err)
	}

	// Drain reader
	jsonData, err := io.ReadAll(jsonReader)
	if err != nil {
		return nil, fmt.Errorf("unable to drain all json reader content: %w", err)
	}

	// Prepare loaders
	schemaLoader := gojsonschema.NewBytesLoader(jsonschema.BundleV1RuleSetSchema())
	documentLoader := gojsonschema.NewBytesLoader(jsonData)

	// Validate
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, fmt.Errorf("ruleset validation failed %w", err)
	}
	if !result.Valid() {
		return result.Errors(), errors.New("ruleset not valid")
	}

	// No error
	return nil, nil
}
