package Hosting_System_Plugin_Library

type Plugin struct {
	Name      string
	Listeners []Listener
	Version   string
	Author    string
}

type PluginLoader struct {
	Plugins map[string]Plugin
}

type PluginInterface interface {
	GetName() string
	GetVersion() string
	GetAuthor() string
	RegisterListeners() []Listener
}
