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
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type qcStatemQctypeWeb struct{}

func (this *qcStatemQctypeWeb) getStatementOid() *asn1.ObjectIdentifier {
	return &util.IdEtsiQcsQcType
}

func (l *qcStatemQctypeWeb) Initialize() error {
	return nil
}

func (l *qcStatemQctypeWeb) CheckApplies(c *x509.Certificate) bool {
	if !util.IsExtInCert(c, util.QcStateOid) {
		return false
	}
	return util.IsQCStatementPresent(c, util.IdEtsiQcsQcType.String())
}

func (l *qcStatemQctypeWeb) Execute(c *x509.Certificate) *lint.LintResult {
	if len(c.QCStatements.ParsedStatements.Types) != 1 {
		return &lint.LintResult{Status: lint.Error, Details: "invalid number of QcType objects"}
	}
	qcType := c.QCStatements.ParsedStatements.Types[0]

	if len(qcType.TypeIdentifiers) == 0 {
		return &lint.LintResult{Status: lint.Error, Details: "no QcType present, sequence of OIDs is empty"}
	}

	for _, t := range qcType.TypeIdentifiers {
		if t.Equal(util.IdEtsiQcsQctWeb) {
			return &lint.LintResult{Status: lint.Pass}
		}
	}

	return &lint.LintResult{Status: lint.Warn, Details: fmt.Sprintf("etsi Type does not indicate certificate as a 'web' certificate")}

}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_qcstatem_qctype_web",
		Description:   "Checks that a QC Statement of the type Id-etsi-qcs-QcType features features at least the type IdEtsiQcsQctWeb",
		Citation:      "ETSI EN 319 412 - 5 V2.2.1 (2017 - 11) / Section 4.2.3",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn319_412_5_V2_2_1_Date,
		Lint:          &qcStatemQctypeWeb{},
	})
}
