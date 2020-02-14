package etsi

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
	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/test"
	"testing"
)

func TestEvAltRegNumOrgidEtsi(t *testing.T) {
	m := map[string]lint.LintStatus{
		"evAllGood.pem":                             lint.NA,
		"EvAltRegNumCert54OrgIdInvalid.pem":         lint.Error,
		"oiLEI.pem":                                 lint.Pass,
		"EvAltRegNumCert56JurContryNotMatching.pem": lint.Pass,
		"EvAltRegNumCert57NtrJurSopMissing.pem":     lint.Pass,
		"EvAltRegNumCert58ValidNtr.pem":             lint.Pass,
		"EvAltRegNumCert59Valid.pem":                lint.Pass,
		"EvAltRegNumCert60OrgIdInvalid.pem":         lint.Error,
		"EvAltRegNumCert61Valid.pem":                lint.Pass,
		"EvAltRegNumCert62OrgIdLenZero.pem":         lint.Error,
		"EvAltRegNumCert63Valid.pem":                lint.Pass,
	}
	for inputPath, expected := range m {
		out := test.TestLint("e_ev_orgid_etsi", inputPath)

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
