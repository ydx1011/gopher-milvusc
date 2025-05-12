package milvusc

import (
	"context"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"github.com/xfali/xlog"
	"github.com/ydx1011/gopher-core/bean"
	"github.com/ydx1011/yfig"
)

const (
	BuildinValueMilvuscSources = "gopher.milvusc"
)

type Sources struct {
	Address  string
	Username string
	Password string
}
type Processor struct {
	logger xlog.Logger
}

type Opt func(*Processor)

func NewProcessor(opts ...Opt) *Processor {
	ret := &Processor{
		logger: xlog.GetLogger(),
	}

	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func (p *Processor) Init(conf yfig.Properties, container bean.Container) error {
	dss := map[string]*Sources{}
	err := conf.GetValue(BuildinValueMilvuscSources, &dss)
	if len(dss) == 0 {
		p.logger.Errorln("No Database")
		return nil
	}
	for k, v := range dss {
		client, err := milvusclient.New(context.Background(), &milvusclient.ClientConfig{
			Address:  v.Address,
			Username: v.Username,
			Password: v.Password,
		})
		if err != nil {
			p.logger.Errorln(err)
			return nil
		}
		//添加到注入容器
		container.RegisterByName(k, client)
	}

	return err
}

func (p *Processor) Process() error {
	return nil
}

func (p *Processor) Classify(o interface{}) (bool, error) {
	//switch v := o.(type) {
	//case redis.Client:
	//	err := p.parseBean(v)
	//	return true, err
	//}
	return false, nil
}

func (p *Processor) BeanDestroy() error {
	return nil
}
