package client

import (
	"fmt"
	"github.com/2336260845/hippo_search/config"
	"github.com/2336260845/hippo_search/gen-go/query_analysis"
	"github.com/2336260845/hippo_search/gen-go/recall"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/sirupsen/logrus"
	"time"
)

type ThriftClient struct {
	QueryAnalysisClient *query_analysis.QueryAnalysisServiceClient
	RecallClient        *recall.RecallServiceClient
}

var thriftClient *ThriftClient

func GetAllClient() *ThriftClient {
	return thriftClient
}

func ThriftInit(conf *config.Config) {
	thriftClient.update(conf)
	go thriftClient.thriftClientDo(conf)
}

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

func (tc *ThriftClient) thriftClientDo(conf *config.Config) {
	for {
		tc.update(conf)

		logrus.Infof("thriftClientDo update client")
		time.Sleep(time.Second * time.Duration(conf.TimeCircle.QueryCut))
	}
}

func (tc *ThriftClient) update(conf *config.Config) {
	qasc, err := newQueryAnalysisServiceClient(conf)
	if err != nil {
		logrus.Errorf(fmt.Sprintf("ThriftInit newQueryAnalysisServiceClient error, err=%+v", err.Error()))
	}

	nrsc, err := newRecallServiceClient(conf)
	if err != nil {
		logrus.Errorf(fmt.Sprintf("ThriftInit newRecallServiceClient error, err=%+v", err.Error()))
	}

	thriftClient = &ThriftClient{}
	thriftClient.QueryAnalysisClient = qasc
	thriftClient.RecallClient = nrsc
}
