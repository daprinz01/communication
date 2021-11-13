package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"persianblack.com/communication/models"

	"github.com/labstack/echo/v4"
)

// SendSMS is used to send SMS messages
func SendSMS(c echo.Context) (err error) {
	log.Println("Send sms request received...")
	// var err error
	var errorResponse models.ErrorResponse
	request := new(models.SendSmsRequest)
	if err = c.Bind(request); err != nil {
		log.Println(fmt.Sprintf("Error occured while trying to marshal request: %s", err))
		errorResponse.Errorcode = "03"
		errorResponse.ErrorMessage = "Invalid request, phonenumber must have a value"
		c.JSON(http.StatusBadRequest, errorResponse)
		return err
	}
	// decoder := json.NewDecoder(r.Body)
	// err = decoder.Decode(&request)
	// if err != nil {
	// 	log.Println("Invalid request, phonenumber must have a value")
	// 	errorResponse.Errorcode = "03"
	// 	errorResponse.ErrorMessage = "Invalid request, phonenumber must have a value"

	// 	response, err := json.MarshalIndent(errorResponse, "", "")
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(response)
	// 	return
	// }
	// defer r.Body.Close()
	if len(request.Message) < 2 || request.Phone == "" {
		log.Println("Invalid request, phonenumber must have a value")
		errorResponse.Errorcode = "03"
		errorResponse.ErrorMessage = "Invalid request, phonenumber must have a value"

		// response, err := json.MarshalIndent(errorResponse, "", "")
		// if err != nil {
		// 	log.Println(err)
		// }
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(response)
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
				log.Println(data["sid"])
				log.Println("Successfully sent sms")
			}
		} else {
			var data map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			err := decoder.Decode(&data)
			if err == nil {
				log.Println(data)
				log.Println("Error occured sms")
			}
			log.Println(resp.Status)
		}
	}()
	successResponse := &models.SuccessResponse{
		ResponseCode:        "00",
		ResponseDescription: "SMS received for sending...",
		ResponseMessage:     nil,
	}
	// response, err := json.MarshalIndent(successResponse, "", "")
	// if err != nil {
	// 	log.Println(err)
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(response)
	c.JSON(http.StatusOK, successResponse)
	return nil
}
