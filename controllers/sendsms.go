package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"persianblack.com/communication/models"

	"github.com/labstack/echo/v4"
)

// SendSMS is used to send SMS messages
func SendSMS(c echo.Context) (err error) {
	fields := log.Fields{"microservice": "persian.black.devtroy.communication.service", "function": "SendNewsletter", "application": "communication"}
	log.WithFields(fields).Info("Send sms request received...")
	// var err error
	var errorResponse models.ErrorResponse
	request := new(models.SendSmsRequest)
	if err = c.Bind(request); err != nil {
		errorResponse.Errorcode = "03"
		errorResponse.ErrorMessage = "Invalid request, phonenumber must have a value"
		log.WithFields(fields).WithError(err).WithFields(log.Fields{"responseCode": errorResponse.Errorcode, "responseDescription": errorResponse.ErrorMessage}).Error(fmt.Sprintf("Error occured while trying to marshal request: %s", err))
		c.JSON(http.StatusBadRequest, errorResponse)
		return err
	}
	if len(request.Message) < 2 || request.Phone == "" {
		errorResponse.Errorcode = "03"
		errorResponse.ErrorMessage = "Invalid request, phonenumber must have a value"
		log.WithFields(fields).WithError(err).WithFields(log.Fields{"responseCode": errorResponse.Errorcode, "responseDescription": errorResponse.ErrorMessage}).Error("Invalid request, phonenumber must have a value")
		c.JSON(http.StatusBadRequest, errorResponse)
		return nil
	}
	go func() {
		accountSid := os.Getenv("TWILIO_SID")
		authToken := os.Getenv("TWILIO_AUTH_TOKEN")
		twilioSmsURL := os.Getenv("TWILIO_ENDPOINT")
		twilioNumber := os.Getenv("TWILIO_NUMBER")
		msgData := url.Values{}
		msgData.Set("To", request.Phone)
		msgData.Set("From", twilioNumber)
		msgData.Set("Body", request.Message)
		msgDataReader := *strings.NewReader(msgData.Encode())
		client := &http.Client{}
		req, _ := http.NewRequest("POST", twilioSmsURL, &msgDataReader)
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, _ := client.Do(req)
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var data map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			err := decoder.Decode(&data)
			if err == nil {
				log.WithFields(fields).Info(data["sid"])
				log.WithFields(fields).Info("Successfully sent sms")
			}
			log.WithFields(fields).WithError(err).WithFields(log.Fields{"httpResponseCode": resp.StatusCode, "httpResponseStatusMessage": resp.Status}).Error("SMS not sent. Twilio response body: ", data)
		} else {
			var data map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			err := decoder.Decode(&data)
			if err == nil {
				log.WithFields(fields).Error(data)
				log.WithFields(fields).Error("Error occured sms")
			}
			log.Println(resp.Status)
			log.WithFields(fields).WithError(err).WithFields(log.Fields{"httpResponseCode": resp.StatusCode, "httpResponseStatusMessage": resp.Status}).Error("Error occured while sending sms. Twilio response body: ", data)
		}
	}()
	successResponse := &models.SuccessResponse{
		ResponseCode:        "00",
		ResponseDescription: "SMS received for sending...",
		ResponseMessage:     nil,
	}
	log.WithFields(fields).Info("Successfully sent SMS to phone number ", request.Phone)
	c.JSON(http.StatusOK, successResponse)
	return nil
}
