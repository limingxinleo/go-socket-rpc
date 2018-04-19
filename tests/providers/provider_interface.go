package providers

type ProviderInterface interface {
	Register()
	GetInstance() interface{}
}
