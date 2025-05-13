package nimbusec

const (
	IssueEventBlacklist  = "blacklist"
	IssueEventDefacement = "defacement"
	IssueEventMalware    = "malware"

	IssueEventCMSVersion    = "cms-version"
	IssueEventCMSVulnerable = "cms-vulnerable"
	IssueEventCMSTampered   = "cms-tampered"

	IssueEventSuspiciousLink     = "suspicious-link"
	IssueEventSuspiciousRequest  = "suspicious-request"
	IssueEventSuspiciousRedirect = "suspicious-redirect"
	IssueEventParkedDomain       = "parked"

	IssueEventTLSCipherSuite        = "tls-ciphersuite"
	IssueEventTLSProtocol           = "tls-protocol"
	IssueEventTLSExpires            = "tls-expires"
	IssueEventTLSHostnameMismatch   = "tls-hostname"
	IssueEventTLSNoTrust            = "tls-notrust"
	IssueEventTLSSigAlg             = "tls-sigalg"
	IssueEventTLSLegacy             = "tls-legacy"
	IssueEventTLSMisconfiguredChain = "tls-misconfigured-chain"
	IssueEventTLSRevokedCertificate = "tls-revoked-cert"

	IssueEventBaselineEmpty    = "baseline-empty"
	IssueEventContentViolation = "content-violation"
	IssueEventNoHTTPSRedirect  = "no-https-redirect"
	IssueEventSeospam          = "seospam"
	IssueEventWebshell         = "webshell"

	IssueEventHijackResource = "hijack-resource"
	IssueEventHijackLink     = "hijack-link"
	IssueEvent404Link        = "404-link"
	IssueEventTakeoverDNS    = "takeover-dns"
	IssueEventSRIMissing     = "sri-missing"
	IssueEventSRIInvalid     = "sri-invalid"

	IssueEventOpenDir  = "config-opendir"
	IssueEventPHPError = "config-phperror"
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
