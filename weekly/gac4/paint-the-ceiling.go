package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'variantsCount' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER s0
 *  3. INTEGER k
 *  4. INTEGER b
 *  5. INTEGER m
 *  6. LONG_INTEGER a
 */

func binarySearch(search int64, arr []int64) int64 {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2
		if arr[median] < search {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) {
		return int64(len(arr) - 1)
	} else if low == 0 {
		return 0
	} else if arr[low] == search {
		return int64(low)
	}

	return int64(low - 1)
}

func variantsCount(n int32, s0 int32, k int32, b int32, m int32, a int64) int64 {
	s := make([]int64, n)
	s[0] = int64(s0)

	for i := int32(1); i < n; i++ {
		s[i] = (int64(k)*s[i-1]+int64(b))%int64(m) + 1 + s[i-1]
	}

	var count int64
	z := int64(math.Sqrt(float64(a)))
	zi := binarySearch(z, s)
	if s[zi] > z {
		return 0
	}
	count = (zi + 1) * (zi + 1)
	for i := int32(0); i <= int32(zi); i++ {
		if s[i] > a {
			break
		}
		y := a / s[i]
		x := binarySearch(y, s)
		if s[x] > y {
			continue
		}

		count = count + 2*(x-zi)
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s0Temp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	s0 := int32(s0Temp)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	b := int32(bTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	a, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := variantsCount(n, s0, k, b, m, a)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
