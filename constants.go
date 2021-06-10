package nimbusec

const (
	IssueEventBlacklist             = "blacklist"
	IssueEventChangedFile           = "changed-file"
	IssueEventCMSVersion            = "cms-version"
	IssueEventCMSVulnerable         = "cms-vulnerable"
	IssueEventDefacement            = "defacement"
	IssueEventMalware               = "malware"
	IssueEventSuspiciousLink        = "suspicious-link"
	IssueEventSuspiciousRequest     = "suspicious-request"
	IssueEventTLSCipherSuite        = "tls-ciphersuite"
	IssueEventTLSProtocol           = "tls-protocol"
	IssueEventTLSExpires            = "tls-expires"
	IssueEventTLSHostnameMismatch   = "tls-hostname"
	IssueEventTLSNoTrust            = "tls-notrust"
	IssueEventTLSSigAlg             = "tls-sigalg"
	IssueEventTLSLegacy             = "tls-legacy"
	IssueEventTLSMisconfiguredChain = "tls-misconfigured-chain"
	IssueEventTLSRevokedCertificate = "tls-revoked-cert"
	IssueEventWebshell              = "webshell"
)

const (
	RoleAgent    Role = "agent"
	RoleReadOnly Role = "readonly"
)
