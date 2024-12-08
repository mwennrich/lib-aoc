package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Atoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	CheckErr(err)
	return i
}

func ReadInput(path string) []string {
	readFile, err := os.Open(path)
	CheckErr(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

func GetIntSlice(line string) []int {
	var ints []int
	for _, s := range strings.Fields(line) {
		ints = append(ints, Atoi(s))
	}
	return ints
}

func GetArrayFromLines(lines []string) [][]string {
	var arr [][]string
	for _, line := range lines {
		arr = append(arr, strings.Split(line, ""))
	}
	return arr
}

func DeepCopyMap(m [][]string) [][]string {
	var copy [][]string
	for _, row := range m {
		var newRow []string
		newRow = append(newRow, row...)
		copy = append(copy, newRow)
	}
	return copy
}

func GetXY(p string) (int, int) {
	x, y := 0, 0
	fmt.Sscanf(p, "%d,%d", &x, &y)
	return x, y
}

func SetXY(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
