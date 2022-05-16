package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	m := GetMessage()
	if !strings.Contains(m, "🗺️") {
		t.Errorf("🗺️ not found in %s", m)
	}
}
