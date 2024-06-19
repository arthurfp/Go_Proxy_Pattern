package proxy

import "testing"

func TestProxy_Request(t *testing.T) {
	realSubject := &RealSubject{}
	proxy := NewProxy(realSubject)
	expected := "Proxy: Controlling access to RealSubject\nRealSubject: Handling Request"
	result := proxy.Request()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
