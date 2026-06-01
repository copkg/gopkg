package client

import (
	"context"
	"errors"
	"github.com/copkg/gopkg/config"
	"github.com/copkg/gopkg/logger"
	etcd_client "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"log"
	"sync"
)

var m sync.Map

func Get(serviceName string) client.XClient {
	share.Trace = true
	if cl, ok := m.Load(serviceName); ok && cl != nil {
		p := cl.(*client.XClientPool)
		return p.Get()
	}
	d, err := etcd_client.NewEtcdV3Discovery(config.Conf.GetString("etcd.path"), serviceName, config.Conf.GetStringSlice("etcd.hosts"), false, nil)
	if err != nil {
		log.Panicf("获取client pool err:%s", err.Error())
		return nil
	}
	pool := client.NewXClientPool(config.Conf.GetInt("service.pool_size"), serviceName, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	//defer pool.Close()
	m.Store(serviceName, pool)
	return pool.Get()
}
func Call(serviceName, serviceMethod string, req, res interface{}) error {
	cli := Get(serviceName)
	if cli == nil {
		log.Printf("创建rpc连接失败")
		return errors.New("创建rpc连接失败")
	}
	if err := cli.Call(context.Background(), serviceMethod, req, res); err != nil {
		log.Printf("call rpc service err:%s", err.Error())
		return err
	}
	return nil
}
func CallWithContext(ctx context.Context, serviceName, serviceMethod string, req, res interface{}) error {
	cli := Get(serviceName)
	if cli == nil {
		log.Printf("创建rpc连接失败")
		return errors.New("rpc服务连接失败")
	}
	if err := cli.Call(ctx, serviceMethod, req, res); err != nil {
		logger.Errorf("call rpc service err:%s-%s,err:", serviceName, serviceMethod, err.Error())
		return err
	}
	return nil
}
