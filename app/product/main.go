package main

import (
	"net"
	"time"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/product/biz/dal"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/product/conf"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		klog.Fatal("Error loading .env file")
	}

	// 初始化数据库连接
	dal.Init()

	// 获取配置并初始化 server 选项
	opts := kitexInit()

	// 创建 server 实例
	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)

	// 启动 server
	err = svr.Run()
	if err != nil {
		klog.Error("Server run error:", err)
	}
}

func kitexInit() (opts []server.Option) {
	// 获取配置
	cfg := conf.GetConf()

	// 解析地址
	addr, err := net.ResolveTCPAddr("tcp", cfg.Kitex.Address)
	if err != nil {
		klog.Fatal("Error resolving address:", err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	r, err := consul.NewConsulRegister(cfg.Registry.RegistryAddress[0])
	// 配置服务信息
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: cfg.Kitex.Service,
	}), server.WithRegistry(r))

	// 配置日志
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())

	// 配置异步日志写入
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Kitex.LogFileName,
			MaxSize:    cfg.Kitex.LogMaxSize,
			MaxBackups: cfg.Kitex.LogMaxBackups,
			MaxAge:     cfg.Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)

	// 注册 server 退出时的钩子，确保日志可以正常刷新
	server.RegisterShutdownHook(func() {
		err := asyncWriter.Sync()
		if err != nil {
			klog.Error("Error during log sync:", err)
		}
	})

	return
}
