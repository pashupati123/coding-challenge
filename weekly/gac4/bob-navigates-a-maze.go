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
 * Complete the 'minMoves' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY maze
 *  2. INTEGER x
 *  3. INTEGER y
 */

const INF = 999999999

type Pair struct {
	x, y int
}

type Graph struct {
	maze  [][]int32
	gold  []Pair
	dist  [101][101][11]int32
	dp    [11][2048]int32
	limit int
}

func min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func (g *Graph) bfs(index int) {
	m := len(g.maze)
	n := len(g.maze[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			g.dist[i][j][index] = INF
		}
	}
	queue := make([]Pair, 1)
	p := g.gold[index]
	queue[0] = p
	g.dist[p.x][p.y][index] = 0

	visited := make(map[Pair]bool)
	visited[p] = true

	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]

			for _, dir := range dirs {
				x := node.x + dir[0]
				y := node.y + dir[1]
				p := Pair{x, y}
				if visited[p] || x < 0 || y < 0 || x >= m || y >= n {
					continue
				}

				if g.maze[x][y] != 1 {
					visited[p] = true
					g.dist[x][y][index] = g.dist[node.x][node.y][index] + 1
					queue = append(queue, Pair{x, y})
				}
			}
		}
	}
}

func (g *Graph) bitMask(index int, mask int, x int32, y int32) int32 {
	if mask == g.limit {
		return g.dist[x][y][index]
	}

	if g.dp[index][mask] != -1 {
		return g.dp[index][mask]
	}

	res := int32(math.MaxInt32)

	for i := 0; i < len(g.gold); i++ {
		if (mask & (1 << uint32(i))) == 0 {
			newMask := mask | (1 << uint32(i))
			res = min(res, g.bitMask(i, newMask, x, y)+g.dist[g.gold[i].x][g.gold[i].y][index])
		}
	}
	g.dp[index][mask] = res
	return res
}

func minMoves(maze [][]int32, x int32, y int32) int32 {
	g := &Graph{
		maze: maze,
	}
	p := make([]Pair, 1)
	p[0] = Pair{0, 0}
	for i, row := range maze {
		for j, col := range row {
			if col == 2 {
				p = append(p, Pair{i, j})
			}
		}
	}

	g.gold = p
	l := len(p)
	g.limit = (1 << uint32(l)) - 1
	for i := 0; i < l; i++ {
		g.bfs(i)
	}

	for i := 0; i < 11; i++ {
		for j := 0; j < 2048; j++ {
			g.dp[i][j] = -1
		}
	}
	ans := g.bitMask(0, 1, x, y)
	if ans >= INF {
		return -1
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	mazeRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	mazeColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var maze [][]int32
	for i := 0; i < int(mazeRows); i++ {
		mazeRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var mazeRow []int32
		for _, mazeRowItem := range mazeRowTemp {
			mazeItemTemp, err := strconv.ParseInt(mazeRowItem, 10, 64)
			checkError(err)
			mazeItem := int32(mazeItemTemp)
			mazeRow = append(mazeRow, mazeItem)
		}

		if len(mazeRow) != int(mazeColumns) {
			panic("Bad input")
		}

		maze = append(maze, mazeRow)
	}

	xTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	x := int32(xTemp)

	yTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	y := int32(yTemp)

	result := minMoves(maze, x, y)

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
