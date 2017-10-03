package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

var (
	i, j   int    = 1, 2
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	rand.Seed(10)
	fmt.Println("The unchanged random number is", rand.Intn(100))

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("The varying random number is", r1.Intn(100), "&", r1.Intn(100))

	fmt.Printf("Now you have %g problems.", math.Sqrt(7))

	fmt.Println("Value of pi is", math.Pi)

	fmt.Println("Add func is", add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println("Swap func is", a, b)

	fmt.Print("The naked return func ")
	fmt.Print(split(17))
	fmt.Println()

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum for for-loop is", sum)

	n := 1
	for n < 1000 {
		n += n
	}
	fmt.Println("This process is same as 2^n = ", n)

	fmt.Println("Handling complex numbers = ", sqrt(2), sqrt(-4))

	fmt.Print("SWITCH Loop: Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
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

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
