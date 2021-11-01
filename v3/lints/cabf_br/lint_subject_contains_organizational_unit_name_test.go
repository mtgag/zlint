package cabf_br

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubjectContainsOrganizationalUnitName(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Certificate is issued before authoritative date",
			InputFilename:  "subjectDnWithOuBeforeAuthoritativeDate.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Certificate is issued after authoritative date and subject does not contain organizational unit name",
			InputFilename:  "subjectDnWithoutOuEntry.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:            "Certificate is issued after authoritative date and subject contains organizational unit name",
			InputFilename:   "subjectDnWithProhibitedOuEntry.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `subject:organizationalUnitName is prohibited for certificates issued on or after September 1, 2022`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_subject_contains_organizational_unit_name", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
