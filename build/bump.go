// Copyright Splunk Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var _ = goyek.Define(goyek.Task{
	Name:  "bump",
	Usage: "go get -u -t ./...",
	Action: func(tf *goyek.TF) {
		ForGoModules(tf, func(tf *goyek.TF) {
			cmd.Exec(tf, "go get -u -t ./...")
		}, dirBuild) // '/build' should be bumped manually as golangci-lint's deps often have breaking changes
	},
})
