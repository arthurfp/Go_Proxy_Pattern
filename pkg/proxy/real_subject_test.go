package proxy

import "testing"

func TestRealSubject_Request(t *testing.T) {
	realSubject := &RealSubject{}
	expected := "RealSubject: Handling Request"
	result := realSubject.Request()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
