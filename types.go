package nimbusec

import (
	"encoding/json"
	"fmt"
	"time"
)

type ExternalID struct {
	ExternalID string `json:"externalId"`
	Comment    string `json:"comment"`
}

type DomainID int64
type Domain struct {
	ID          DomainID     `json:"id"`
	Bundle      BundleID     `json:"bundle"`
	Name        string       `json:"name"`
	URL         string       `json:"url"`
	ExternalIDs []ExternalID `json:"externalIds,omitempty"`
}

type DomainFilter struct {
	Name       string `url:"name,omitempty"`
	ExternalID string `url:"externalId,omitempty"`
}

type DomainApplication struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Path       string `json:"path"`
	Category   string `json:"category"`
	Source     string `json:"source"`
	Latest     bool   `json:"latest"`
	Vulnerable bool   `json:"vulnerable"`
}

type DomainMetadata struct {
	Domain    DomainID          `json:"domain"`
	Redirects []Redirect        `json:"redirects"`
	Headers   map[string]string `json:"headers"`
}

type DomainStats struct {
	Domains     int `json:"domains"`
	Malware     int `json:"malware"`
	Webshell    int `json:"webshell"`
	Application int `json:"application"`
	TLS         int `json:"tls"`
	Reputation  int `json:"reputation"`
	Defacement  int `json:"defacement"`
}

type Redirect struct {
	URL        string `json:"url"`
	Address    string `json:"address"`
	StatusCode int    `json:"statusCode"`
}

type BundleID string
type Bundle struct {
	ID           BundleID `json:"id"`
	Name         string   `json:"name"`
	StartDate    int64    `json:"startDate"`
	EndDate      int64    `json:"endDate"`
	Price        int      `json:"price"`
	Currency     string   `json:"currency"`
	Active       int      `json:"active"`
	Capacity     int      `json:"capacity"`
	ChargingType string   `json:"chargingType"`

	Features struct {
		Defacement struct {
			Available bool `json:"available"`
			Nimbusec  bool `json:"nimbusec"`
			ZoneH     bool `json:"zoneh"`
		} `json:"defacement"`
		UnwantedContent struct {
			Available bool `json:"available"`
		} `json:"unwantedContent"`
		Malware struct {
			Available bool `json:"available"`
			Nimbusec  bool `json:"nimbusec"`
			Ikarus    bool `json:"ikarus"`
			Avira     bool `json:"avira"` // Deprecated. Avira is not used as of 01.03.2022 anymore
			LastLine  bool `json:"lastline"`
			ClamAV    bool `json:"clamav"`
		} `json:"malware"`
		Reputation struct {
			Available bool `json:"available"`
		} `json:"reputation"`
		TLS struct {
			Available bool `json:"available"`
		} `json:"tls"`
		Webshell struct {
			Available bool `json:"available"`
		} `json:"webshell"`
		Application struct {
			Available bool `json:"available"`
		} `json:"application"`
		Scanning struct {
			Available        bool `json:"available"`
			FastScanInterval int  `json:"fastScanInterval"`
			DeepScanInterval int  `json:"deepScanInterval"`
			Quota            int  `json:"quota"`
			FromEU           bool `json:"fromEU"`
			FromUS           bool `json:"fromUS"`
			FromASIA         bool `json:"fromAsia"`
			Mobile           bool `json:"mobile"`
		} `json:"scanning"`
		Notification struct {
			Available   bool `json:"available"`
			EMail       bool `json:"email"`
			TextMessage bool `json:"textMessage"`
		} `json:"notification"`
	} `json:"features"`
}

type NotificationID int64
type Notification struct {
	ID         NotificationID `json:"id"`
	Domain     DomainID       `json:"domain"`
	User       UserID         `json:"user"`
	Transport  string         `json:"transport"`
	Blacklist  int            `json:"blacklist"`
	Defacement int            `json:"defacement"`
	Malware    int            `json:"malware"`
	Webhook    *string        `json:"webhook,omitempty"`
}

