package Hosting_System_Plugin_Library

type HostingPlugin struct {
	Name      string
	Listeners []HostingListener
	Version   string
	Author    string
	Load      func() error
	Debug     func() error
}

type HostingPluginLoader struct {
	Plugins map[string]HostingPlugin
}

type HostingPluginInterface interface {
	GetName() string
	GetVersion() string
	GetAuthor() string
	RegisterListener(listener HostingListener) []HostingListener
	RegisterListeners(listeners []HostingListener) []HostingListener
	GetListeners() []HostingListener
	GetListenerByName(name string) HostingListener
}
