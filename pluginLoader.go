package Hosting_System_Plugin_Library

func (plugin *HostingPlugin) RegisterListener(listener HostingListener) error {
	plugin.Listeners = append(plugin.Listeners, listener)
	return nil
}

func (plugin *HostingPlugin) RegisterListeners(listeners []HostingListener) error {
	for _, listener := range listeners {
		err := plugin.RegisterListener(listener)
		if err != nil {
			return err
		}
	}
	return nil
}

func (plugin *HostingPlugin) GetListenerByName(name string) HostingListener {
	for _, listener := range plugin.Listeners {
		if listener.Name == name {
			return listener
		}
	}
	return HostingListener{}
}

func (plugin *HostingPlugin) GetName() string {
	return plugin.Name
}

func (plugin *HostingPlugin) GetVersion() string {
	return plugin.Version
}

func (plugin *HostingPlugin) GetAuthor() string {
	return plugin.Author
}

func (plugin *HostingPlugin) GetListeners() []HostingListener {
	return plugin.Listeners
}
