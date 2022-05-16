package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	m := GetMessage()
	if !strings.Contains(m, "Hello") {
		t.Errorf("Hello not found in %s", m)
	}

}
