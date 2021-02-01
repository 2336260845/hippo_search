module github.com/2336260845/hippo_search

go 1.15

require (
	github.com/apache/thrift/lib/go/thrift v0.0.0-20210120171102-e27e82c46ba4
	github.com/gin-gonic/gin v1.6.3
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/viper v1.7.1
)

replace github.com/apache/thrift/lib/go/thrift => ./go_thrift //gomod总是自动拉取最新代码，暂时未找到更好的办法，先替换为本地文件;本地版本为v0.13.0
