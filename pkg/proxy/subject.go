package proxy

// Subject defines the interface for RealSubject and Proxy.
type Subject interface {
	Request() string
}
