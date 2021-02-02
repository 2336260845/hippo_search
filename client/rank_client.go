package client

import (
	"fmt"
	"github.com/2336260845/hippo_search/config"
	"github.com/2336260845/hippo_search/gen-go/rank"
	"github.com/apache/thrift/lib/go/thrift"
)

func newRankServiceClient(conf *config.Config) (client *rank.RankServiceClient, err error) {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(conf.ServerConfig.RankAddress)
	if err != nil {
		return nil, fmt.Errorf("newRecallServiceClient NewTSocket error, err=%+v", err.Error())
	}

	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return nil, fmt.Errorf("newRecallServiceClient GetTransport error, err=%+v", err.Error())
	}

	client = rank.NewRankServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		return nil, fmt.Errorf("newRankServiceClient NewQueryAnalysisServiceClientFactory error, err=%+v", err.Error())
	}

	return client, nil
}