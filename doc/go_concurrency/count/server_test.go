package count

import "testing"

func TestRunHelloWorldHandler(t *testing.T) {
	v := RunHelloWorldHandler(true)
	if v != "Hello, World!\n" {
		t.Fatalf("Error with: %v", v)
	}
}
