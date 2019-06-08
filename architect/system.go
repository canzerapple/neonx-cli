package architect

type System struct {
	gateways  []*HttpService
	services  []*RpcService
}
