module github.com/xiaobudongzhang/micro-user-web

go 1.14

replace github.com/xiaobudongzhang/micro-basic => /wwwroot/microdemo/micro-basic

replace github.com/xiaobudongzhang/micro-inventory-srv => /wwwroot/microdemo/micro-inventory-srv

replace github.com/xiaobudongzhang/micro-payment-srv => /wwwroot/microdemo/micro-payment-srv

replace github.com/xiaobudongzhang/micro-order-srv => /wwwroot/microdemo/micro-order-srv

replace github.com/xiaobudongzhang/micro-plugins => /wwwroot/microdemo/micro-plugins

replace github.com/xiaobudongzhang/micro-auth => /wwwroot/microdemo/micro-auth

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-log/log v0.2.0
	github.com/go-redis/redis v6.15.7+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/micro-in-cn/tutorials/microservice-in-micro v0.0.0-20200523172022-7bc5e0adb1bf // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.5.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/go-plugins/config/source/grpc/v2 v2.5.0 // indirect
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.5.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/xiaobudongzhang/micro-auth v1.1.1
	github.com/xiaobudongzhang/micro-basic v1.1.2
	github.com/xiaobudongzhang/micro-plugins v0.0.0-00010101000000-000000000000
	github.com/xiaobudongzhang/micro-user-srv v0.0.0-20200412120407-21b917b570be
)
