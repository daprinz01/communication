package util

import (
	"io"
)

type file struct {
	Name     string
	Header   map[string][]string
	CopyFunc func(w io.Writer) error
}

// SendEmail is used to send email to users
// func SendEmail(emailMessage models.SendEmailRequest) error {
// 	//Read variables from environment
// 	var smtpHost, smtpPortKey, smtpUser, smtpPassword string
// 	smtpHost = os.Getenv("SMTP_HOST")
// 	smtpPortKey = os.Getenv("SMTP_PORT")
// 	smtpUser = os.Getenv("SMTP_USER")
// 	smtpPassword = os.Getenv("SMTP_PASSWORD")

// 	smtpPort, err := strconv.Atoi(smtpPortKey)
// 	if err != nil {
// 		log.Println(fmt.Sprintf("Invalid port number passed: %s", err))
// 	}
// 	// // Set up authentication information.
// 	// auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

// 	// // Connect to the server, authenticate, set the sender and recipient,
// 	// // and send the email all in one step.
// 	// // to := []string{"recipient@example.net"}
// 	// msg := []byte(message)
// 	// err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, from, to, msg)
// 	// if err != nil {
// 	// 	log.Println(fmt.Sprintf("Error occured sending Email: err"))
// 	// 	return err
// 	// }
// 	// return nil

// 	m := gomail.NewMessage()
// 	m.SetHeader("From", fmt.Sprintf("%s <%s>", emailMessage.From.Name, emailMessage.From.Email))
// 	// m.SetHeader("To", "okechukwuprince@hotmail.com", "cora@example.com")
// 	// Set To Email Addresses to a maximum of 10
// 	if len(emailMessage.To) == 1 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name))
// 	} else if len(emailMessage.To) == 2 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name))
// 	} else if len(emailMessage.To) == 3 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name))
// 	} else if len(emailMessage.To) == 4 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name), m.FormatAddress(emailMessage.To[3].Email, emailMessage.To[3].Name))
// 	} else if len(emailMessage.To) == 5 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name), m.FormatAddress(emailMessage.To[3].Email, emailMessage.To[3].Name), m.FormatAddress(emailMessage.To[4].Email, emailMessage.To[4].Name))
// 	} else if len(emailMessage.To) == 6 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name), m.FormatAddress(emailMessage.To[3].Email, emailMessage.To[3].Name), m.FormatAddress(emailMessage.To[4].Email, emailMessage.To[4].Name), m.FormatAddress(emailMessage.To[5].Email, emailMessage.To[5].Name))
// 	} else if len(emailMessage.To) == 7 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name), m.FormatAddress(emailMessage.To[3].Email, emailMessage.To[3].Name), m.FormatAddress(emailMessage.To[4].Email, emailMessage.To[4].Name), m.FormatAddress(emailMessage.To[5].Email, emailMessage.To[5].Name), m.FormatAddress(emailMessage.To[6].Email, emailMessage.To[6].Name))
// 	} else if len(emailMessage.To) == 8 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name), m.FormatAddress(emailMessage.To[3].Email, emailMessage.To[3].Name), m.FormatAddress(emailMessage.To[4].Email, emailMessage.To[4].Name), m.FormatAddress(emailMessage.To[5].Email, emailMessage.To[5].Name), m.FormatAddress(emailMessage.To[6].Email, emailMessage.To[6].Name), m.FormatAddress(emailMessage.To[7].Email, emailMessage.To[7].Name))
// 	} else if len(emailMessage.To) == 9 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name), m.FormatAddress(emailMessage.To[3].Email, emailMessage.To[3].Name), m.FormatAddress(emailMessage.To[4].Email, emailMessage.To[4].Name), m.FormatAddress(emailMessage.To[5].Email, emailMessage.To[5].Name), m.FormatAddress(emailMessage.To[6].Email, emailMessage.To[6].Name), m.FormatAddress(emailMessage.To[7].Email, emailMessage.To[7].Name), m.FormatAddress(emailMessage.To[8].Email, emailMessage.To[8].Name))
// 	} else if len(emailMessage.To) >= 10 {
// 		m.SetHeader("To", m.FormatAddress(emailMessage.To[0].Email, emailMessage.To[0].Name), m.FormatAddress(emailMessage.To[1].Email, emailMessage.To[1].Name), m.FormatAddress(emailMessage.To[2].Email, emailMessage.To[2].Name), m.FormatAddress(emailMessage.To[3].Email, emailMessage.To[3].Name), m.FormatAddress(emailMessage.To[4].Email, emailMessage.To[4].Name), m.FormatAddress(emailMessage.To[5].Email, emailMessage.To[5].Name), m.FormatAddress(emailMessage.To[6].Email, emailMessage.To[6].Name), m.FormatAddress(emailMessage.To[7].Email, emailMessage.To[7].Name), m.FormatAddress(emailMessage.To[8].Email, emailMessage.To[8].Name), m.FormatAddress(emailMessage.To[9].Email, emailMessage.To[9].Name))
// 	}

