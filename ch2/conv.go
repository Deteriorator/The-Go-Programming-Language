/*
 * @program: The-Go-Programming-Language
 * @author: Leon
 * @create: 2020-09-13 22:50
 */
package tempconv


func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }


