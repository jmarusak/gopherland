package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type CelsiusFlag struct { Celsius }

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (v *CelsiusFlag) Set(s string) error {
	var value float64
	var unit string
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
		case "C", "°C":
			v.Celsius = Celsius(value)
			return nil
		case "F":
			v.Celsius = FToC(Fahrenheit(value))
			return nil
	}
	return fmt.Errorf("invalid temperature %g", s)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}
