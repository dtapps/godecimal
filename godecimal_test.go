package godecimal

import (
	"testing"
)

func TestDecimal(t *testing.T) {
	t.Log(Decimal(2.3333))
}

func TestRound(t *testing.T) {
	t.Log(Round(2.3333, 1))
	t.Log(Round(2.3333, 2))
	t.Log(Round(2.3333, 3))
	t.Log(Round(2.3333, 4))
}
