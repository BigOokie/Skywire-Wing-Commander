// Copyright © 2018 BigOokie
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
package utils

import (
	"testing"
)

func Test_FileExists_NotOk(t *testing.T) {
	badFilename := "this-file-does-not-exist.test"
	if FileExists(badFilename) {
		t.Errorf("File (%s) does not exist but was detected as existing.", badFilename)
	}
}

func Test_FileExists_Ok(t *testing.T) {
	goodFilename := "./utils_test.go"
	if !FileExists(goodFilename) {
		t.Errorf("File (%s) does not exist but should.", goodFilename)
	}
}

func Test_UserHome(t *testing.T) {
	if UserHome() == "" {
		t.Error("Failed to determine User Home Path")
	}
}

func Test_AppVerInfo(t *testing.T) {
	//expectappverinfostr := "v0.0.0 [abcdefg] 2019-01-01"
	expectappvernumstr := "v0.0.0"
	InitAppVersionInfo("0.0.0", "abcdefg", "master", "dirty", "test summary", "2019-01-01")

	//if AppVersionInfoString() != expectappverinfostr {
	if AppVersionInfoString() != expectappvernumstr {
		t.Error("App Version Info String is incorrect", AppVersionInfoString())
	}

	if AppVersionNumberString() != expectappvernumstr {
		t.Error("App Version Numbner String is incorrect", AppVersionNumberString())
	}
}
