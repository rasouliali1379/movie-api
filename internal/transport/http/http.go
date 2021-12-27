package http

type IRest interface {
	Start(address string) error
	Shutdown() error
}
