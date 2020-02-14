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
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/util"
)

type evOrgIdEtsi struct{}

func (l *evOrgIdEtsi) Initialize() error {
	return nil
}

func (l *evOrgIdEtsi) CheckApplies(c *x509.Certificate) bool {
	_, isPsd2 := util.IsQcStatemPresent(c, &util.IdEtsiPsd2Statem)
	orgId := util.GetSubjectOrgId(c.RawSubject)
	if isPsd2 || !util.IsEV(c.PolicyIdentifiers) || !orgId.IsPresent {
		// PSD2 certificates follow different requirements
		return false
	}
	return true
}

func (l *evOrgIdEtsi) Execute(c *x509.Certificate) *lint.LintResult {
	orgId := util.GetSubjectOrgId(c.RawSubject)

	errStr, _ := util.ParseCabfOrgId(orgId.Value, true)
	if errStr != "" {
		return &lint.LintResult{Status: lint.Error, Details: errStr}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ev_orgid_etsi",
		Description:   "If the subject:organizationIdentifier field is present in an EV certificate, then this lint checks that the format of its contents is in conformance to the CAB/F EV Guidelines and ETSI certificate profiles",
		Citation:      "CA/Browser Forum EV Guidelines v1.7, Sec. 9.2.8 and ETSI TS 119 412-1 V1.3.1 (2019-08), Sec. 5.1.4",
		Source:        lint.EtsiEsi,
		EffectiveDate: util.EtsiEn119_412_1_V1_3_1_Date,
		Lint:          &evOrgIdEtsi{},
	})
}
