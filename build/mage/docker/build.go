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

package docker

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	exec "golang.org/x/sys/execabs"

	"github.com/elastic/harp/build/artifact"
	"github.com/elastic/harp/build/mage/git"
)

var dockerTemplate = strings.TrimSpace(`
# syntax=docker/dockerfile:experimental

# Arguments
ARG BUILD_DATE={{.BuildDate}}
ARG VERSION={{.Version}}
ARG VCS_REF={{.VcsRef}}

# Builder arguments
ARG TOOLS_IMAGE={{.ToolImageName}}
ARG GOLANG_IMAGE={{.GolangImage}}
ARG GOLANG_VERSION={{.GolangVersion}}

FROM $TOOLS_IMAGE as compiler

# Back to project root
WORKDIR $GOPATH/src/workspace

# Copy go.mod
COPY --chown=golang:golang go.mod .
COPY --chown=golang:golang go.sum .
RUN go mod download

# Copy project go module
COPY --chown=golang:golang . .

{{ if .Cmd.HasModule }}
# Go to cmd
WORKDIR $GOPATH/src/workspace/{{ .Cmd.Module }}
{{ else }}
# Stay at root path
WORKDIR $GOPATH/src/workspace
{{ end }}

# Clean existing binaries
RUN set -eux; \
	rm -f bin/*

# Update vendor
RUN set -eux; \
	go mod vendor

# Build final target
RUN set -eux; \
    mage

# Compress binaries
RUN set -eux; \
    upx -9 bin/* && \
    chmod +x bin/*

## -------------------------------------------------------------------------------------------------

# hadolint ignore=DL3007
FROM gcr.io/distroless/static:latest

# Arguments
ARG BUILD_DATE={{.BuildDate}}
ARG VERSION={{.Version}}
ARG VCS_REF={{.VcsRef}}

# Metadata
LABEL \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="{{.Cmd.Name}}" \
    org.label-schema.description="{{.Cmd.Description}}" \
    org.label-schema.url="https://{{.Cmd.Package}}" \
    org.label-schema.vcs-url="https://{{.Cmd.Package}}.git" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vendor="Elastic" \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0"

{{ if .Cmd.HasModule }}
COPY --from=compiler /go/src/workspace/{{.Cmd.Module}}/bin/{{.Cmd.Kebab}}-linux-amd64 /usr/bin/{{.Cmd.Kebab}}
{{ else }}
COPY --from=compiler /go/src/workspace/bin/{{.Cmd.Kebab}}-linux-amd64 /usr/bin/{{.Cmd.Kebab}}
{{ end }}

COPY --from=compiler /tmp/group /tmp/passwd /etc/
COPY --from=compiler --chown=65534:65534 /tmp/.config /

USER nobody:nobody
WORKDIR /

ENTRYPOINT [ "/usr/bin/{{.Cmd.Kebab}}" ]
CMD ["--help"]
`)

// Build a docker container for given command.
func Build(cmd *artifact.Command) func() error {
	return func() error {
		mg.Deps(git.CollectInfo)

		// Retrieve golang attributes
		golangImage := golangImage
		if os.Getenv("GOLANG_IMAGE") != "" {
			golangImage = os.Getenv("GOLANG_IMAGE")
		}
		golangVersion := golangVersion
		if os.Getenv("GOLANG_VERSION") != "" {
			golangVersion = os.Getenv("GOLANG_VERSION")
		}

		// Docker image name
		toolImageName := toolImage
		if os.Getenv("TOOL_IMAGE_NAME") != "" {
			toolImageName = os.Getenv("TOOL_IMAGE_NAME")
		}

		buf, err := merge(dockerTemplate, map[string]interface{}{
			"ToolImageName": toolImageName,
			"BuildDate":     time.Now().Format(time.RFC3339),
			"Version":       git.Tag,
			"VcsRef":        git.Revision,
			"GolangImage":   golangImage,
			"GolangVersion": golangVersion,
			"Cmd":           cmd,
		})
		if err != nil {
			return err
		}

		// Check if we want to generate dockerfile output
		if os.Getenv("DOCKERFILE_ONLY") != "" {
			return os.WriteFile(fmt.Sprintf("Dockerfile.%s", cmd.Kebab()), buf.Bytes(), 0o600)
		}

		// Prepare command
		//nolint:gosec // expected behavior
		c := exec.Command("docker", "build",
			"-t", fmt.Sprintf("elastic/%s", cmd.Kebab()),
			"-f", "-",
			"--build-arg", fmt.Sprintf("BUILD_DATE=%s", time.Now().Format(time.RFC3339)),
			"--build-arg", fmt.Sprintf("VERSION=%s", git.Tag),
			"--build-arg", fmt.Sprintf("VCS_REF=%s", git.Revision),
			".",
		)

		// Prepare environment
		c.Env = os.Environ()
		c.Env = append(c.Env, "DOCKER_BUILDKIT=1")
		c.Stderr = os.Stderr
		c.Stdout = os.Stdout
		c.Stdin = buf // Pass Dockerfile as stdin

		// Execute
		err = c.Run()
		if err != nil {
			return fmt.Errorf("unable to run command: %w", err)
		}

		// Check execution error
		if !sh.CmdRan(err) {
			return fmt.Errorf("running '%s' failed with exit code %d", c.String(), sh.ExitStatus(err))
		}

		// No error
		return nil
	}
}
