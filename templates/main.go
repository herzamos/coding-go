package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/herzamos/coding-go/utils"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part (1 or 2)")
	flag.Parse()
	// fmt.Println("Running part", part)
	if part == 1 {
	 	ans := part1(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}