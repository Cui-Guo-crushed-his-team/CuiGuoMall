package rpc

import (
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/conf"
	checkoututils "github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/order/utils"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/checkout/checkoutservice"
	"sync"

	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/rpc/kitex_gen/payment/paymentservice"

	"github.com/cloudwego/kitex/client"
)

var (
	CheckoutClient checkoutservice.Client
	PaymentClient  paymentservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	serviceName    string
	commonSuite    client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		})

		initPaymentClient()
		initCheckoutClient()
	})
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	checkoututils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	checkoututils.MustHandleError(err)
}
