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

// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package file

const (
	// DefaultSSLDirectory defines the location where the SSL certificates will be generated
	// This directory contains all the SSL certificates that are specified in Ingress rules.
	// The name of each file is <namespace>-<secret name>.pem. The content is the concatenated
	// certificate and key.
	DefaultSSLDirectory = "/opt/ibm/router/nginx/ssl"
)

var (
	directories = []string{
		"/opt/ibm/router/nginx/template",
		DefaultSSLDirectory,
	}

	files = []string{
		"/tmp/nginx.pid",
	}
)
