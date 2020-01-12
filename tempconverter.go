package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const AbsoluteZeroC Celsius = -273.15
const FreezingC Celsius = 0
const BoilingC Celsius = 100

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

// CToF は摂氏の温度を華氏へ変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

// FToC は華氏の温度を摂氏へ変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC は絶対温度を摂氏へ変換します。
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

// CToK は摂氏を絶対温度へ変換します。
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }
