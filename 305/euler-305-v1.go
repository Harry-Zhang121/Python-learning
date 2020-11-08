package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/cheggaaa/pb/v3"
)

func intToArray(num int) []int {
	tmpArr := make([]int, 0)
	for ; num > 0; num /= 10 {
		tmpArr = append(tmpArr, num%10)
	}

	arr := make([]int, 0)
	for i := len(tmpArr) - 1; i >= 0; i-- {
		arr = append(arr, tmpArr[i])
	}

	return arr

}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func shiftWindow(window []int, next []int) []int {
	var windowLen int = len(window)

	for i := 0; i < windowLen-1; i++ {
		window[i] = window[i+1]
	}

	window[windowLen-1] = next[0]

	next = next[1:]
	return next
}

func f(num int, wg *sync.WaitGroup) int {
	// Start a progress bar
	bar := pb.StartNew(num)

	target := intToArray(num)
	var targetLen int = len(target)

	var counter int = 0
	var nextNum int = 1
	var position int

	// Init the window. fill it with numbers
	window := make([]int, targetLen, targetLen)
	position = (-1) * targetLen
	position++
	for ; position < 1; nextNum++ {
		for nextNumArr := intToArray(nextNum); len(nextNumArr) > 0; {
			nextNumArr = shiftWindow(window, nextNumArr)
			position++
		}
	}
	// Initialization complete

	// Start searching for num
	for ; ; nextNum++ {
		for nextNumArr := intToArray(nextNum); len(nextNumArr) > 0; {
			//  fmt.Println(window)
			if equalSlice(window, target) {
				counter++
				bar.Increment()
			}
			if counter == num {
				fmt.Printf("f(%v) = %v\n", num, position)
				bar.Finish()
				wg.Done()
				break
				// return position
			}

			nextNumArr = shiftWindow(window, nextNumArr)
			position++

		}
	}
}

func job(left int, right int) int {
	var sum int = 0
	var wg sync.WaitGroup

	for i := left; i <= right; i++ {
		wg.Add(1)
		go f(int(math.Pow(3, float64(i))), &wg)
		// fmt.Printf("Calulating f(%v)\n", int(math.Pow(3, float64(i))))
	}

	wg.Wait()
	return sum
}

// func main() {
// 	job(7, 8)
// }
