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
	fmt.Println("2. Program HighestPalindrome")
	fmt.Println("3. Pogram BalancedBracket ")
	fmt.Scan(&inputProgram)

	switch inputProgram {
	case 1:
		fmt.Println("\nProgram WeightedStrings dipilih.")
		fmt.Println(WeightedStrings())
	case 2:
		fmt.Println("\nProgram HighestPalindrome dipilih.")
		fmt.Println(HighestPalindrome())
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

// Nomor 3 

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

// Nomor 2

func GetHigher(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func UpdateRes(res, newStr []string) string {
	length := len(res)
	r := []string{}
	
	r = append(r, res[:length/2]...)
	r = append(r, newStr...)
	r = append(r, res[length/2:]...)
	return strings.Join(r, "")
}

func isPalindrome(res *string, str []string, changeRemaining *int) bool {
	length := len(str)

	if length == 0 {
		return true
	}

	if length == 1 {
		return false
	} 

	if str[0] != str[length-1]  {
		if *changeRemaining == 0 {
			return false
		}

		*changeRemaining -= 1

		val1, _ := strconv.Atoi(string(str[0]))
		val2, _ := strconv.Atoi(string(str[length-1]))

		higher := GetHigher(val1, val2)

		str[0] = fmt.Sprint(higher)
		str[length-1] = fmt.Sprint(higher)

	}
	
	*res = UpdateRes(strings.Split(*res, ""), []string{str[0], str[length-1]})

	return isPalindrome(res, str[1:length-1], changeRemaining)
}

func HighestPalindrome() (res string) {
	var str string
	var k int
	fmt.Println("\nInput string :")
	fmt.Scan(&str)
	fmt.Println("\nInput changeRemaining(int) :")
	fmt.Scan(&k)
	val := isPalindrome(&res, strings.Split(str, ""), &k)
	if (!val) {
		res = "-1"
	}
	return
}