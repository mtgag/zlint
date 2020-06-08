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

package etsi

import (
	"testing"

	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/test"
)

func TestQcStatemQcLimitValue(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "QcStmtValidLimitValue.pem",
			InputFilename:  "QcStmtValidLimitValue.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtInvalidLimitValue.pem",
			InputFilename:  "QcStmtInvalidLimitValue.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiInvalidAlphabeticCurrencyLengthCert29.pem",
			InputFilename:  "QcStmtEtsiInvalidAlphabeticCurrencyLengthCert29.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiInvalidAlphabeticCurrencyCert30.pem",
			InputFilename:  "QcStmtEtsiInvalidAlphabeticCurrencyCert30.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "QcStmtEtsiValidNumericCurrencyCert31.pem",
			InputFilename:  "QcStmtEtsiValidNumericCurrencyCert31.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "QcStmtEtsiInvalidNumericCurrencyCert32.pem",
			InputFilename:  "QcStmtEtsiInvalidNumericCurrencyCert32.pem",
			ExpectedResult: lint.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_qcstatem_qclimitvalue_valid", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