// 	// set cc maximum of 10

// 	if len(emailMessage.CC) == 1 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name))
// 	} else if len(emailMessage.CC) == 2 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name))
// 	} else if len(emailMessage.CC) == 3 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name))
// 	} else if len(emailMessage.CC) == 4 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name), m.FormatAddress(emailMessage.CC[3].Email, emailMessage.CC[3].Name))
// 	} else if len(emailMessage.CC) == 5 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name), m.FormatAddress(emailMessage.CC[3].Email, emailMessage.CC[3].Name), m.FormatAddress(emailMessage.CC[4].Email, emailMessage.CC[4].Name))
// 	} else if len(emailMessage.CC) == 6 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name), m.FormatAddress(emailMessage.CC[3].Email, emailMessage.CC[3].Name), m.FormatAddress(emailMessage.CC[4].Email, emailMessage.CC[4].Name), m.FormatAddress(emailMessage.CC[5].Email, emailMessage.CC[5].Name))
// 	} else if len(emailMessage.CC) == 7 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name), m.FormatAddress(emailMessage.CC[3].Email, emailMessage.CC[3].Name), m.FormatAddress(emailMessage.CC[4].Email, emailMessage.CC[4].Name), m.FormatAddress(emailMessage.CC[5].Email, emailMessage.CC[5].Name), m.FormatAddress(emailMessage.CC[6].Email, emailMessage.CC[6].Name))
// 	} else if len(emailMessage.CC) == 8 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name), m.FormatAddress(emailMessage.CC[3].Email, emailMessage.CC[3].Name), m.FormatAddress(emailMessage.CC[4].Email, emailMessage.CC[4].Name), m.FormatAddress(emailMessage.CC[5].Email, emailMessage.CC[5].Name), m.FormatAddress(emailMessage.CC[6].Email, emailMessage.CC[6].Name), m.FormatAddress(emailMessage.CC[7].Email, emailMessage.CC[7].Name))
// 	} else if len(emailMessage.CC) == 9 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name), m.FormatAddress(emailMessage.CC[3].Email, emailMessage.CC[3].Name), m.FormatAddress(emailMessage.CC[4].Email, emailMessage.CC[4].Name), m.FormatAddress(emailMessage.CC[5].Email, emailMessage.CC[5].Name), m.FormatAddress(emailMessage.CC[6].Email, emailMessage.CC[6].Name), m.FormatAddress(emailMessage.CC[7].Email, emailMessage.CC[7].Name), m.FormatAddress(emailMessage.CC[8].Email, emailMessage.CC[8].Name))
// 	} else if len(emailMessage.CC) >= 10 {
// 		m.SetHeader("Cc", m.FormatAddress(emailMessage.CC[0].Email, emailMessage.CC[0].Name), m.FormatAddress(emailMessage.CC[1].Email, emailMessage.CC[1].Name), m.FormatAddress(emailMessage.CC[2].Email, emailMessage.CC[2].Name), m.FormatAddress(emailMessage.CC[3].Email, emailMessage.CC[3].Name), m.FormatAddress(emailMessage.CC[4].Email, emailMessage.CC[4].Name), m.FormatAddress(emailMessage.CC[5].Email, emailMessage.CC[5].Name), m.FormatAddress(emailMessage.CC[6].Email, emailMessage.CC[6].Name), m.FormatAddress(emailMessage.CC[7].Email, emailMessage.CC[7].Name), m.FormatAddress(emailMessage.CC[8].Email, emailMessage.CC[8].Name), m.FormatAddress(emailMessage.CC[9].Email, emailMessage.CC[9].Name))
// 	}

// 	//set Bcc maximum of 10

