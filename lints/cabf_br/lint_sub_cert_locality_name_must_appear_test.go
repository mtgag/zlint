package cabf_br

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"testing"

	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/test"
)

func TestSubCertLocalityNameMustAppear(t *testing.T) {
	inputPath := "subCertLocalityNameMustAppear.pem"
	expected := lint.Error
	out := test.TestLint("e_sub_cert_locality_name_must_appear", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertLocalityNameDoesNotNeedToAppear(t *testing.T) {
	inputPath := "subCertLocalityNameDoesNotNeedToAppear.pem"
	expected := lint.Pass
	out := test.TestLint("e_sub_cert_locality_name_must_appear", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}