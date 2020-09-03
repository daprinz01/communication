package models

// SendEmailRequest is used to get email request
type SendEmailRequest struct {
	From           EmailAddress      `json:"from"`
	To             []EmailAddress    `json:"to"`
	CC             []EmailAddress    `json:"cc"`
	BCC            []EmailAddress    `json:"bcc"`
	Subject        string            `json:"subject"`
	Message        string            `json:"message"`
	AttachmentName []EmailAttachment `json:"attachments"`
}

// NewsLetterList is used to send newsletter to multiple recipients
type NewsLetterList []struct {
	Name    string
	Address string
}

// EmailAddress is used to collect name and email of recipients
type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// EmailAttachment is used to send attachments
type EmailAttachment struct {
	FileName string `json:"fileName"`
	Base64   string `json:"base64"`
}

// OtpRequest is used to collect name and email of recipients
type OtpRequest struct {
	Email       string `json:"email"`
	Purpose     string `json:"purpose"`
	Application string `json:"application"`
}

// SendSmsRequest is used to receive sms request
type SendSmsRequest struct {
	Phone   string `json:"phonenumber"`
	Message string `json:"message"`
}
