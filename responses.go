package mepost

import "time"

// ApiResponse represents a generic API response.
type ApiResponse[T any] struct {
	Success bool            `json:"success"`
	Data    T               `json:"data,omitempty"`
	Errors  []ErrorResponse `json:"errors,omitempty"`
}

// BaseResult represents a generic structure for a paginated list response.
type BaseResult[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
}

// ErrorResponse represents an error response structure.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

// AddDomainResponse represents the response for adding a domain.
type AddDomainResponse struct {
	DKIM   DNSRecord `json:"dkim"`
	DMARC  DNSRecord `json:"dmarc"`
	Domain string    `json:"domain"`
	SPF    DNSRecord `json:"spf"`
}

// CancelWarmUpResponse represents the response for cancelling an IP warm-up.
type CancelWarmUpResponse struct {
	CancelledAt string `json:"cancelledAt"`
	IPAddress   string `json:"ipAddress"`
	StartedAt   string `json:"startedAt"`
}

// GetMessageInfoResponse represents the response for retrieving message information.
type GetMessageInfoResponse struct {
	Email             string             `json:"email"`
	EmailClicksCount  int                `json:"emailClicksCount"`
	EmailClicksDetail []EmailClickDetail `json:"emailClicksDetail"`
	EmailReadsCount   int                `json:"emailReadsCount"`
	EmailReadsDetail  []EmailReadDetail  `json:"emailReadsDetail"`
	State             string             `json:"state"`
	Subject           string             `json:"subject"`
	TemplateID        string             `json:"templateId"`
}

// GetScheduleInfoResponse represents the response for retrieving schedule information.
type GetScheduleInfoResponse struct {
	Details          ScheduleDetails `json:"details"`
	EmailReadsCount  int             `json:"emailReadsCount"`
	EmailReadsUnique int             `json:"emailReadsUnique"`
	HardBounceCount  int             `json:"hardBounceCount"`
	LinkClicksCount  int             `json:"linkClicksCount"`
	OtherBounceCount int             `json:"otherBounceCount"`
	SenderFromEmail  string          `json:"senderFromEmail"`
	SenderFromName   string          `json:"senderFromName"`
	SoftBounceCount  int             `json:"softBounceCount"`
	State            string          `json:"state"`
	Subject          string          `json:"subject"`
	TemplateID       string          `json:"templateId"`
	UnsubscribeCount int             `json:"unsubscribeCount"`
}

// RemoveDomainResponse represents the response for removing a domain.
type RemoveDomainResponse struct {
	Domain    string `json:"domain"`
	RemovedAt string `json:"removedAt"`
}

// SetIpGroupResponse represents the response for setting an IP group.
type SetIpGroupResponse struct {
	IpAddress string  `json:"ipAddress"`
	IpGroup   IPGroup `json:"ipGroup"`
}

// StartWarmUpResponse represents the response for starting an IP warm-up.
type StartWarmUpResponse struct {
	EndAt     string `json:"endAt"`
	IPAddress string `json:"ipAddress"`
	StartAt   string `json:"startAt"`
	Status    string `json:"status"`
}

// DNSRecord represents a DNS record for domain verification.
type DNSRecord struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

// EmailClickDetail represents details of an email click.
type EmailClickDetail struct {
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	IP          string `json:"ip"`
	URL         string `json:"url"`
}

// EmailReadDetail represents details of an email read.
type EmailReadDetail struct {
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	IP          string `json:"ip"`
}

// ScheduleDetails represents the details of a schedule.
type ScheduleDetails struct {
	Clicks       []EmailTransactionEvent `json:"clicks"`
	HardBounces  []EmailTransactionEvent `json:"hardBounces"`
	Reads        []EmailTransactionEvent `json:"reads"`
	SoftBounces  []EmailTransactionEvent `json:"softBounces"`
	Unsubscribes []EmailTransactionEvent `json:"unsubscribes"`
}

