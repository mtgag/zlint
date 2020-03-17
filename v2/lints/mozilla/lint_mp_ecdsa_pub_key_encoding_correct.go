package mozilla

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

/************************************************
https://www.mozilla.org/en-US/about/governance/policies/security-group/certs/policy/

When ECDSA keys are encoded in a SubjectPublicKeyInfo structure, the algorithm field MUST be one of the following, as
specified by RFC 5480, Section 2.1.1:

The encoded AlgorithmIdentifier for a P-256 key MUST match the following hex-encoded
bytes: > 301306072a8648ce3d020106082a8648ce3d030107.

The encoded AlgorithmIdentifier for a P-384 key MUST match the following hex-encoded
bytes: > 301006072a8648ce3d020106052b81040022.

The above encodings consist of an ecPublicKey OID (1.2.840.10045.2.1) with a named curve parameter of the corresponding
curve OID. Certificates MUST NOT use the implicit or specified curve forms.

This lint covers the previous part. Next part is covered by e_mp_ecdsa_signature_encoding_correct.

When a root or intermediate certificate's ECDSA key is used to produce a signature, only the following algorithms may
be used, and with the following encoding requirements:

If the signing key is P-256, the signature MUST use ECDSA with SHA-256. The encoded AlgorithmIdentifier MUST match the
following hex-encoded bytes: 300a06082a8648ce3d040302.

If the signing key is P-384, the signature MUST use ECDSA with SHA-384. The encoded AlgorithmIdentifier MUST match the
following hex-encoded bytes: 300a06082a8648ce3d040303.

The above encodings consist of the corresponding OID with the parameters field omitted, as specified by RFC 5758,
Section 3.2. Certificates MUST NOT include a NULL parameter. Note this differs from RSASSA-PKCS1-v1_5, which includes
an explicit NULL.

************************************************/

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2/lint"
	"github.com/zmap/zlint/v2/util"
)

type ecdsaPubKeyAidEncoding struct{}

var ecSPKIAlgorithmIDToDER = [2][]byte{
	// encoded AlgorithmIdentifier for a P-256 key
	{0x30, 0x13, 0x06, 0x07, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x02, 0x01, 0x06, 0x08, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x03, 0x01, 0x07},
	// encoded AlgorithmIdentifier for a P-384 key
	{0x30, 0x10, 0x06, 0x07, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x02, 0x01, 0x06, 0x05, 0x2b, 0x81, 0x04, 0x00, 0x22},
}

func (l *ecdsaPubKeyAidEncoding) Initialize() error {
	return nil
}

func (l *ecdsaPubKeyAidEncoding) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.ECDSA
}

func (l *ecdsaPubKeyAidEncoding) Execute(c *x509.Certificate) *lint.LintResult {

	encodedPublicKey, err := util.GetPublicKeyAidEncoded(c)

	if err != nil {
		return &lint.LintResult{Status: lint.Error, Details: "error reading public key algorithm identifier from TBS"}
	}

	for _, encoding := range ecSPKIAlgorithmIDToDER {
		if bytes.Equal(encodedPublicKey, encoding) {
			return &lint.LintResult{Status: lint.Pass}
		}
	}

	return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Wrong encoding of ECC public key. Got the unsupported %s", hex.EncodeToString(encodedPublicKey))}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_mp_ecdsa_pub_key_encoding_correct",
		Description:   "The encoded algorithm identifiers for ECDSA public keys MUST match specific bytes",
		Citation:      "Mozilla Root Store Policy / Section 5.1.2",
		Source:        lint.MozillaRootStorePolicy,
		EffectiveDate: util.MozillaPolicy27Date,
		Lint:          &ecdsaPubKeyAidEncoding{},
	})
}