// 	if len(emailMessage.BCC) == 1 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name))
// 	} else if len(emailMessage.BCC) == 2 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name))
// 	} else if len(emailMessage.BCC) == 3 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name))
// 	} else if len(emailMessage.BCC) == 4 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name), m.FormatAddress(emailMessage.BCC[3].Email, emailMessage.BCC[3].Name))
// 	} else if len(emailMessage.BCC) == 5 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name), m.FormatAddress(emailMessage.BCC[3].Email, emailMessage.BCC[3].Name), m.FormatAddress(emailMessage.BCC[4].Email, emailMessage.BCC[4].Name))
// 	} else if len(emailMessage.BCC) == 6 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name), m.FormatAddress(emailMessage.BCC[3].Email, emailMessage.BCC[3].Name), m.FormatAddress(emailMessage.BCC[4].Email, emailMessage.BCC[4].Name), m.FormatAddress(emailMessage.BCC[5].Email, emailMessage.BCC[5].Name))
// 	} else if len(emailMessage.BCC) == 7 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name), m.FormatAddress(emailMessage.BCC[3].Email, emailMessage.BCC[3].Name), m.FormatAddress(emailMessage.BCC[4].Email, emailMessage.BCC[4].Name), m.FormatAddress(emailMessage.BCC[5].Email, emailMessage.BCC[5].Name), m.FormatAddress(emailMessage.BCC[6].Email, emailMessage.BCC[6].Name))
// 	} else if len(emailMessage.BCC) == 8 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name), m.FormatAddress(emailMessage.BCC[3].Email, emailMessage.BCC[3].Name), m.FormatAddress(emailMessage.BCC[4].Email, emailMessage.BCC[4].Name), m.FormatAddress(emailMessage.BCC[5].Email, emailMessage.BCC[5].Name), m.FormatAddress(emailMessage.BCC[6].Email, emailMessage.BCC[6].Name), m.FormatAddress(emailMessage.BCC[7].Email, emailMessage.BCC[7].Name))
// 	} else if len(emailMessage.BCC) == 9 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name), m.FormatAddress(emailMessage.BCC[3].Email, emailMessage.BCC[3].Name), m.FormatAddress(emailMessage.BCC[4].Email, emailMessage.BCC[4].Name), m.FormatAddress(emailMessage.BCC[5].Email, emailMessage.BCC[5].Name), m.FormatAddress(emailMessage.BCC[6].Email, emailMessage.BCC[6].Name), m.FormatAddress(emailMessage.BCC[7].Email, emailMessage.BCC[7].Name), m.FormatAddress(emailMessage.BCC[8].Email, emailMessage.BCC[8].Name))
// 	} else if len(emailMessage.BCC) >= 10 {
// 		m.SetHeader("Bcc", m.FormatAddress(emailMessage.BCC[0].Email, emailMessage.BCC[0].Name), m.FormatAddress(emailMessage.BCC[1].Email, emailMessage.BCC[1].Name), m.FormatAddress(emailMessage.BCC[2].Email, emailMessage.BCC[2].Name), m.FormatAddress(emailMessage.BCC[3].Email, emailMessage.BCC[3].Name), m.FormatAddress(emailMessage.BCC[4].Email, emailMessage.BCC[4].Name), m.FormatAddress(emailMessage.BCC[5].Email, emailMessage.BCC[5].Name), m.FormatAddress(emailMessage.BCC[6].Email, emailMessage.BCC[6].Name), m.FormatAddress(emailMessage.BCC[7].Email, emailMessage.BCC[7].Name), m.FormatAddress(emailMessage.BCC[8].Email, emailMessage.BCC[8].Name), m.FormatAddress(emailMessage.BCC[9].Email, emailMessage.BCC[9].Name))
// 	}
// 	// m.FormatAddress("testmail1@yahoo.com", "testMail1"),
// 	// m.FormatAddress("testmail2@yahoo.com", "testMail2"),

// 	// m.SetHeader("Bcc", "Persian Support <support@persianblack.com>")
// 	// m.SetAddressHeader("Cc", "daprinz.op@gmail.com", "Prince Gmail")
// 	m.SetHeader("Subject", emailMessage.Subject)
// 	m.SetBody("text/html", emailMessage.Message)
// 	m.Attach("/usr/local/bin/log/communication.log")

// 	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

// 	// Send the email to Bob, Cora and Dan.
// 	if err := d.DialAndSend(m); err != nil {
// 		return err
// 	}
// 	return nil
// }

// SendNewsletter used to send newsletter
func SendNewsletter() {
	// // The list of recipients.
	// var list models.NewsLetterList

	// d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")
	// s, err := d.Dial()
	// if err != nil {
	// 	panic(err)
	// }

	// m := gomail.NewMessage()
	// for _, r := range list {
	// 	m.SetHeader("From", "no-reply@example.com")
	// 	m.SetAddressHeader("To", "okechukwuprince@hotmail.com")
	// 	m.SetHeader("Subject", "Newsletter #1")
	// 	m.SetBody("text/html", fmt.Sprintf("Hello Prince!"))

	// 	if err := gomail.Send(s, m); err != nil {
	// 		log.Printf("Could not send email to okechukwuprince@hotmail.com: %v", err)
	// 	}
	// 	m.Reset()
	// }
}
