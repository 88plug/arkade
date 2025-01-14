// Copyright (c) arkade author(s) 2022. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package helm

import "testing"

func Test_GetHelmURL_GitBash(t *testing.T) {
	arch := "amd64"
	os := "mingw64_nt-10.0-18362"
	helmVersion := "v3.9.3"

	got := GetHelmURL(arch, os, helmVersion)
	want := "https://get.helm.sh/helm-v3.9.3-windows-amd64.tar.gz"
	if got != want {
		t.Fatalf("want: %s, but got: %s", want, got)
	}
}
