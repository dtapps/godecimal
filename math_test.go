package godecimal

import "testing"

func TestMath(t *testing.T) {
	t.Log(Abs(111))
	t.Log(Abs(-111))

	t.Log(Floor(3.4))
	t.Log(Floor(3.5))

	t.Log(Ceil(3.4))
	t.Log(Ceil(3.5))

	t.Log(Round(3.4))
	t.Log(Round(3.5))

	t.Log(RoundPoint(3.4))
	t.Log(RoundPoint(3.5))
	t.Log(RoundPoint(3.14))
	t.Log(RoundPoint(3.14159))

	t.Log(Max(4, 3))

	t.Log(Min(4, 3))
}
