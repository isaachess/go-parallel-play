package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	min_routines := 1
	max_routines := 1000
	for {
		text, _ := reader.ReadString('\n')
		text2 := strings.Replace(text, "\n", "", -1)
		num, _ := strconv.Atoi(text2)
		num_routines := num / 10000
		if num_routines < min_routines {
			num_routines = min_routines
		} else if num_routines > max_routines {
			num_routines = max_routines
		}
		fmt.Println("num_routines", num_routines)
		start := time.Now()
		sum := sum_me(num, num_routines)
		fmt.Println("time", time.Since(start))
		fmt.Println("sum", sum)
	}
}

func sum_me(stop int, num_routines int) int {
	diff := float64(stop) / float64(num_routines)
	floor_diff := math.Floor(diff)
	int_diff := int(floor_diff)
	ch := make(chan int)
	start := 0
	for i := 0; i < num_routines; i++ {
		new_stop := start + int_diff
		if new_stop > stop {
			new_stop = stop
		}
		go sum_parallel(start, new_stop, ch)
		start = start + int_diff + 1
	}

	final_sum := 0
	for i := 0; i < num_routines; i++ {
		final_sum += <-ch
	}
	return final_sum
}

func sum_parallel(start int, stop int, ch chan int) {
	sum := sum_numbers(start, stop)
	ch <- sum
}

func sum_numbers(start int, stop int) int {
	final := 0
	for i := start; i <= stop; i++ {
		final += i
	}
	return final
}
