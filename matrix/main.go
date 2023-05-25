package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	mat := [][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
		{1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}
	fmt.Printf("updateMatrix(mat): %v\n", updateMatrix(mat))
}

func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	for counter := range result {
		result[counter] = make([]int, len(mat[0]))
	}

	checker := map[string]int{}

	for row := range mat {
		for col := range mat[row] {
			if _, ok := checker[fmt.Sprint(col, row)]; !ok {
				findClosestPath(mat, col, row, checker)
			}
		}
	}

	for row := range mat {
		for col := range mat[row] {
			secondFindClosestPath(mat, col, row, checker)

		}
	}

	for key, value := range checker {
		address := strings.Split(key, " ")
		j, _ := strconv.Atoi(address[1])
		i, _ := strconv.Atoi(address[0])
		result[j][i] = value
	}
	return result
}

func findClosestPath(mat [][]int, i, j int, checker map[string]int) int {
	rows := len(mat)
	colmns := len(mat[0])

	if val, ok := checker[fmt.Sprint(i, j)]; ok {
		return val
	}

	checker[fmt.Sprint(i, j)] = 10000

	if isZero(mat, i, j) {
		checker[fmt.Sprint(i, j)] = 0
		return 0
	} else {
		var a, b, c, d int = 10000, 10000, 10000, 10000
		if check(i-1, j, rows, colmns) {
			a = findClosestPath(mat, i-1, j, checker)
		}
		if check(i+1, j, rows, colmns) {
			b = findClosestPath(mat, i+1, j, checker)
		}
		if check(i, j-1, rows, colmns) {
			c = findClosestPath(mat, i, j-1, checker)
		}
		if check(i, j+1, rows, colmns) {
			d = findClosestPath(mat, i, j+1, checker)
		}
		e := min(c, d)
		f := min(a, b)
		res := 0
		if f > e {
			checker[fmt.Sprint(i, j)] = e + 1
			res = e + 1
		} else {
			checker[fmt.Sprint(i, j)] = f + 1
			res = f + 1
		}

		return res
	}
}

func secondFindClosestPath(mat [][]int, i, j int, checker map[string]int) {
	rows := len(mat)
	colmns := len(mat[0])

	if isZero(mat, i, j) {

		return
	} else {

		var a, b, c, d int = 10000, 10000, 10000, 10000
		if check(i-1, j, rows, colmns) {
			a = checker[fmt.Sprint(i-1, j)]
		}
		if check(i+1, j, rows, colmns) {
			b = checker[fmt.Sprint(i+1, j)]
		}
		if check(i, j-1, rows, colmns) {
			c = checker[fmt.Sprint(i, j-1)]
		}
		if check(i, j+1, rows, colmns) {
			d = checker[fmt.Sprint(i, j+1)]
		}
		e := min(c, d)
		f := min(a, b)
		res := 0
		if f > e {
			checker[fmt.Sprint(i, j)] = e + 1
			res = e + 1
		} else {
			checker[fmt.Sprint(i, j)] = f + 1
			res = f + 1
		}
		checker[fmt.Sprint(i, j)] = res

	}
}

func check(i, j, rows, colmns int) bool {
	if i >= colmns || i <= -1 || j >= rows || j <= -1 {
		return false
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func isZero(mat [][]int, i, j int) bool {
	return mat[j][i] == 0
}
