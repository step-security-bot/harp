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

package selector

import (
	"testing"

	bundlev1 "github.com/elastic/harp/api/gen/go/harp/bundle/v1"
)

func Test_matchCel_IsSatisfiedBy(t *testing.T) {
	type fields struct {
		expressions []string
	}
	type args struct {
		object interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    bool
	}{
		{
			name:    "nil",
			wantErr: true,
		},
		{
			name:    "empty",
			args:    args{},
			wantErr: true,
		},
		{
			name: "not supported type",
			fields: fields{
				expressions: []string{"p.is_cso_compliant()"},
			},
			args: args{
				object: struct{}{},
			},
			wantErr: false,
			want:    false,
		},
		{
			name: "supported type: empty object",
			fields: fields{
				expressions: []string{"p.is_cso_compliant()"},
			},
			args: args{
				object: &bundlev1.Package{},
			},
			want: false,
		},
		{
			name: "supported type: invalid return",
			fields: fields{
				expressions: []string{"8"},
			},
			args: args{
				object: &bundlev1.Package{},
			},
			wantErr: true,
			want:    false,
		},
		{
			name: "supported type: match",
			fields: fields{
				expressions: []string{"p.is_cso_compliant()"},
			},
			args: args{
				object: &bundlev1.Package{
					Name: "infra/aws/billing/eu-central-1/rds/postgres/billing/admin_account",
					Labels: map[string]string{
						"patched": "true",
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := MatchCEL(tt.fields.expressions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Error got %v, expected %v", err, tt.wantErr)
				return
			}
			if s == nil {
				return
			}
			if got := s.IsSatisfiedBy(tt.args.object); got != tt.want {
				t.Errorf("celMatcher.IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}