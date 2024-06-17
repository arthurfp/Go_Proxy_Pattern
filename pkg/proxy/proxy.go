package proxy

// Proxy provides a surrogate or placeholder for the RealSubject to control access to it.
type Proxy struct {
	realSubject *RealSubject
}

// NewProxy creates a new Proxy with the given RealSubject.
func NewProxy(realSubject *RealSubject) *Proxy {
	return &Proxy{realSubject: realSubject}
}

// Request handles the request and controls access to the RealSubject.
func (p *Proxy) Request() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	return "Proxy: Controlling access to RealSubject\n" + p.realSubject.Request()
}
