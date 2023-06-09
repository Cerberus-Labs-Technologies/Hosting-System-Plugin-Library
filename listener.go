package Hosting_System_Plugin_Library

type HostingListener struct {
	Name    string
	Handler func() error
}
