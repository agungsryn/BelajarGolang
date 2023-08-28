package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputProgram int
	fmt.Println("Pilih program dan tekan enter: ")
	fmt.Println("1. Program WeightedStrings")
	fmt.Println("2. Program Not Ready")
	fmt.Println("3. Pogram BalancedBracket ")
	fmt.Scan(&inputProgram)

	switch inputProgram {
	case 1:
		fmt.Println("\nProgram WeightedStrings dipilih.")
		fmt.Println(WeightedStrings())
	case 3:
		fmt.Println("\nProgram BalancedBracket dipilih.")
		fmt.Println(BalancedBracket())
	}

}

func WeightedStrings() (res []string) {

	var inputStr string
	var inputQueries string
	var queries []int
	fmt.Println("\nInput string :")
	fmt.Scan(&inputStr)
	fmt.Println("\nInput queries (pisahkan dengan dengan ,(koma) ) :")
	fmt.Scan(&inputQueries)

	// 1. init alphabet mapz
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetMap := make(map[string]int)
	for i, char := range alphabet {
		alphabetMap[string(char)] = i + 1
	}

	numStrings := strings.Split(inputQueries, ",")
	for _, numStr := range numStrings {
		numInt, _ := strconv.Atoi(numStr)
		queries = append(queries, numInt)
	}

	// 2. collect str value
	storedValueMap := make(map[string]int)
	for _, v := range inputStr {
		storedValueMap[string(v)] += alphabetMap[string(v)]
	}

	// 3. construct response
	i := 0
	for _, v := range inputStr {
		if total, ok := storedValueMap[string(v)]; ok {

			if total == queries[i] {
				res = append(res, "Yes")
			} else {
				res = append(res, "No")
			}

			delete(storedValueMap, string(v))
			i++
		}
	}

	return
}

type bracket string

const (
	bracketA bracket = "("
	bracketB bracket = "{"
	bracketC bracket = "["
)

func (b *bracket) getBracketCloser() string {
	switch *b {
	case bracketA:
		return ")"
	case bracketB:
		return "}"
	case bracketC:
		return "]"
	}
	return ""
}

func findBracket(str string) bracket {
	switch str {
	case "(":
		return bracketA
	case "{":
		return bracketB
	case "[":
		return bracketC
	}
	return ""
}
func BalancedBracket() string {
	completedIndexes := []bool{}
	var str string
	fmt.Println("\nInput string :")
	fmt.Scan(&str)
	str = "{ [ ()]}"
	str = strings.ReplaceAll(str, " ", "")
	for range str {
		completedIndexes = append(completedIndexes, false)
	}

	var currentBracket bracket = ""
	var totalSkip = 0
	var totalDepthEqualBracket = 0

	for i, v := range str {
		if !completedIndexes[i] {
			completedIndexes[i] = true
			//find current bracket
			currentBracket = findBracket(string(v))
			found := false
			for j, w := range str[i+1:] {
				switch string(w) {
				case string(currentBracket):
					//check if we have multiple open bracket
					totalDepthEqualBracket++
				case currentBracket.getBracketCloser():
					if totalDepthEqualBracket > 0 {
						totalDepthEqualBracket--
					} else {
						completedIndexes[i+j+1] = true
						if totalSkip%2 != 0 && totalSkip != 0 {
							return "NO"
						}
						found = true
						totalSkip = 0
					}
				}

				if found {
					break
				}
				totalSkip++
			}
		}
	}

	return "YES"
}