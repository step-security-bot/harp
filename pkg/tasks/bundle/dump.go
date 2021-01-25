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

package bundle

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jmespath/go-jmespath"

	bundlev1 "github.com/elastic/harp/api/gen/go/harp/bundle/v1"
	"github.com/elastic/harp/pkg/bundle"
	"github.com/elastic/harp/pkg/sdk/types"
	"github.com/elastic/harp/pkg/tasks"
)

// DumpTask implements secret-container dumping task.
type DumpTask struct {
	ContainerReader tasks.ReaderProvider
	OutputWriter    tasks.WriterProvider
	PathOnly        bool
	DataOnly        bool
	MetadataOnly    bool
	JMESPathFilter  string
}

// Run the task.
func (t *DumpTask) Run(ctx context.Context) error {
	var (
		reader io.Reader
		err    error
	)

	// Create input reader
	reader, err = t.ContainerReader(ctx)
	if err != nil {
		return fmt.Errorf("unable to open input bundle: %w", err)
	}

	// Load bundle
	b, err := bundle.FromContainerReader(reader)
	if err != nil {
		return fmt.Errorf("unable to load bundle content: %w", err)
	}

	// Create output writer
	writer, err := t.OutputWriter(ctx)
	if err != nil {
		return fmt.Errorf("unable to open writer: %w", err)
	}

	if t.DataOnly {
		return t.dumpData(writer, b)
	}

	if t.MetadataOnly {
		return t.dumpMetadata(writer, b)
	}

	if t.PathOnly {
		return t.dumpPath(writer, b)
	}

	if t.JMESPathFilter != "" {
		return t.dumpFilter(writer, b)
	}

	// Dump full structure.
	if err := bundle.JSON(writer, b); err != nil {
		return fmt.Errorf("unable to generate JSON: %w", err)
	}

	// No error
	return nil
}

func (t *DumpTask) dumpData(writer io.Writer, b *bundlev1.Bundle) error {
	// Check arguments
	if types.IsNil(writer) {
		return fmt.Errorf("unable to process nil writer")
	}
	if b == nil {
		return fmt.Errorf("unable to process nil bundle")
	}

	// Convert bundle as a map
	bMap, err := bundle.AsMap(b)
	if err != nil {
		return fmt.Errorf("unable to convert bundle content: %w", err)
	}

	// Encode as JSON
	if err := json.NewEncoder(writer).Encode(bMap); err != nil {
		return fmt.Errorf("unable to marshal JSON bundle content: %w", err)
	}

	return nil
}

func (t *DumpTask) dumpMetadata(writer io.Writer, b *bundlev1.Bundle) error {
	// Check arguments
	if types.IsNil(writer) {
		return fmt.Errorf("unable to process nil writer")
	}
	if b == nil {
		return fmt.Errorf("unable to process nil bundle")
	}

	metaMap := bundle.KV{}

	// Export bundle metadata
	for _, p := range b.Packages {
		metadata := bundle.KV{}
		// Has annotations
		if len(p.Annotations) > 0 {
			out, err := json.Marshal(p.Annotations)
			if err != nil {
				return fmt.Errorf("unable to encode annotations as JSON: %w", err)
			}

			// Assign json
			metadata["harp.elastic.io/v1/bundle#annotations"] = string(out)
		}
		// Has labels
		if len(p.Labels) > 0 {
			out, err := json.Marshal(p.Labels)
			if err != nil {
				return fmt.Errorf("unable to encode labels as JSON: %w", err)
			}

			// Assign json
			metadata["harp.elastic.io/v1/bundle#labels"] = string(out)
		}

		// Assign to package
		metaMap[p.Name] = metadata
	}

	// Encode as JSON
	if err := json.NewEncoder(writer).Encode(metaMap); err != nil {
		return fmt.Errorf("unable to marshal JSON bundle metadata: %w", err)
	}

	return nil
}

func (t *DumpTask) dumpPath(writer io.Writer, b *bundlev1.Bundle) error {
	// Check arguments
	if types.IsNil(writer) {
		return fmt.Errorf("unable to process nil writer")
	}
	if b == nil {
		return fmt.Errorf("unable to process nil bundle")
	}

	// Extract bundle paths
	paths, err := bundle.Paths(b)
	if err != nil {
		return fmt.Errorf("unable to extract bundle paths: %w", err)
	}

	// Print a xargs compatible list
	for _, p := range paths {
		_, err = fmt.Fprintf(writer, "%s\n", p)
		if err != nil {
			return fmt.Errorf("unable to write package path '%s' to stdout: %w", p, err)
		}
	}

	return nil
}

func (t *DumpTask) dumpFilter(writer io.Writer, b *bundlev1.Bundle) error {
	// Check arguments
	if types.IsNil(writer) {
		return fmt.Errorf("unable to process nil writer")
	}
	if b == nil {
		return fmt.Errorf("unable to process nil bundle")
	}

	// Filter bundle with JMESPath expression
	res, err := jmespath.Search(t.JMESPathFilter, b)
	if err != nil {
		return fmt.Errorf("unable to process JMESPath filter '%s': %w", t.JMESPathFilter, err)
	}

	// Encode response
	out, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("unable to encode JMESPath filter result: %w", err)
	}

	// Write to writer
	if _, err := fmt.Fprintf(writer, "%s", string(out)); err != nil {
		return fmt.Errorf("unable to write JSON to stdout: %w", err)
	}

	return nil
}