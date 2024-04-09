package main

import (
	"awesomeProject/karim"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	defer func() {
		f := time.Now()
		fmt.Println(t, f, f.Nanosecond()-t.Nanosecond(), f.Second()-t.Second())

	}()
	print("Руслан")
	a := 2423
	var b int = 5775
	fmt.Println(a*b + a - b/a + a)
	c := karim.Ruslan(2121212, 212121, 212121)
	fmt.Println(c)
	for i := 0; i < 60000000000; i++ {
		karim.Ruslan(84233.0, 12899, 93123.0)

	}
}
func print(david string) {
	fmt.Println(david)

}
func summa(david int, david2 int) int {
	return david + david2

}
