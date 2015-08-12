package logic

import (
	"fmt"
	"strconv"
	"strings"
)

//Description:
// Number is a palindrome if it is equal to the number with digits in reversed order. For example, 5, 44, 171, 4884 are palindromes and 43, 194, 4773 are not palindromes.

// Write a method palindrome_chain_length which takes a positive number and returns the number of special steps needed to obtain a palindrome. The special step is: "reverse the digits, and add to the original number". If the resulting number is not a palindrome, repeat the procedure with the sum until the resulting number is a palindrome.

// If the input number is already a palindrome, the number of steps is 0.

// Input will always be a positive integer.

// For example, start with 87:

// 87 + 78 = 165; 165 + 561 = 726; 726 + 627 = 1353; 1353 + 3531 = 4884

// 4884 is a palindrome and we needed 4 steps to obtain it, so palindrome_chain_length(87) == 4

func PalindromeChainLength() {
	order := 2
	var inputsample = []int{
		87,
		88,
		89,
	}

	var resultsample = []int{
		4,
		0,
		24,
	}

	var chain_length int = 0
	input := inputsample[order]
	for {
		palindromeResult, resultNum := PalindromeProcess(input)
		input = resultNum
		if palindromeResult == true {
			if resultsample[order] == chain_length {
				fmt.Printf("well done! result : %d\n", chain_length)
				return
			} else {
				fmt.Printf("failed! result : %d\n", chain_length)
				return
			}
		}
		chain_length++
	}
}

func PalindromeProcess(input int) (bool, int) {
	fmt.Printf("**process**\n")

	str_n := strings.Split(strconv.Itoa(input), "")
	str_rev := getReverse(str_n)
	fmt.Printf("str_n : %v\n", str_n)
	fmt.Printf("str_rev : %v\n", str_rev)

	for cnt, n := range str_n {
		if n != str_rev[cnt] {
			nn := strings.Join(str_n, "")
			normalNum, err := strconv.Atoi(nn)
			if err != nil {
				fmt.Printf("normalNum strconv.Atoi error : %v\n", err)
			}

			rr := strings.Join(str_rev, "")
			reverseNum, err := strconv.Atoi(rr)
			if err != nil {
				fmt.Printf("reverseNum strconv.Atoi error : %v\n", err)
			}

			return false, (normalNum + reverseNum)
		}
	}
	return true, input
}

func getReverse(words []string) []string {
	resultwords := make([]string, len(words))
	for x, _ := range words {
		resultwords[x] = words[len(words)-x-1]
	}
	return resultwords
}
