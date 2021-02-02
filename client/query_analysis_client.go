package client

import (
	"fmt"
	"github.com/2336260845/hippo_search/config"
	"github.com/2336260845/hippo_search/gen-go/query_analysis"
	"github.com/apache/thrift/lib/go/thrift"
)

//query分析 thrift client
func newQueryAnalysisServiceClient(conf *config.Config) (client *query_analysis.QueryAnalysisServiceClient, err error) {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(conf.ServerConfig.QueryAddress)
	if err != nil {
		return nil, fmt.Errorf("NewQueryAnalysisServiceClient NewTSocket error, err=%+v", err.Error())
	}

	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return nil, fmt.Errorf("NewQueryAnalysisServiceClient GetTransport error, err=%+v", err.Error())
	}

	client = query_analysis.NewQueryAnalysisServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		return nil, fmt.Errorf("NewQueryAnalysisServiceClient NewQueryAnalysisServiceClientFactory error, err=%+v", err.Error())
	}

	return client, nil
}
