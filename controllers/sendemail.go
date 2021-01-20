package controllers

import (
	"communication/models"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

// SendEmail is used to send email to customers
// Supports To, CC, BCC to a maximum of 10
func SendEmail(c echo.Context) (err error) {
	log.Println("Send email request received...")

	var errorResponse models.ErrorResponse
	request := new(models.SendEmailRequest)

	if err = c.Bind(request); err != nil {
		log.Println(fmt.Sprintf("Error occured while trying to marshal request: %s", err))
		return
	}

	// decoder := json.NewDecoder(r.Body)
	// err = decoder.Decode(&request)
	// defer r.Body.Close()
	if len(request.To) < 1 || request.From.Email == "" {
		log.Println("Invalid request, From and To must have a value")
		errorResponse.Errorcode = "03"
		errorResponse.ErrorMessage = "Invalid request, From and To must have a value"

		// response, err := json.MarshalIndent(errorResponse, "", "")
		// if err != nil {
		// 	log.Println(err)
		// }
		c.JSON(http.StatusBadRequest, errorResponse)
		// w.WriteHeader(http.StatusUnauthorized)
		// w.Write(response)
		return
	}
	// Initialise gomail library
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(request.From.Email, request.From.Name))
	// Set To Email Addresses to a maximum of 10
	if len(request.To) == 1 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name))
	} else if len(request.To) == 2 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name))
	} else if len(request.To) == 3 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name))
	} else if len(request.To) == 4 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name), m.FormatAddress(request.To[3].Email, request.To[3].Name))
	} else if len(request.To) == 5 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name), m.FormatAddress(request.To[3].Email, request.To[3].Name), m.FormatAddress(request.To[4].Email, request.To[4].Name))
	} else if len(request.To) == 6 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name), m.FormatAddress(request.To[3].Email, request.To[3].Name), m.FormatAddress(request.To[4].Email, request.To[4].Name), m.FormatAddress(request.To[5].Email, request.To[5].Name))
	} else if len(request.To) == 7 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name), m.FormatAddress(request.To[3].Email, request.To[3].Name), m.FormatAddress(request.To[4].Email, request.To[4].Name), m.FormatAddress(request.To[5].Email, request.To[5].Name), m.FormatAddress(request.To[6].Email, request.To[6].Name))
	} else if len(request.To) == 8 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name), m.FormatAddress(request.To[3].Email, request.To[3].Name), m.FormatAddress(request.To[4].Email, request.To[4].Name), m.FormatAddress(request.To[5].Email, request.To[5].Name), m.FormatAddress(request.To[6].Email, request.To[6].Name), m.FormatAddress(request.To[7].Email, request.To[7].Name))
	} else if len(request.To) == 9 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name), m.FormatAddress(request.To[3].Email, request.To[3].Name), m.FormatAddress(request.To[4].Email, request.To[4].Name), m.FormatAddress(request.To[5].Email, request.To[5].Name), m.FormatAddress(request.To[6].Email, request.To[6].Name), m.FormatAddress(request.To[7].Email, request.To[7].Name), m.FormatAddress(request.To[8].Email, request.To[8].Name))
	} else if len(request.To) >= 10 {
		m.SetHeader("To", m.FormatAddress(request.To[0].Email, request.To[0].Name), m.FormatAddress(request.To[1].Email, request.To[1].Name), m.FormatAddress(request.To[2].Email, request.To[2].Name), m.FormatAddress(request.To[3].Email, request.To[3].Name), m.FormatAddress(request.To[4].Email, request.To[4].Name), m.FormatAddress(request.To[5].Email, request.To[5].Name), m.FormatAddress(request.To[6].Email, request.To[6].Name), m.FormatAddress(request.To[7].Email, request.To[7].Name), m.FormatAddress(request.To[8].Email, request.To[8].Name), m.FormatAddress(request.To[9].Email, request.To[9].Name))
	}

	// set cc maximum of 10

	if len(request.CC) == 1 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name))
	} else if len(request.CC) == 2 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name))
	} else if len(request.CC) == 3 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name))
	} else if len(request.CC) == 4 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name), m.FormatAddress(request.CC[3].Email, request.CC[3].Name))
	} else if len(request.CC) == 5 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name), m.FormatAddress(request.CC[3].Email, request.CC[3].Name), m.FormatAddress(request.CC[4].Email, request.CC[4].Name))
	} else if len(request.CC) == 6 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name), m.FormatAddress(request.CC[3].Email, request.CC[3].Name), m.FormatAddress(request.CC[4].Email, request.CC[4].Name), m.FormatAddress(request.CC[5].Email, request.CC[5].Name))
	} else if len(request.CC) == 7 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name), m.FormatAddress(request.CC[3].Email, request.CC[3].Name), m.FormatAddress(request.CC[4].Email, request.CC[4].Name), m.FormatAddress(request.CC[5].Email, request.CC[5].Name), m.FormatAddress(request.CC[6].Email, request.CC[6].Name))
	} else if len(request.CC) == 8 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name), m.FormatAddress(request.CC[3].Email, request.CC[3].Name), m.FormatAddress(request.CC[4].Email, request.CC[4].Name), m.FormatAddress(request.CC[5].Email, request.CC[5].Name), m.FormatAddress(request.CC[6].Email, request.CC[6].Name), m.FormatAddress(request.CC[7].Email, request.CC[7].Name))
	} else if len(request.CC) == 9 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name), m.FormatAddress(request.CC[3].Email, request.CC[3].Name), m.FormatAddress(request.CC[4].Email, request.CC[4].Name), m.FormatAddress(request.CC[5].Email, request.CC[5].Name), m.FormatAddress(request.CC[6].Email, request.CC[6].Name), m.FormatAddress(request.CC[7].Email, request.CC[7].Name), m.FormatAddress(request.CC[8].Email, request.CC[8].Name))
	} else if len(request.CC) >= 10 {
		m.SetHeader("Cc", m.FormatAddress(request.CC[0].Email, request.CC[0].Name), m.FormatAddress(request.CC[1].Email, request.CC[1].Name), m.FormatAddress(request.CC[2].Email, request.CC[2].Name), m.FormatAddress(request.CC[3].Email, request.CC[3].Name), m.FormatAddress(request.CC[4].Email, request.CC[4].Name), m.FormatAddress(request.CC[5].Email, request.CC[5].Name), m.FormatAddress(request.CC[6].Email, request.CC[6].Name), m.FormatAddress(request.CC[7].Email, request.CC[7].Name), m.FormatAddress(request.CC[8].Email, request.CC[8].Name), m.FormatAddress(request.CC[9].Email, request.CC[9].Name))
	}

	//set Bcc maximum of 10

	if len(request.BCC) == 1 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name))
	} else if len(request.BCC) == 2 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name))
	} else if len(request.BCC) == 3 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name))
	} else if len(request.BCC) == 4 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name), m.FormatAddress(request.BCC[3].Email, request.BCC[3].Name))
	} else if len(request.BCC) == 5 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name), m.FormatAddress(request.BCC[3].Email, request.BCC[3].Name), m.FormatAddress(request.BCC[4].Email, request.BCC[4].Name))
	} else if len(request.BCC) == 6 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name), m.FormatAddress(request.BCC[3].Email, request.BCC[3].Name), m.FormatAddress(request.BCC[4].Email, request.BCC[4].Name), m.FormatAddress(request.BCC[5].Email, request.BCC[5].Name))
	} else if len(request.BCC) == 7 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name), m.FormatAddress(request.BCC[3].Email, request.BCC[3].Name), m.FormatAddress(request.BCC[4].Email, request.BCC[4].Name), m.FormatAddress(request.BCC[5].Email, request.BCC[5].Name), m.FormatAddress(request.BCC[6].Email, request.BCC[6].Name))
	} else if len(request.BCC) == 8 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name), m.FormatAddress(request.BCC[3].Email, request.BCC[3].Name), m.FormatAddress(request.BCC[4].Email, request.BCC[4].Name), m.FormatAddress(request.BCC[5].Email, request.BCC[5].Name), m.FormatAddress(request.BCC[6].Email, request.BCC[6].Name), m.FormatAddress(request.BCC[7].Email, request.BCC[7].Name))
	} else if len(request.BCC) == 9 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name), m.FormatAddress(request.BCC[3].Email, request.BCC[3].Name), m.FormatAddress(request.BCC[4].Email, request.BCC[4].Name), m.FormatAddress(request.BCC[5].Email, request.BCC[5].Name), m.FormatAddress(request.BCC[6].Email, request.BCC[6].Name), m.FormatAddress(request.BCC[7].Email, request.BCC[7].Name), m.FormatAddress(request.BCC[8].Email, request.BCC[8].Name))
	} else if len(request.BCC) >= 10 {
		m.SetHeader("Bcc", m.FormatAddress(request.BCC[0].Email, request.BCC[0].Name), m.FormatAddress(request.BCC[1].Email, request.BCC[1].Name), m.FormatAddress(request.BCC[2].Email, request.BCC[2].Name), m.FormatAddress(request.BCC[3].Email, request.BCC[3].Name), m.FormatAddress(request.BCC[4].Email, request.BCC[4].Name), m.FormatAddress(request.BCC[5].Email, request.BCC[5].Name), m.FormatAddress(request.BCC[6].Email, request.BCC[6].Name), m.FormatAddress(request.BCC[7].Email, request.BCC[7].Name), m.FormatAddress(request.BCC[8].Email, request.BCC[8].Name), m.FormatAddress(request.BCC[9].Email, request.BCC[9].Name))
	}

	m.SetHeader("Subject", request.Subject)
	m.SetBody("text/html", request.Message)

	go func() {
		var smtpHost, smtpPortKey, smtpUser, smtpPassword, attachmentPath string
		smtpHost = os.Getenv("SMTP_HOST")
		smtpPortKey = os.Getenv("SMTP_PORT")
		smtpUser = os.Getenv("SMTP_USER")
		smtpPassword = os.Getenv("SMTP_PASSWORD")
		attachmentPath = os.Getenv("ATTACHMENT_PATH")
		smtpPort, err := strconv.Atoi(smtpPortKey)
		if err != nil {
			log.Println(fmt.Sprintf("Invalid port number passed: %s", err))
		}
		for i := 0; i < len(request.AttachmentName); i++ {

			log.Println(fmt.Sprintf("Attachment: %s was found....", request.AttachmentName[i].FileName))
			file, err := os.Create(fmt.Sprintf("%s%s", attachmentPath, request.AttachmentName[i].FileName))

			if err != nil {
				log.Println(fmt.Sprintf("Error occured while creating file on host: %s", err))
			}

			defer file.Close()
			dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(request.AttachmentName[i].Base64))

			fileSize, err := io.Copy(file, dec)
			if err != nil {
				log.Println(fmt.Sprintf("Error occured uploading attachment %s to server: %s", request.AttachmentName[i].FileName, err))
			}

			log.Println(fmt.Sprintf("Attachment %s Uploaded successfully. Wrote %d bytes", request.AttachmentName[i].FileName, fileSize))

			m.Attach(fmt.Sprintf("%s%s", attachmentPath, request.AttachmentName[i].FileName))
		}

		d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
		}
		log.Println(fmt.Sprintf("Successfully sent email for %s", request.From))
	}()
	successResponse := &models.SuccessResponse{
		ResponseCode:        "00",
		ResponseDescription: "Email received for sending...",
		ResponseMessage:     nil,
	}
	log.Println("Email sent successfully...")
	// response, err := json.MarshalIndent(successResponse, "", "")
	// if err != nil {
	// 	log.Println(err)
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(response)
	c.JSON(http.StatusOK, successResponse)
	return
}

