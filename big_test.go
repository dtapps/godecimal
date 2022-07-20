package godecimal

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	priceString := NewString("100")
	t.Log("从字符串创建：", priceString)

	priceFloat := NewFloat(0.20)
	t.Log("从浮点数创建：", priceFloat)

	priceInt := NewInt(300)
	t.Log("从整数创建：", priceInt)

	priceUint := NewUint(400)
	t.Log("从无符合整数创建：", priceUint)

	t.Log("加：", priceString.Add(priceFloat), reflect.TypeOf(priceString))
	t.Log("减", priceString.Sub(priceFloat), reflect.TypeOf(priceString))
	t.Log("乘：", priceString.Mul(priceFloat), reflect.TypeOf(priceString))
	t.Log("除：", priceString.Quo(priceFloat), reflect.TypeOf(priceString))
}

// https://www.itranslater.com/qa/details/2582643503239005184
// https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision
func TestOperation(t *testing.T) {

	a1 := NewFloat(0.1890)
	t.Log("a1：", a1)

	a2 := NewFloat(0.50)
	t.Log("a2：", a2)

	t.Log("a1+a2：", a1.Add(a2), a1.Add(a2).String(), a1.Add(a2).Float64(), reflect.TypeOf(a1.Add(a2).Float64()), a1.Add(a2).MoneyFloat64(), reflect.TypeOf(a1.Add(a2).MoneyFloat64()))
	t.Log("a1-a2：", a1.Sub(a2), a1.Sub(a2).String(), a1.Sub(a2).Float64(), reflect.TypeOf(a1.Sub(a2).Float64()), a1.Sub(a2).MoneyFloat64(), reflect.TypeOf(a1.Sub(a2).MoneyFloat64()))
	t.Log("a1*a2：", a1.Mul(a2), a1.Mul(a2).String(), a1.Mul(a2).Float64(), reflect.TypeOf(a1.Mul(a2).Float64()), a1.Mul(a2).MoneyFloat64(), reflect.TypeOf(a1.Mul(a2).MoneyFloat64()))
	t.Log("a1/a2：", a1.Quo(a2), a1.Quo(a2).String(), a1.Quo(a2).Float64(), reflect.TypeOf(a1.Quo(a2).Float64()), a1.Quo(a2).MoneyFloat64(), reflect.TypeOf(a1.Quo(a2).MoneyFloat64()))

	a3 := NewFloat(.035)
	t.Log("a3：", a3)

	a4 := NewFloat(.08875)
	t.Log("a4：", a4)

	t.Log("a3+a4：", a3.Add(a4), a3.Add(a4).String(), a3.Add(a4).Float64(), reflect.TypeOf(a3.Add(a4).Float64()), a3.Add(a4).MoneyFloat64(), reflect.TypeOf(a3.Add(a4).MoneyFloat64()))
	t.Log("a3-a4：", a3.Sub(a4), a3.Sub(a4).String(), a3.Sub(a4).Float64(), reflect.TypeOf(a3.Sub(a4).Float64()), a3.Sub(a4).MoneyFloat64(), reflect.TypeOf(a3.Sub(a4).MoneyFloat64()))
	t.Log("a3*a4：", a3.Mul(a4), a3.Mul(a4).String(), a3.Mul(a4).Float64(), reflect.TypeOf(a3.Mul(a4).Float64()), a3.Mul(a4).MoneyFloat64(), reflect.TypeOf(a3.Mul(a4).MoneyFloat64()))
	t.Log("a3/a4：", a3.Quo(a4), a3.Quo(a4).String(), a3.Quo(a4).Float64(), reflect.TypeOf(a3.Quo(a4).Float64()), a3.Quo(a4).MoneyFloat64(), reflect.TypeOf(a3.Quo(a4).MoneyFloat64()))
}

func TestDecimal1(t *testing.T) {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal, reflect.TypeOf(subtotal)) // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax, reflect.TypeOf(preTax))      // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                     // Taxes: 37.482861375
	fmt.Println("Total:", total)                                 // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax))      // Tax rate: 0.08875

	t.Log(total.String())
	t.Log(total.Float64())
	t.Log(total.Floor())

}

func Test3(t *testing.T) {

	var x1, y1 float64 = 10, 3
	z1 := x1 / y1
	t.Log(x1, y1, z1)

	x2, y2 := big.NewFloat(10), big.NewFloat(3)
	z2 := new(big.Float).Quo(x2, y2)
	t.Log(x2, y2, z2)

}