// EmailTransactionEvent represents an event related to an email transaction.
type EmailTransactionEvent struct {
	BounceCode    string `json:"bounceCode,omitempty"`
	City          string `json:"city,omitempty"`
	CountryCode   string `json:"countryCode,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	Data          string `json:"data,omitempty"`
	EventType     string `json:"eventType,omitempty"`
	ID            string `json:"id,omitempty"`
	IP            string `json:"ip,omitempty"`
	StatID        string `json:"statId,omitempty"`
	SubscriberID  string `json:"subscriberId,omitempty"`
	TransactionID string `json:"transactionId,omitempty"`
}

// Schedule represents a scheduled email or marketing campaign.
type Schedule struct {
	Approved         bool      `json:"approved"`
	AuthorizedToSend bool      `json:"authorizedToSend"`
	CreatedAt        time.Time `json:"createdAt"`
	CreditAmount     float64   `json:"creditAmount"`
	EmailGroupId     int       `json:"emailGroupId"`
	JobStatus        string    `json:"jobStatus"`
	JobType          string    `json:"jobType"`
	Reason           string    `json:"reason"`
	ResultType       string    `json:"resultType"`
	ScheduledAt      time.Time `json:"scheduledAt"`
	StatId           string    `json:"statId"`
	Template         Template  `json:"template"`
	UpdatedAt        time.Time `json:"updatedAt"`
	UUID             string    `json:"uuid"`
}

// Template represents the structure of an email template.
type Template struct {
	Config    string    `json:"config"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	RawHtml   string    `json:"rawHtml"`
	RawText   string    `json:"rawText"`
	Subject   string    `json:"subject"`
	UpdatedAt time.Time `json:"updatedAt"`
	UUID      string    `json:"uuid"`
}

// EmailGroup represents a group of emails.
type EmailGroup struct {
	CompanyId             int       `json:"companyId"`
	CreatedAt             time.Time `json:"createdAt"`
	GeneralScore          int       `json:"generalScore"`
	IsWeb                 bool      `json:"isWeb"`
	Name                  string    `json:"name"`
	NewsletterScore       int       `json:"newsletterScore"`
	Priority              int       `json:"priority"`
	TotalActiveSubscriber int       `json:"totalActiveSubscriber"`
	TotalSubscriber       int       `json:"totalSubscriber"`
	TotalUnsubscribe      int       `json:"totalUnsubscribe"`
	UpdatedAt             time.Time `json:"updatedAt"`
	UUID                  string    `json:"uuid"`
}

// EmailGroupWithCounts represents an email group with additional statistics.
type EmailGroupWithCounts struct {
	CompanyId             int       `json:"companyId"`
	CreatedAt             time.Time `json:"createdAt"`
	GeneralScore          int       `json:"generalScore"`
	IsWeb                 bool      `json:"isWeb"`
	Name                  string    `json:"name"`
	NewsletterScore       int       `json:"newsletterScore"`
	Priority              int       `json:"priority"`
	TotalActiveSubscriber int       `json:"totalActiveSubscriber"`
	TotalBounced          int       `json:"totalBounced"`
	TotalSubscriber       int       `json:"totalSubscriber"`
	TotalUnsubscribe      int       `json:"totalUnsubscribe"`
	UpdatedAt             time.Time `json:"updatedAt"`
	UUID                  string    `json:"uuid"`
}

// IPGroup represents a group of IP addresses.
type IPGroup struct {
	CompanyId   int         `json:"companyId"`
	CreatedAt   time.Time   `json:"createdAt"`
	IpAddresses []IpAddress `json:"ipAddresses"`
	Name        string      `json:"name"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	UUID        string      `json:"uuid"`
}

// IpAddress represents details of an IP address.
type IpAddress struct {
	CompanyId  int       `json:"companyId"`
	CreatedAt  time.Time `json:"createdAt"`
	Ip         string    `json:"ip"`
	IpGroupId  int       `json:"ipGroupId"`
	ReverseDNS string    `json:"reverseDNS,omitempty"`
	Status     string    `json:"status"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UUID       string    `json:"uuid"`
}