// SendNewsletter is used to send email to customers
// Supports unlimited To addresses. CC and BCC are not supported
func SendNewsletter(c echo.Context) (err error) {
	log.Println("Send newsletter request received...")
	// var err error
	var errorResponse models.ErrorResponse
	request := new(models.SendEmailRequest)
	// decoder := json.NewDecoder(r.Body)
	// err = decoder.Decode(&request)
	// defer r.Body.Close()
	if err = c.Bind(request); err != nil {
		log.Println(fmt.Sprintf("Error occured while trying to marshal request: %s", err))
		return
	}
	if len(request.To) < 1 || request.From.Email == "" {
		log.Println("Invalid request, From and To must have a value")
		errorResponse.Errorcode = "03"
		errorResponse.ErrorMessage = "Invalid request, From and To must have a value"

		// response, err := json.MarshalIndent(errorResponse, "", "")
		// if err != nil {
		// 	log.Println(err)
		// }
		// w.WriteHeader(http.StatusUnauthorized)
		// w.Write(response)
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	go func() {
		var smtpHost, smtpPortKey, smtpUser, smtpPassword, attachmentPath string
		smtpHost = os.Getenv("SMTP_HOST")
		smtpPortKey = os.Getenv("SMTP_PORT")
		smtpUser = os.Getenv("SMTP_USER")
		smtpPassword = os.Getenv("SMTP_PASSWORD")
		attachmentPath = os.Getenv("ATTACHMENT_PATH")
		smtpPort, err := strconv.Atoi(smtpPortKey)
		if err != nil {
			log.Println(fmt.Sprintf("Invalid port number passed: %s", err))
		}
		// d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)
		// s, err := d.Dial()
		// if err != nil {
		// 	log.Println(err)
		// }

		m := gomail.NewMessage()

		for _, recipient := range request.To {
			m.SetAddressHeader("From", request.From.Email, request.From.Name)
			m.SetAddressHeader("To", recipient.Email, recipient.Name)
			m.SetHeader("Subject", request.Subject)
			m.SetBody("text/html", request.Message)

			for i := 0; i < len(request.AttachmentName); i++ {

				log.Println(fmt.Sprintf("Attachment: %s was found....", request.AttachmentName[i].FileName))
				file, err := os.Create(fmt.Sprintf("%s%s", attachmentPath, request.AttachmentName[i].FileName))

				if err != nil {
					log.Println(fmt.Sprintf("Error occured while creating file on host: %s", err))
				}

				defer file.Close()
				dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(request.AttachmentName[i].Base64))

				fileSize, err := io.Copy(file, dec)
				if err != nil {
					log.Println(fmt.Sprintf("Error occured uploading attachment %s to server: %s", request.AttachmentName[i].FileName, err))
				}

				log.Println(fmt.Sprintf("Attachment %s Uploaded successfully. Wrote %d bytes", request.AttachmentName[i].FileName, fileSize))

				m.Attach(fmt.Sprintf("%s%s", attachmentPath, request.AttachmentName[i].FileName))
			}

			// if err := gomail.Send(s, m); err != nil {
			// 	log.Printf("Could not send Newsletter: %s to %s: %v", request.Subject, recipient.Email, err)
			// }
			d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

			// Send the email to Bob, Cora and Dan.
			if err := d.DialAndSend(m); err != nil {
				log.Println(err)
			}
			m.Reset()
		}
	}()
	successResponse := &models.SuccessResponse{
		ResponseCode:        "00",
		ResponseDescription: "Newsletter received for sending...",
		ResponseMessage:     nil,
	}
	log.Println("Email sent successfully...")
	// response, err := json.MarshalIndent(successResponse, "", "")
	// if err != nil {
	// 	log.Println(err)
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(response)
	c.JSON(http.StatusOK, successResponse)
	return
}
