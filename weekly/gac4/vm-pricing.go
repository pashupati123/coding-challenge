package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'interpolate' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY instances
 *  3. FLOAT_ARRAY price
 */

func interpolate(n int32, instances []int32, price []float32) string {
	m := make(map[int32]float32)
	newInstances := make([]int32, 0)
	for i, q := range instances {
		if price[i] > 0.0 {
			m[q] = price[i]
			newInstances = append(newInstances, q)
		}
	}

	var index int
	if m[n] != 0 {
		return fmt.Sprintf("%.2f", m[n])
	} else if len(newInstances) == 1 {
		return fmt.Sprintf("%.2f", m[newInstances[0]])
	} else if n < newInstances[0] {
		index = 0
	} else if n > newInstances[len(newInstances)-1] {
		index = len(newInstances) - 2
	} else {
		for i := 0; i < len(newInstances)-1; i++ {
			if n > newInstances[i] && n < newInstances[i+1] {
				index = i
				break
			}
		}
	}

	x1 := newInstances[index]
	x2 := newInstances[index+1]
	y1 := m[x1]
	y2 := m[x2]

	return fmt.Sprintf("%.2f", (y2-y1)*float32(n-x1)/float32(x2-x1)+y1)
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

	instancesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var instances []int32

	for i := 0; i < int(instancesCount); i++ {
		instancesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		instancesItem := int32(instancesItemTemp)
		instances = append(instances, instancesItem)
	}

	priceCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var price []float32

	for i := 0; i < int(priceCount); i++ {
		priceItemTemp, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
		checkError(err)
		priceItem := float32(priceItemTemp)
		price = append(price, priceItem)
	}

	result := interpolate(n, instances, price)

	fmt.Fprintf(writer, "%s\n", result)

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