// Subscriber represents an email subscriber.
type Subscriber struct {
	Bounced      bool          `json:"bounced"`
	ConfirmCode  string        `json:"confirmCode"`
	ConfirmIp    string        `json:"confirmIp"`
	Confirmed    bool          `json:"confirmed"`
	CreatedAt    time.Time     `json:"createdAt"`
	CustomFields []CustomField `json:"customFields"`
	EmailAddress string        `json:"emailAddress"`
	EmailGroupId int           `json:"emailGroupId"`
	SubscribedAt time.Time     `json:"subscribedAt"`
	Unsubscribed bool          `json:"unsubscribed"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	UUID         string        `json:"uuid"`
}

// CustomField represents a custom field for a subscriber.
type CustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CompanyDomain represents a domain associated with a company.
type CompanyDomain struct {
	AwsRegion      string    `json:"awsRegion"`
	AwsVerified    bool      `json:"awsVerified"`
	Company        Company   `json:"company"`
	CompanyId      int       `json:"companyId"`
	CreatedAt      time.Time `json:"createdAt"`
	DkimContent    string    `json:"dkimContent"`
	DkimName       string    `json:"dkimName"`
	DkimPrivateKey string    `json:"dkimPrivateKey"`
	DkimSelector   string    `json:"dkimSelector"`
	DkimVerified   bool      `json:"dkimVerified"`
	DmarcContent   string    `json:"dmarcContent"`
	DmarcName      string    `json:"dmarcName"`
	DmarcVerified  bool      `json:"dmarcVerified"`
	Domain         string    `json:"domain"`
	HasAwsIdentity bool      `json:"hasAwsIdentity"`
	IsVerified     bool      `json:"isVerified"`
	SpfContent     string    `json:"spfContent"`
	SpfName        string    `json:"spfName"`
	SpfVerified    bool      `json:"spfVerified"`
	UpdatedAt      time.Time `json:"updatedAt"`
	UUID           string    `json:"uuid"`
}

// Company represents a company.
type Company struct {
	CompanyPlan   CompanyPlan `json:"companyPlan"`
	CompanyPlanID int         `json:"companyPlanID"`
	CreatedAt     time.Time   `json:"createdAt"`
	FooterHtml    string      `json:"footerHtml"`
	FooterText    string      `json:"footerText"`
	Name          string      `json:"name"`
	OwnerId       int         `json:"ownerId"`
	Priority      int         `json:"priority"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	UUID          string      `json:"uuid"`
}

// CompanyPlan represents a plan associated with a company.
type CompanyPlan struct {
	CompanyId             int         `json:"companyID"`
	CreatedAt             time.Time   `json:"createdAt"`
	CurrentUsage          int         `json:"currentUsage"`
	EndedAt               time.Time   `json:"endedAt"`
	LastBilled            time.Time   `json:"lastBilled"`
	PricingPlan           PricingPlan `json:"pricingPlan"`
	PricingPlanId         int         `json:"pricingPlanID"`
	SelectedContactsLimit int         `json:"selectedContactsLimit"`
	SelectedDataRetention int         `json:"selectedDataRetention"`
	SelectedEmailLimit    int         `json:"selectedEmailLimit"`
	StartedAt             time.Time   `json:"startedAt"`
	Status                string      `json:"status"`
	UpdatedAt             time.Time   `json:"updatedAt"`
	UUID                  string      `json:"uuid"`
}

// PricingPlan represents a pricing plan.
type PricingPlan struct {
	CreatedAt    time.Time `json:"createdAt"`
	DailyLimit   int       `json:"dailyLimit"`
	MaximumEmail int       `json:"maximumEmail"`
	Name         string    `json:"name"`
	PlanType     string    `json:"planType"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UUID         string    `json:"uuid"`
}
