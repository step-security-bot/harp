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

package release

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/zntrio/harp/build/artifact"
)

var formulaTemplate = strings.TrimSpace(`# typed: false
# frozen_string_literal: true

# Code generated by Harp build tool
class {{ .Formula }} < Formula
  desc "{{ .Description }}"
  homepage "https://{{ .Repository }}"
  license "Apache-2.0"
  stable do
    on_macos do
      if Hardware::CPU.intel?
        url "https://{{ .Repository }}/releases/download/cmd%2F{{ .Bin }}%2F{{ .Release }}/{{ .Bin }}-darwin-amd64-{{ .Release }}.tar.xz"
        sha256 "{{ sha256file (printf "dist/%s-darwin-amd64-%s.tar.xz" .Bin .Release) }}"
      elsif Hardware::CPU.arm?
        url "https://{{ .Repository }}/releases/download/cmd%2F{{ .Bin }}%2F{{ .Release }}/{{ .Bin }}-darwin-arm64-{{ .Release }}.tar.xz"
        sha256 "{{ sha256file (printf "dist/%s-darwin-arm64-%s.tar.xz" .Bin .Release) }}"
      end
    end
    on_linux do
      if Hardware::CPU.intel?
        if Hardware::CPU.is_64_bit?
          url "https://{{ .Repository }}/releases/download/cmd%2F{{ .Bin }}%2F{{ .Release }}/{{ .Bin }}-linux-amd64-{{ .Release }}.tar.xz"
          sha256 "{{ sha256file (printf "dist/%s-linux-amd64-%s.tar.xz" .Bin .Release) }}"
        end
      elsif Hardware::CPU.arm?
        if Hardware::CPU.is_64_bit?
          url "https://{{ .Repository }}/releases/download/cmd%2F{{ .Bin }}%2F{{ .Release }}/{{ .Bin }}-linux-arm64-{{ .Release }}.tar.xz"
          sha256 "{{ sha256file (printf "dist/%s-linux-arm64-%s.tar.xz" .Bin .Release) }}"
        else
          url "https://{{ .Repository }}/releases/download/cmd%2F{{ .Bin }}%2F{{ .Release }}/{{ .Bin }}-linux-arm7-{{ .Release }}.tar.xz"
          sha256 "{{ sha256file (printf "dist/%s-linux-arm7-%s.tar.xz" .Bin .Release) }}"
        end
      end
    end
  end

  # Source definition
  head do
    url "https://{{ .Repository }}.git", :branch => "main"

    # build dependencies
    depends_on "go" => :build
    depends_on "mage" => :build
  end

  def install
    ENV.deparallelize

    if build.head?
      # Prepare build environment
      ENV["GOPATH"] = buildpath
      (buildpath/"src/{{ .Repository }}").install Dir["{*,.git,.gitignore}"]

      # Mage tools
      ENV.prepend_path "PATH", buildpath/"tools/bin"

      # In {{ .Repository }} command
      cd "src/{{ .Repository }}/cmd/{{ .Bin }}" do
        system "go", "mod", "vendor"
        system "mage", "compile"
      end

      # Install builded command
      cd "src/{{ .Repository }}/cmd/{{ .Bin }}/bin" do
        # Install binaries
        if OS.mac? && Hardware::CPU.arm?
          bin.install "{{ .Bin }}-darwin-arm64" => "{{ .Bin }}"
        elsif OS.mac?
          bin.install "{{ .Bin }}-darwin-amd64" => "{{ .Bin }}"
        elsif OS.linux?
          bin.install "{{ .Bin }}-linux-amd64" => "{{ .Bin }}"
        end
      end
    elsif OS.mac? && Hardware::CPU.arm?
      # Install binaries
      bin.install "{{ .Bin }}-darwin-arm64" => "{{ .Bin }}"
    elsif OS.mac?
      bin.install "{{ .Bin }}-darwin-amd64" => "{{ .Bin }}"
    elsif OS.linux?
      bin.install "{{ .Bin }}-linux-amd64" => "{{ .Bin }}"
    end

    # Final message
    ohai "Install success!"
  end

  def caveats
    <<~EOS
      If you have previously built {{ .Bin }} from source, make sure you're not using
      $GOPATH/bin/{{ .Bin }} instead of this one. If that's the case you can simply run
      rm -f $GOPATH/bin/{{ .Bin }}
    EOS
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/{{ .Bin }} version")
  end
end
`)

type formulaModel struct {
	Repository  string
	Bin         string
	Formula     string
	Description string
	Release     string
}

// HomebrewFormula generates HomeBrew formula for given command.
func HomebrewFormula(cmd *artifact.Command) func() error {
	sha256sum := func(filename string) (string, error) {
		// Open file
		f, err := os.Open(filename)
		if err != nil {
			return "", err
		}

		// Prepare hasher
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			return "", err
		}

		// No error
		return fmt.Sprintf("%x", h.Sum(nil)), nil
	}
	return func() error {
		// Compile template
		formulaTpl, err := template.New("Formula").Funcs(map[string]interface{}{
			"sha256file": sha256sum,
		}).Parse(formulaTemplate)
		if err != nil {
			return err
		}

		// Merge data
		var buf bytes.Buffer
		if errTmpl := formulaTpl.Execute(&buf, &formulaModel{
			Repository:  cmd.Package,
			Bin:         cmd.Kebab(),
			Formula:     cmd.Camel(),
			Description: cmd.Description,
			Release:     os.Getenv("RELEASE"),
		}); errTmpl != nil {
			return errTmpl
		}

		// Write output to Stdout
		_, errWrite := buf.WriteTo(os.Stdout)
		return errWrite
	}
}
