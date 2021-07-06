package main

// Oops, I started thinking of this problem with the first cell being X=1, Y=1, I forgot to start at zero. Sorry :)

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var w, h, c, r int
	flag.IntVar(&w, "w", 2, "The width")
	flag.IntVar(&h, "h", 2, "The height")
	flag.IntVar(&c, "c", 0, "The column")
	flag.IntVar(&r, "r", 0, "The row")
	flag.Parse()

	if c > w {
		fmt.Printf("Cannot get column %d when there are %d columns\n", c, w)
		os.Exit(1)
	}
	if r > h {
		fmt.Printf("Cannot get row %d when there are %d rows\n", c, w)
		os.Exit(1)
	}
	if w < 1 {
		fmt.Println("Cannot work on a table with less than 1 column")
		os.Exit(1)
	}
	if h < 1 {
		fmt.Println("Cannot work on a table with less than 1 row")
		os.Exit(1)
	}

	if c > 0 && r > 0 {
		fmt.Printf("value is %d\n", findValue(w, h, r, c))
	} else {
		printTable(w, h)
	}
}

func printTable(width, height int) {
	for r := 1; r <= height; r++ {
		for c := 1; c <= width; c++ {
			fmt.Printf("%3d ", findValue8(width, height, r, c))
		}
		fmt.Print("\n")
	}
}

// findValue returns the value for a specific cell
//
// How I found this algorithm:
//
// - First I tried with "if col == 1" (which was the easiest)
// - Then I added "else if col == 2" (then I realized it was related to the previous col)
// - Then I found this "row + col - 1" stuff that applies to all columns
// - Then I found out that if i >= width, I must add only width and not count upwards
// - Then I added "if col == 2 && row == height" and "else if col == 3 && row == height - 1" to substract 1
// - Then I found that it can be shortened and universalized as "if height-row-col == -2"
// - Then I added "if height-row-col == -3" to substract 3, and "height-row-col == -4" to substract 6, yay there is a pattern, let's find out!
// - And finally, this pattern results in the second loop, when height-row-col < 0
//
// I quickly abandoned filling a table (and maybe then extracting a value), because simply initializing the table with empty values takes twice the time, without even starting to calculate the values
func findValue(width, height, row, col int) int {
	result := col
	// Add 1 + 2 + 3 + ... + (row+col-1)-1, with some limits...
	for i := 1; i < row+col-1; i++ {
		if i < width {
			result += i
		} else {
			result += width
		}
	}
	// Substract values at the end...
	if height-row-col < 0 {
		for i := -1; i > height-row-col; i-- {
			result += i
		}
	}
	return result
}

// Here, I tried to integrate the second loop in the first one
//
// The benchmark show that it is slower for large tables, which is easily explained because the 2nd loop overhead is there only if height-row-col < 0 which is checked only once per value, whereas moving the test in the first loop executes that test row+col-1 times for each value
func findValue2(width, height, row, col int) int {
	result := col
	// Add 1 + 2 + 3 + ... + (row+col-1)-1, with some limits...
	for i := 1; i < row+col-1; i++ {
		if i < width {
			result += i
		} else {
			result += width
		}
		// Substract values at the end...
		if height-row-col < 0 && i < -(height-row-col) {
			result -= i
		}
	}
	return result
}

// A smarter move here, have the height-row-col < 0 test first, and then execute only one different loop according to the test value
func findValue3(width, height, row, col int) int {
	result := col
	if height-row-col < 0 {
		for i := 1; i < row+col-1; i++ {
			if i < width {
				result += i
			} else {
				result += width
			}
			if i < -(height - row - col) {
				result -= i
			}
		}
	} else {
		for i := 1; i < row+col-1; i++ {
			if i < width {
				result += i
			} else {
				result += width
			}
		}
	}
	return result

}

// But adding 1+2+...+n is equivalent to (n/2) * (n+1), so we can avoid a loop
// See https://www.youtube.com/watch?v=pqV_6IaVsxg
//
func sumOneToN(n int) int {
	return int((float32(n) / 2) * (float32(n) + 1))
}
func findValue4(width, height, row, col int) int {
	offset := row - 1
	colOffset := col + offset
	value := sumOneToN(colOffset) - offset
	if colOffset-height > 0 {
		value -= sumOneToN(colOffset - height)
	}
	if colOffset-width > 1 {
		value -= sumOneToN(colOffset - width - 1)
	}
	return value
}

// What if we remove the function call?
func findValue5(width, height, row, col int) int {
	offset := row - 1
	colOffset := col + offset
	value := int((float32(colOffset)/2)*(float32(colOffset)+1)) - offset
	if colOffset-height > 0 {
		value -= int((float32(colOffset-height) / 2) * (float32(colOffset-height) + 1))
	}
	if colOffset-width > 1 {
		value -= int((float32(colOffset-width-1) / 2) * (float32(colOffset-width-1) + 1))
	}
	return value
}

// What if we cache the previous calculated values?
// Let's try a 10 million entries cache...
var sumCache = [10000000]int{}

func cacheSumOneToN(n int) int {
	val := sumCache[n]
	if val == 0 {
		val = int((float32(n) / 2) * (float32(n) + 1))
		sumCache[n] = val
	}
	return val
}
func findValue6(width, height, row, col int) int {
	offset := row - 1
	colOffset := col + offset
	value := cacheSumOneToN(colOffset) - offset
	if colOffset-height > 0 {
		value -= cacheSumOneToN(colOffset - height)
	}
	if colOffset-width > 1 {
		value -= cacheSumOneToN(colOffset - width - 1)
	}
	return value
}

// And what if we cache the entries AND remove the function call?
var sumCache2 = [10000000]int{}

func findValue7(width, height, row, col int) int {
	offset := row - 1
	colOffset := col + offset
	value := sumCache2[colOffset]
	if value == 0 {
		value = int((float32(colOffset) / 2) * (float32(colOffset) + 1))
		sumCache2[colOffset] = value
	}
	value -= offset
	if colOffset-height > 0 {
		toSub := sumCache2[colOffset-height]
		if toSub == 0 {
			toSub = int((float32(colOffset-height) / 2) * (float32(colOffset-height) + 1))
			sumCache2[colOffset-height] = toSub
		}
		value -= toSub
	}
	if colOffset-width > 1 {
		toSub := sumCache2[colOffset-width-1]
		if toSub == 0 {
			toSub = int((float32(colOffset-width-1) / 2) * (float32(colOffset-width-1) + 1))
			sumCache2[colOffset-width-1] = toSub
		}
		value -= toSub
	}
	return value
}

// Let's take findValue5 and avoid converting to float...
// Or, talking the math vocabulary : (n/2) * (n+1) = (n * (n+1)) / 2
// (n × n+1) is always even, so there is no lost decimal when dividing
func findValue8(width, height, row, col int) int {
	offset := row - 1
	colOffset := col + offset
	value := (colOffset*(colOffset+1))/2 - offset
	if colOffset-height > 0 {
		value -= ((colOffset - height) * (colOffset - height + 1)) / 2
	}
	if colOffset-width > 1 {
		value -= ((colOffset - width - 1) * (colOffset - width)) / 2
	}
	return value
}
