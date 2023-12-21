package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-function-notification-email-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-function-notification-email-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-function-notification-email-rmq-kube/config"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	db   *database.Mysql
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		rmq:  rmq,
		db:   db,
	}
}

func (c *DPFMAPICaller) AsyncFunction(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
	queueMessage rabbitmq.RabbitmqMessage,
) ([]byte, []error) {
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	c.process(nil, &mtx, input, output, accepter, &errs, log, queueMessage)

	return nil, nil
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	d, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := d.(string)
	if !ok {
		return false
	}
	return result == "success"
}

func getBoolPtr(b bool) *bool {
	return &b
}
