package client

import (
	"fmt"
	"github.com/2336260845/hippo_search/config"
	"github.com/2336260845/hippo_search/gen-go/recall"
	"github.com/apache/thrift/lib/go/thrift"
)

func newRecallServiceClient(conf *config.Config) (client *recall.RecallServiceClient, err error) {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(conf.ServerConfig.RecallAddress)
	if err != nil {
		return nil, fmt.Errorf("newRecallServiceClient NewTSocket error, err=%+v", err.Error())
	}

	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return nil, fmt.Errorf("newRecallServiceClient GetTransport error, err=%+v", err.Error())
	}

	client = recall.NewRecallServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		return nil, fmt.Errorf("newRecallServiceClient NewQueryAnalysisServiceClientFactory error, err=%+v", err.Error())
	}

	return client, nil
}