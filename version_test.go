// Copyright 2021 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package pe

import (
	"testing"
)

var peVersionResourceTests = []struct {
	in               string
	out              error
	versionResources map[string]string
}{
	{
		getAbsoluteFilePath("test/putty.exe"),
		nil,
		map[string]string{"CompanyName": "Simon Tatham", "FileDescription": "SSH, Telnet and Rlogin client", "FileVersion": "Release 0.73 (with embedded help)", "InternalName": "PuTTY", "OriginalFilename": "PuTTY", "ProductName": "PuTTY suite", "ProductVersion": "Release 0.73"},
	},
	{
		getAbsoluteFilePath("test/brave.exe"),
		nil,
		map[string]string{"CompanyName": "Brave Software, Inc.", "FileDescription": "Brave Browser", "FileVersion": "80.1.7.92", "InternalName": "chrome_exe"},
	},
	{
		getAbsoluteFilePath("test/impbyord.exe"),
		nil,
		map[string]string{},
	},
	{
		getAbsoluteFilePath("test/WdBoot.sys"),
		nil,
		map[string]string{"CompanyName": "Microsoft Corporation", "FileDescription": "Microsoft antimalware boot driver", "FileVersion": "4.18.1906.3 (GitEnlistment(winpbld).190621-1227)", "InternalName": "WdBoot"},
	},
	{
		getAbsoluteFilePath("test/shimeng.dll"),
		nil,
		map[string]string{"CompanyName": "Microsoft Corporation", "FileDescription": "Shim Engine DLL", "FileVersion": "10.0.17763.1 (WinBuild.160101.0800)", "OriginalFilename": "Shim Engine DLL (IAT)", "LegalCopyright": "© Microsoft Corporation. All rights reserved.", "InternalName": "Shim Engine DLL (IAT)", "ProductName": "Microsoft® Windows® Operating System", "ProductVersion": "10.0.17763.1"},
	},
}

func TestParseVersionResources(t *testing.T) {
	for _, tt := range peVersionResourceTests {
		t.Run(tt.in, func(t *testing.T) {
			file, err := New(tt.in, &Options{})
			if err != nil {
				t.Fatalf("New(%s) failed, reason: %v", tt.in, err)
			}

			got := file.Parse()
			if got != nil {
				t.Errorf("Parse(%s) got %v, want %v", tt.in, got, tt.out)
			}
			vers, err := file.ParseVersionResources()
			if err != nil {
				t.Fatalf("ParseVersionResurces(%s) failed, reason: %v", tt.in, err)
			}
			for k, v := range tt.versionResources {
				val, ok := vers[k]
				if !ok {
					t.Errorf("%s: should have %s version resource", tt.in, k)
				}
				if val != v {
					t.Errorf("%s: expected: %s version resource got: %s. Available resources: %v", tt.in, v, val, vers)
				}
			}
		})
	}
}
