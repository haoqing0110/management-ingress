/*
Copyright 2017 The Kubernetes Authors.

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

package file

import (
	"testing"
)

func TestNewFakeFS(t *testing.T) {
	fs, err := NewFakeFS()
	if err != nil {
		t.Fatalf("unexpected error creating filesystem abstraction: %v", err)
	}

	if fs == nil {
		t.Fatal("expected a filesystem but none returned")
	}

	_, err = fs.Stat("/opt/ibm/router/nginx/conf/nginx.conf")
	if err != nil {
		t.Fatalf("unexpected error reading default nginx.conf file: %v", err)
	}
}
