package wechat

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
)

type WorkConf struct {
	AgentID       int
	CorpID        string
	Secret        string
	HttpDebug     bool
	CacheAddrs    []string
	CachePassword string
	CacheDB       int
	Callback      string
	Scopes        []string
	File          string
	ErrorFile     string
	Stdout        bool
	Level         string
}

var Work *work.Work

func NewWork(c *WorkConf) *work.Work {
	app, err := work.NewWork(&work.UserConfig{
		CorpID:  c.CorpID,  // 企业微信的app id，所有企业微信共用一个。
		AgentID: c.AgentID, // 内部应用的app id
		Secret:  c.Secret,  // 内部应用的app secret
		OAuth: work.OAuth{
			Callback: c.Callback, //
			Scopes:   c.Scopes,
		},
		Log: work.Log{
			Level: c.Level,
			// 可以重定向到你的目录下，如果未设置File和Error，默认会在当前目录下的wechat文件夹下生成日志
			File:   c.File,
			Error:  c.ErrorFile,
			Stdout: c.Stdout, //  是否打印在终端
		},
		HttpDebug: c.HttpDebug,
		// 可选，不传默认走程序内存
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    c.CacheAddrs,
			Password: c.CachePassword,
			DB:       c.CacheDB,
		}),
	})
	if err != nil {
		panic(err)
	}
	return app
}
