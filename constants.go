package nimbusec

const (
	IssueEventBlacklist             = "blacklist"
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
	IssueEventBaselineEmpty         = "baseline-empty"
	IssueEventContentViolation      = "content-violation"
	IssueEventNoHTTPSRedirect       = "no-https-redirect"
)

const (
	RoleAgent    Role = "agent"
	RoleReadOnly Role = "readonly"
)

const (
	ScreenshotCurrent  ScreenshotType = "current"
	ScreenshotPrevious ScreenshotType = "previous"

	RegionEU   Region = "EU"
	RegionUS   Region = "US"
	RegionASIA Region = "ASIA"

	ViewportDesktop Viewport = "desktop"
	ViewportMobile  Viewport = "mobile"
)
