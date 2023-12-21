package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-function-notification-email-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-function-notification-email-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"net/smtp"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) process(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
	queueMessage rabbitmq.RabbitmqMessage,
) {
	c.SendEmail(input, output, errs, log, queueMessage)
}

func (c *DPFMAPICaller) SendEmail(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
	queueMessage rabbitmq.RabbitmqMessage,
) {
	emailFrom := c.conf.Email.EmailFrom
	emailAuthPass := c.conf.Email.EmailAuthPass
	emailHost := c.conf.Email.EmailHost
	emailAddress := c.conf.Email.EmailAddress

	emailAuth := smtp.PlainAuth("", emailFrom, emailAuthPass, emailHost)

	convertedInputJsonString, err := json.Marshal(queueMessage.Data())
	if err != nil {
		log.Error(err)
	}

	for _, v := range *input.Message.Partner {
		if v.PartnerFunction == "DELIVERTO" {
			msg := []byte("" +
				"From: DeliveryTo <" + emailFrom + ">\r\n" +
				"To: " + v.EmailAddress + "\r\n" +
				"Subject: 出荷案内\r\n" +
				"\r\n" +
				"" + string(convertedInputJsonString) + "\r\n" +
				"")

			//fmt.Printf("%v %v %v %v ", emailAddress, emailAuth, emailFrom, msg)

			err = smtp.SendMail(emailAddress, emailAuth, emailFrom, []string{v.EmailAddress}, msg)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