type NotificationUpdate struct {
	Transport  string `json:"transport"`
	Blacklist  int    `json:"blacklist"`
	Defacement int    `json:"defacement"`
	Malware    int    `json:"malware"`
}

type UserID string
type User struct {
	ID           UserID `json:"id"`
	Login        string `json:"login"`
	Mail         string `json:"mail"`
	Role         string `json:"role"`
	Company      string `json:"company"`
	Surname      string `json:"surname"`
	Forename     string `json:"forename"`
	Title        string `json:"title"`
	Mobile       string `json:"mobile"`
	Password     string `json:"password"`
	SignatureKey string `json:"signatureKey"`
}

type UserDomains struct {
	Domains []DomainID `json:"domains"`
}

type UserConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}
type DomainConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ScreenshotType string

type Screenshot struct {
	Target   string `json:"target"`
	Previous Image  `json:"previous"`
	Current  Image  `json:"current"`
}

type Image struct {
	Date time.Time `json:"date"`
	MIME string    `json:"mime"`
	URL  string    `json:"url"`
}

type Region string
type Viewport string
type ScreenshotFilter struct {
	Region   Region   `url:"region,omitempty"`
	Viewport Viewport `url:"viewport,omitempty"`
}

type IssueID int64
type Issue struct {
	ID          IssueID      `json:"id"`
	Domain      DomainID     `json:"domain"`
	Status      IssueStatus  `json:"status"`
	Event       string       `json:"event"`
	Category    string       `json:"category"`
	Severity    int          `json:"severity"`
	FirstSeen   time.Time    `json:"firstSeen"`
	LastSeen    time.Time    `json:"lastSeen"`
	Regions     []string     `json:"regions,omitempty"`
	Viewports   []string     `json:"viewports,omitempty"`
	Details     interface{}  `json:"details,omitempty"`
	ExternalIDs []ExternalID `json:"externalIds,omitempty"`
}

type IssueUpdate struct {
	Status      IssueStatus  `json:"status"`
	Comment     string       `json:"comment"`
	ExternalIDs []ExternalID `json:"externalIds"`
}

type IssueFilter struct {
	Status     IssueStatus `url:"status,omitempty"`
	Severity   int         `url:"severity,omitempty"`
	Event      string      `url:"event,omitempty"`
	Category   string      `url:"category,omitempty"`
	Limit      int         `url:"limit,omitempty"`
	Sort       string      `url:"sort,omitempty"`
	ExternalID string      `url:"externalId,omitempty"`
}

type IssueStatus string

const (
	IssueStatusPending       IssueStatus = "pending"
	IssueStatusAcknowledged  IssueStatus = "acknowledged"
	IssueStatusIgnored       IssueStatus = "ignored"
	IssueStatusFalsePositive IssueStatus = "falsepositive"
)

type IssuesSummary struct {
	Domain DomainID `json:"domain"`

	MalwareSeverity     int `json:"malwareSeverity"`
	WebshellSeverity    int `json:"webshellSeverity"`
	ApplicationSeverity int `json:"applicationSeverity"`
	TLSSeverity         int `json:"tlsSeverity"`
	ReputationSeverity  int `json:"reputationSeverity"`
	DefacementSeverity  int `json:"defacementSeverity"`

	MalwareCount     int `json:"malwareCount"`
	WebShellCount    int `json:"webShellCount"`
	ApplicationCount int `json:"applicationCount"`
	TLSCount         int `json:"tlsCount"`
	ReputationCount  int `json:"reputationCount"`
	DefacementCount  int `json:"defacementCount"`
}

type IssueHistory map[string]IssueCount
type IssueCount struct {
	Malware     int `json:"malware"`
	Webshell    int `json:"webshell"`
	Application int `json:"application"`
	TLS         int `json:"tls"`
	Reputation  int `json:"reputation"`
	Defacement  int `json:"defacement"`
}

type ApplicationOutdatedDetails struct {
	Name          string `json:"name"`
	URL           string `json:"url,omitempty"`
	Path          string `json:"path,omitempty"`
	Version       string `json:"version"`
	LatestVersion string `json:"latestVersion"`
}

