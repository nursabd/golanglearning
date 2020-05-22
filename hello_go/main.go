package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
} // ОДИН ИЗ СПОСОБОВ

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (c, d int) {
	c = sum * 4 / 9
	d = sum - c
	return

}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

/*type Vertex struct {
	X int
	Y int
}*/

// defer fmt.Println("world")
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	return

	names := [4]string{
		"Nursultan",
		"Zarina",
		"Aman",
		"Marat",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/*primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	return

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(sqrt(2), sqrt(-4))

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(sum1)*/

	/*
			var c, python, java bool

		var k, j int = 1, 2
			var za, nu int = 1, 2
			am := 3
			c, python, java := true, false, "no!"
			fmt.Println(za, nu, am, c, python, java)
			return
			fmt.Println(k, j, c, python, java)
			var i int
			fmt.Println(i, c, python, java)                            // без указание дают 0==false
			fmt.Println("The time is", time.Now())                     //time
			fmt.Println("My favorite number is", rand.Intn(50))        // random
			fmt.Printf("Now you have %g problems. \n", math.Sqrt(625)) // SQRT
			fmt.Println(add(25, 31))                                   // ДОБАВЛЯЕМ ЧИСЛА
			a, b := swap("hello", "nurs")                              // Метод свап
			fmt.Println(a, b)
			fmt.Println(split(17)) // метод разделить
	*/

}
