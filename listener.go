package Hosting_System_Plugin_Library

type HostingListener struct {
	Name    string
	Handler func() error
}

type HostingListenerInterface interface {
	GetName() string
	GetHandler() func() error
	Execute() error
}

func (hl *HostingListener) GetName() string {
	return hl.Name
}

func (hl *HostingListener) GetHandler() func() error {
	return hl.Handler
}

func (hl *HostingListener) Execute() error {
	return hl.Handler()
}
