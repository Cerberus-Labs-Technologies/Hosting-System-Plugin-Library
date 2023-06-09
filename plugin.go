package Hosting_System_Plugin_Library

type HostingPlugin struct {
	Name      string
	Listeners []HostingListener
	Version   string
	Author    string
}

type HostingPluginLoader struct {
	Plugins map[string]HostingPlugin
}

type HostingPluginInterface interface {
	GetName() string
	GetVersion() string
	GetAuthor() string
	RegisterListeners() []HostingListener
}
