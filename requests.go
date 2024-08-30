package mepost

// AddDomainRequest represents the request to add a domain.
type AddDomainRequest struct {
	Domain string `json:"domain"`
}

// RemoveDomainRequest represents the request to add a domain.
type RemoveDomainRequest struct {
	Domain string `json:"domain"`
}

// CancelScheduledMessageRequest represents the request to cancel a scheduled message.
type CancelScheduledMessageRequest struct {
	ScheduledMessageID string `json:"scheduledMessageId"`
}

// CancelWarmUpRequest represents the request to cancel IP warm-up.
type CancelWarmUpRequest struct {
	IpAddress string `json:"ipAddress"`
}

// CreateIpGroupRequest represents the request to create a new IP group.
type CreateIpGroupRequest struct {
	GroupName string `json:"groupName"`
}

// CreateNewGroupRequest represents the request to create a new email group.
type CreateNewGroupRequest struct {
	Name string `json:"name"`
	To   []To   `json:"to"`
}

// CreateSubscriberRequest represents the request to add subscribers to a group.
type CreateSubscriberRequest struct {
	To []To `json:"to"`
}

// DeleteSubscriberRequest represents the request to delete subscribers.
type DeleteSubscriberRequest struct {
	Emails []string `json:"emails"`
}

// RenameGroupRequest represents the request to rename a group.
type RenameGroupRequest struct {
	Name string `json:"name"`
}

// SendMarketingRequest represents the request to send a marketing email.
type SendMarketingRequest struct {
	Attachments   []AttachmentDto   `json:"attachments,omitempty"`
	Customization map[string]string `json:"customization,omitempty"`
	FromEmail     string            `json:"fromEmail"`
	FromName      string            `json:"fromName"`
	Headers       map[string]string `json:"headers,omitempty"`
	Html          string            `json:"html,omitempty"`
	IpGroup       string            `json:"ipGroup,omitempty"`
	ReturnPath    string            `json:"returnPath,omitempty"`
	ScheduledAt   string            `json:"scheduledAt,omitempty"`
	Subject       string            `json:"subject"`
	Text          string            `json:"text,omitempty"`
	To            []string          `json:"to"`
}

// SendMessageByTemplateRequest represents the request to send a message by template.
type SendMessageByTemplateRequest struct {
	Message    MessageDto `json:"message"`
	TemplateID string     `json:"templateId"`
}

// SendTransactionalRequest represents the request to send a transactional email.
type SendTransactionalRequest struct {
	Attachments   []AttachmentDto   `json:"attachments,omitempty"`
	Customization map[string]string `json:"customization,omitempty"`
	FromEmail     string            `json:"fromEmail"`
	FromName      string            `json:"fromName"`
	Headers       map[string]string `json:"headers,omitempty"`
	Html          string            `json:"html,omitempty"`
	IpGroup       string            `json:"ipGroup,omitempty"`
	ReturnPath    string            `json:"returnPath,omitempty"`
	ScheduledAt   string            `json:"scheduledAt,omitempty"`
	Subject       string            `json:"subject"`
	Text          string            `json:"text,omitempty"`
	To            []To              `json:"to"`
}

// SetIpGroupRequest represents the request to set an IP group.
type SetIpGroupRequest struct {
	GroupName string `json:"groupName"`
	IpAddress string `json:"ipAddress"`
}

// StartWarmUpRequest represents the request to start an IP warm-up process.
type StartWarmUpRequest struct {
	IpAddress string `json:"ipAddress"`
}

// To represents the structure for recipient details.
type To struct {
	Customization map[string]string `json:"customization,omitempty"`
	Email         string            `json:"email"`
	Name          string            `json:"name"`
	Type          string            `json:"type,omitempty"`
}

// AttachmentDto represents the structure for an email attachment.
type AttachmentDto struct {
	Base64Content string `json:"base64Content"`
	FileName      string `json:"fileName"`
}

// MessageDto represents the structure for a message.
type MessageDto struct {
	Attachments   []AttachmentDto   `json:"attachments,omitempty"`
	Customization map[string]string `json:"customization,omitempty"`
	FromEmail     string            `json:"fromEmail"`
	FromName      string            `json:"fromName"`
	Headers       map[string]string `json:"headers,omitempty"`
	Html          string            `json:"html,omitempty"`
	IpGroup       string            `json:"ipGroup,omitempty"`
	ReturnPath    string            `json:"returnPath,omitempty"`
	ScheduledAt   string            `json:"scheduledAt,omitempty"`
	Subject       string            `json:"subject"`
	Text          string            `json:"text,omitempty"`
	To            []To              `json:"to"`
}
