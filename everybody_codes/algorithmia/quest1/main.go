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
	fmt.Println("Running part", part)
	if part == 1 {
	 	ans := part1(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else if part == 3 {
		ans := part3(input)
		utils.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1 (input string) int {
	enemies := strings.Split(input, "")
	tot_potions := 0

	for _, enemy := range enemies {
		if enemy == "B" {
			tot_potions += 1
		} else if enemy == "C" {
			tot_potions += 3
		}
	}
	return tot_potions
}

func part2 (input string) int {
	pairs := make([]string, 0, len(input)/2)
	enemies := strings.Split(input, "")
	tot_potions := 0

	for i := 0; i < len(input); i += 2 {
		pairs = append(pairs, input[i:i+2])
	}

	for i, enemy := range enemies {
		if enemy == "B" {
			tot_potions += 1
		} else if enemy == "C" {
			tot_potions += 3
		} else if enemy == "D" {
			tot_potions += 5
		}

		if pairs[i / 2][0] != 'x' && pairs[i / 2][1] != 'x' {
			tot_potions += 1
		}
	}
	
	return tot_potions
}



func part3 (input string) int {
	triplets := make([]string, 0, len(input)/3)
	enemies := strings.Split(input, "")
	tot_potions := 0

	for i := 0; i < len(input); i += 3 {
		triplets = append(triplets, input[i:i+3])
	}

	for i, enemy := range enemies {
		if enemy == "B" {
			tot_potions += 1
		} else if enemy == "C" {
			tot_potions += 3
		} else if enemy == "D" {
			tot_potions += 5
		}

		if triplets[i / 3][utils.mod(i-1, 3)] != 'x' && triplets[i / 3][utils.mod(i+1, 3)] != 'x' {
			tot_potions += 1
		}
	}
	
	return tot_potions
}