type ApplicationVulnerableDetails struct {
	Name            string `json:"name"`
	URL             string `json:"url,omitempty"`
	Path            string `json:"path,omitempty"`
	Version         string `json:"version"`
	Vulnerabilities []struct {
		CVE         string  `json:"cve"`
		Score       float64 `json:"score"`
		Description string  `json:"description"`
		Link        string  `json:"link"`
	} `json:"vulnerabilities"`
}

type DefacementDetails struct {
	URL     string `json:"url"`
	Threat  string `json:"threat"`
	Profile string `json:"profile"`
}

type ZoneHDetails struct {
	URL    string `json:"url"`
	Threat string `json:"threat"`
}

// Initiator holds all information about the initiator of an request/malware issue
type Initiator struct {
	InitType string `json:"type"`
	URL      string `json:"url"`
	Line     string `json:"line"`
	Column   string `json:"column"`
}

type MalwareDetails struct {
	URL        string      `json:"url"`
	Threat     string      `json:"threat"`
	AV         string      `json:"av"`
	Initiators []Initiator `json:"initiators"`
	Profile    string      `json:"profile"`
}

type LastlineDetails struct {
	Score           int `json:"score"`
	AnalysisSubject struct {
		URL     string `json:"url"`
		Referer string `json:"referer"`
	} `json:"analysis_subject"`
	Threat            string   `json:"threat"`
	ThreatClass       string   `json:"threat_class"`
	MaliciousActivity []string `json:"malicious_activity"`
}

type TLSCertificateDetails struct {
	Order      int      `json:"order"`
	NotAfter   int64    `json:"notAfter"`
	NotBefore  int64    `json:"notBefore"`
	Issuer     string   `json:"issuer"`
	CName      string   `json:"cName"`
	AltNames   []string `json:"altNames"`
	SigAlgName string   `json:"sigAlgName"`
	SigAlgOID  string   `json:"sigAlgOID"`
}

type TLSConfigurationDetails struct {
	Protocol string `json:"protocol,omitempty"`
	Cipher   string `json:"cipher,omitempty"`
}

type BlacklistDetails struct {
	Blacklist    string   `json:"blacklist"`
	BlacklistURL string   `json:"blacklistURL"`
	Reasons      []string `json:"reasons"`
	Profile      string   `json:"profile"`
}

type SuspiciousLinkDetails struct {
	URL   string `json:"url"`
	Links []struct {
		Link       string             `json:"link"`
		Blacklists []BlacklistDetails `json:"blacklists"`
	} `json:"links"`
}

type SuspiciousRequestDetails struct {
	Entity     string             `json:"entity"`
	URLs       []string           `json:"urls"`
	Blacklists []BlacklistDetails `json:"blacklists"`
}

type BaselineDetails struct {
	ClientID string `json:"clientID"`
}

type WebshellDetails struct {
	AV          string `json:"av"`
	MD5         string `json:"md5"`
	Path        string `json:"path"`
	Size        int    `json:"size"`
	Owner       string `json:"owner"`
	Group       string `json:"group"`
	MTime       int    `json:"mtime"`
	Threat      string `json:"threat"`
	Permissions string `json:"permissions"`
}

type ContentViolationDetails struct {
	NewHostnames       map[string][]string `json:"newHostnames"`
	NewAResources      map[string]struct{} `json:"newAResources"`
	NewImgResources    map[string]struct{} `json:"newImgResources"`
	NewInputResources  map[string]struct{} `json:"newInputResources"`
	NewButtonResources map[string]struct{} `json:"newButtonResources"`
	NewEventAttributes map[string]struct{} `json:"newEventAttributes"`
	Profile            string              `json:"profile"`
}

type SeospamDetails struct {
	Profile         string        `json:"profile"`
	CrawlerResult   SeospamResult `json:"crawlerResult"`
	GooglebotResult SeospamResult `json:"googlebotResult"`
}

type SeospamResult struct {
	Title       string `json:"Title"`
	Keywords    string `json:"Keywords"`
	Description string `json:"Description"`
}

