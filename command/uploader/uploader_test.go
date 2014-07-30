package uploader

import (
	"fmt"
	"testing"
)

func TestTelemFileToQueryString(t *testing.T) {
	actual := telemToQueryString("\na apple\n b banana\r\nc carrot")
	expected := "a=apple&b=banana&c=carrot"

	if actual != expected {
		t.Fatal(fmt.Sprintf("Expected \"%s\", received \"%s\"", expected, actual))
	}
}
