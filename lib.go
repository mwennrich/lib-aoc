package aoc

import (
	"bufio"
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
