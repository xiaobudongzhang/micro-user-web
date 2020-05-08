module github.com/xiaobudongzhang/micro-user-web

go 1.14

replace github.com/xiaobudongzhang/micro-basic => /data/ndemo/micro-basic

replace github.com/xiaobudongzhang/micro-inventory-srv => /data/ndemo/micro-inventory-srv

replace github.com/xiaobudongzhang/micro-payment-srv => /data/ndemo/micro-payment-srv

replace github.com/xiaobudongzhang/micro-order-srv => /data/ndemo/micro-order-srv

replace github.com/xiaobudongzhang/micro-plugins => /data/ndemo/micro-plugins

replace github.com/xiaobudongzhang/micro-auth => /data/ndemo/micro-auth

require (
	github.com/go-log/log v0.2.0
	github.com/go-redis/redis v6.15.7+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/micro-in-cn/tutorials/microservice-in-micro v0.0.0-20200430044506-2451e30bf530
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.4.0
	github.com/xiaobudongzhang/micro-auth v1.1.1
	github.com/xiaobudongzhang/micro-basic v1.1.2
	github.com/xiaobudongzhang/micro-user-srv v0.0.0-20200412120407-21b917b570be
)
