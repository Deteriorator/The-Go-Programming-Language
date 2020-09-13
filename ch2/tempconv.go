/*
 * @program: The-Go-Programming-Language
 * @author: Leon
 * @create: 2020-09-13 22:49
 */
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }


func main(){
	fmt.Printf("brrrr! %v\n", tempconv)
}