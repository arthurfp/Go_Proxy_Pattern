package proxy

// RealSubject is the concrete implementation of Subject.
type RealSubject struct{}

// Request handles the actual request.
func (r *RealSubject) Request() string {
	return "RealSubject: Handling Request"
}
