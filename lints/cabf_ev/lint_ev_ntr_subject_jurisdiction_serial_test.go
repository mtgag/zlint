package cabf_ev

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

func TestEvAltRegNumSubjectJurisSerial(t *testing.T) {
	m := map[string]lint.LintStatus{
		"EvAltRegNumCert52NoOrgId.pem":              lint.NA,
		"EvAltRegNumCert56JurContryNotMatching.pem": lint.Error,
		"EvAltRegNumCert57NtrJurSopMissing.pem":     lint.Error,
		"EvAltRegNumCert61Valid.pem":                lint.NA,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_ev_ntr_subject_jurisdiction_serial", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
