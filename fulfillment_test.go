package azurempsaas

import "testing"

// Test function for HelloFulfillment()
func TestHelloFulfillment(t *testing.T) {
	v := HelloFulfillment()
	t.Logf("Returned: %s", v)

}
