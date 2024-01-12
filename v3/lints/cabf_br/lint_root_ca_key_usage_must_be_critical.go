package cabf_br

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type rootCAKeyUsageMustBeCritical struct{}

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_root_ca_key_usage_must_be_critical",
			Description:   "Root CA certificates MUST have Key Usage Extension marked critical",
			Citation:      "BRs: 7.1.2.1",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.RFC2459Date,
		},
		Lint: NewRootCAKeyUsageMustBeCritical,
	})
}

func NewRootCAKeyUsageMustBeCritical() lint.LintInterface {
	return &rootCAKeyUsageMustBeCritical{}
}

func (l *rootCAKeyUsageMustBeCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsRootCA(c) && util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *rootCAKeyUsageMustBeCritical) Execute(c *x509.Certificate) *lint.LintResult {
	keyUsageExtension := util.GetExtFromCert(c, util.KeyUsageOID)
	if keyUsageExtension.Critical {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Error}
	}
}
