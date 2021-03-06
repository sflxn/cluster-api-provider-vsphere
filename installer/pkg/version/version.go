// Copyright 2018 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// These fields are set by the compiler using the linker flags upon build via Makefile.
var (
	Version     string
	GitCommit   string
	BuildDate   string
	BuildNumber string
	State       string

	v bool
)

type Build struct {
	Version       string
	GitCommit     string
	BuildDate     string
	BuildNumber   string
	State         string
	PluginVersion int
}

// Show returns whether -version flag is set
func Show() bool {
	return v
}

// String returns a string representation of the version
func String() string {
	return GetBuild().String()
}

// UserAgent returns component/version in HTTP User-Agent header value format
func UserAgent(component string) string {
	v := Version
	if strings.HasPrefix(v, "v") {
		v = v[1:]
	}
	return fmt.Sprintf("%s/%s", component, v)
}

func GetBuild() *Build {
	if BuildNumber == "" {
		BuildNumber = "0"
	}
	return &Build{
		Version:     Version,
		GitCommit:   GitCommit,
		BuildDate:   BuildDate,
		BuildNumber: BuildNumber,
		State:       State,
	}
}

func (v *Build) String() string {
	if v.State == "" {
		v.State = "clean"
	}

	if v.BuildNumber == "" {
		v.BuildNumber = "N/A"
	}
	return fmt.Sprintf("%s git:%s-%s build:%s id:%s runtime:%s", v.Version, v.GitCommit, v.State, v.BuildDate, v.BuildNumber, runtime.Version())
}

func (v *Build) ShortVersion() string {
	return fmt.Sprintf("%s-%s-%s", v.Version, v.BuildNumber, v.GitCommit)
}

// Equal determines if v is equal to b based on BuildNumber
func (v *Build) Equal(b *Build) bool {
	return v.BuildNumber == b.BuildNumber
}

// IsOlder determines if v is older than b based on BuildNumber
func (v *Build) IsOlder(b *Build) (bool, error) {
	if v.Equal(b) {
		return false, nil
	}

	if v.BuildNumber == "" || b.BuildNumber == "" {
		return false, fmt.Errorf("invalid BuildNumber - comparing %q to %q", v.BuildNumber, b.BuildNumber)
	}

	vi, errv := strconv.Atoi(v.BuildNumber)
	bi, errb := strconv.Atoi(b.BuildNumber)
	if errv != nil {
		return false, fmt.Errorf("invalid BuildNumber format %s: %s", v, errv)
	}
	if errb != nil {
		return false, fmt.Errorf("invalid BuildNumber format %s: %s", b, errb)
	}

	buildBefore := vi < bi
	return buildBefore, nil
}

// IsNewer determines if v is newer than b based on BuildNumber
func (v *Build) IsNewer(b *Build) (bool, error) {
	if v.Equal(b) {
		return false, nil
	}
	older, err := v.IsOlder(b)
	if err != nil {
		return false, err
	}
	return !older, nil
}
