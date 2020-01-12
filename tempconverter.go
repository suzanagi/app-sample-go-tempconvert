package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const AbsoluteZeroC Celsius = -273.15
const FreezingC Celsius = 0
const BoilingC Celsius = 100

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// CToF は摂氏の温度を華氏へ変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

// FToC は華氏の温度を摂氏へ変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