type CMSTamperedDetails struct {
	Url          string `json:"url"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Version      string `json:"version"`
	OriginalHash string `json:"originalHash"`
	TamperedHash string `json:"tamperedHash"`
}

type SuspiciousRedirectDetails struct {
	Profile              string     `json:"profile"`
	CurrentRedirects     []Redirect `json:"currentRedirects"`
	ValidRedirectDomains []string   `json:"validRedirectDomains"`
}

type HijackDetails struct {
	URL        string   `json:"url"`
	Initiators []string `json:"initiators"`
}

type TakeoverDNSDetails struct {
	Service     string   `json:"service"`
	Nameservers []string `json:"nameservers"`
}

// UnmarshalJSON unmarshals Issues and attaches the correct Details type instead of the interface{}
func (issue *Issue) UnmarshalJSON(b []byte) error {
	type TempIssueType Issue
	tempIssue := TempIssueType{}
	err := json.Unmarshal(b, &tempIssue) // seperate type needed to prevent recursive unmarshal-loop
	if err != nil {
		return err
	}
	*issue = Issue(tempIssue)
	if issue.Details == nil {
		return nil // no details to cast
	}

	// issue.Details is map[string]interface{} at this point. marshal again, to allow a specific unmarshal
	detail, err := json.Marshal(issue.Details)
	if err != nil {
		return err
	}

	issue.Details, err = UnmarshalDetails(issue.Event, detail)
	if err != nil {
		return err
	}

	return nil
}

// UnmarshalDetails converts the given details value to the correct details type
func UnmarshalDetails(event string, details []byte) (interface{}, error) {
	var err error
	switch event {
	case IssueEventBlacklist:
		specificDetails := BlacklistDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventSuspiciousLink:
		specificDetails := SuspiciousLinkDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventSuspiciousRequest:
		specificDetails := SuspiciousRequestDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventMalware:
		specificDetails := MalwareDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventDefacement:
		specificDetails := DefacementDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventCMSVersion:
		specificDetails := ApplicationOutdatedDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventCMSVulnerable:
		specificDetails := ApplicationVulnerableDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventWebshell:
		specificDetails := WebshellDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventBaselineEmpty:
		specificDetails := BaselineDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventTLSProtocol, IssueEventTLSCipherSuite:
		specificDetails := TLSConfigurationDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventTLSSigAlg, IssueEventTLSNoTrust, IssueEventTLSHostnameMismatch, IssueEventTLSExpires, IssueEventTLSLegacy, IssueEventTLSMisconfiguredChain, IssueEventTLSRevokedCertificate:
		specificDetails := TLSCertificateDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventContentViolation:
		specificDetails := ContentViolationDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventSeospam:
		specificDetails := SeospamDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventCMSTampered:
		specificDetails := CMSTamperedDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventSuspiciousRedirect:
		specificDetails := SuspiciousRedirectDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case EventHijackResource, EventHijackLink, Event404Link:
		specificDetails := HijackDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case EventTakeoverDNS:
		specificDetails := TakeoverDNSDetails{}
		err = json.Unmarshal(details, &specificDetails)
		if err != nil {
			return nil, err
		}
		return specificDetails, nil
	case IssueEventNoHTTPSRedirect:
		// currently, no specific details are held by "no-https-redirect" issues
		return nil, nil
	default:
		return nil, fmt.Errorf("event '%s' unknown | %s", event, string(details))
	}
}

type Role string
type TokenID string
type APIToken struct {
	Name        string    `json:"name"`
	Key         TokenID   `json:"key"`
	Secret      string    `json:"secret"`
	Role        Role      `json:"role"`
	LastCall    time.Time `json:"lastCall"`
	LastVersion int       `json:"lastVersion"`
}

type Agent struct {
	Os      string `json:"os"`
	Arch    string `json:"arch"`
	Version string `json:"version"`
	Md5     string `json:"md5"`
	Sha1    string `json:"sha1"`
	Format  string `json:"format"`
	URL     string `json:"url"`
}
