package godecimal

import (
	"reflect"
	"testing"
)

func TestCalculation(t *testing.T) {
	t.Log("加：", Float64Add(10, 3), reflect.TypeOf(Float64Add(10, 3)))
	t.Log("减", Float64Sub(10, 3), reflect.TypeOf(Float64Sub(10, 3)))
	t.Log("乘：", Float64Mul(10, 3), reflect.TypeOf(Float64Mul(10, 3)))
	t.Log("除：", Float64Quo(10, 3), reflect.TypeOf(Float64Quo(10, 3)))
}
