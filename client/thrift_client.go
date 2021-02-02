package client

import (
	"fmt"
	"github.com/2336260845/hippo_search/config"
	"github.com/2336260845/hippo_search/gen-go/query_analysis"
	"github.com/2336260845/hippo_search/gen-go/rank"
	"github.com/2336260845/hippo_search/gen-go/recall"
	"github.com/sirupsen/logrus"
	"time"
)

type ThriftClient struct {
	QueryAnalysisClient *query_analysis.QueryAnalysisServiceClient
	RecallClient        *recall.RecallServiceClient
	RankClient          *rank.RankServiceClient
}

var thriftClient *ThriftClient

func GetAllClient() *ThriftClient {
	return thriftClient
}

func ThriftInit(conf *config.Config) {
	thriftClient.update(conf)
	go thriftClient.thriftClientDo(conf)
}

func (tc *ThriftClient) thriftClientDo(conf *config.Config) {
	for {
		time.Sleep(time.Second * time.Duration(conf.TimeCircle.QueryCut))
		tc.update(conf)

		logrus.Infof("thriftClientDo update client")
	}
}

func (tc *ThriftClient) update(conf *config.Config) {
	qasc, err := newQueryAnalysisServiceClient(conf)
	if err != nil {
		logrus.Errorf(fmt.Sprintf("ThriftInit newQueryAnalysisServiceClient error, err=%+v", err.Error()))
	} else {
		logrus.Infof("ThriftInit newQueryAnalysisServiceClient success")
	}

	nrsc, err := newRecallServiceClient(conf)
	if err != nil {
		logrus.Errorf(fmt.Sprintf("ThriftInit newRecallServiceClient error, err=%+v", err.Error()))
	} else {
		logrus.Infof("ThriftInit newRecallServiceClient success")
	}

	nranksc, err := newRankServiceClient(conf)
	if err != nil {
		logrus.Errorf(fmt.Sprintf("ThriftInit newRankServiceClient error, err=%+v", err.Error()))
	} else {
		logrus.Infof("ThriftInit newRankServiceClient success")
	}

	thriftClient = &ThriftClient{}
	thriftClient.QueryAnalysisClient = qasc
	thriftClient.RecallClient = nrsc
	thriftClient.RankClient = nranksc
}
