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

func get_num_routines(stop int) int {
	min_routines := 1
	max_routines := 1000
	num_routines := stop / 10000
	if num_routines < min_routines {
		return min_routines
	} else if num_routines > max_routines {
		return max_routines
	}
	return num_routines
}

func get_new_stop(start int, stop int, diff int) int {
	new_stop := start + diff
	if new_stop > stop {
		new_stop = stop
	}
	return new_stop
}

func sum_me(stop int, num_routines int) int {
	diff := float64(stop) / float64(num_routines)
	floor_diff := math.Floor(diff)
	int_diff := int(floor_diff)
	ch := make(chan int)
	start := 0
	for i := 0; i < num_routines; i++ {
		new_stop := get_new_stop(start, stop, int_diff)
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text2 := strings.Replace(text, "\n", "", -1)
		num, _ := strconv.Atoi(text2)
		num_routines := get_num_routines(num)
		fmt.Println("num_routines", num_routines)
		start := time.Now()
		sum := sum_me(num, num_routines)
		fmt.Println("time", time.Since(start))
		fmt.Println("sum", sum)
	}
}
