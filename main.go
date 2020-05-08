package main

import (
	"fmt"
	"net/http"

	"github.com/xiaobudongzhang/micro-user-web/handler"

	"github.com/xiaobudongzhang/micro-basic/config"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	log "github.com/micro/go-micro/v2/util/log"

	"github.com/micro/go-micro/v2/web"

	basic "github.com/xiaobudongzhang/micro-basic/basic"
)

func main() {
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":9088"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// 注册登录接口
	service.HandleFunc("/user/login", handler.Login)
	// 注册退出接口
	service.HandleFunc("/user/logout", handler.Logout)
	service.HandleFunc("/user/test", handler.TestSession)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()

	fmt.Println(etcdCfg.GetHost(), etcdCfg.GetPort())
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
