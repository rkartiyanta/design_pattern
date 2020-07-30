package structural

import (
	"fmt"
	"strconv"
	"strings"
)

func Plus(a, b int) int {
	return a + b
}

func Min(a, b int) int {
	return a - b
}

func Calculate(question string) string {
	result := ""
	if strings.Contains(question, "+") || strings.Contains(question, "-") {
		iPlus := strings.Index(question, "+")
		iMin := strings.Index(question, "-")

		if iPlus < iMin && iPlus > 0 {
			firstNumber, _ := strconv.Atoi(question[0:iPlus])
			tempQuestion := question[iPlus+1:]
			secondNumber := 0
			if strings.Contains(tempQuestion, "+") || strings.Contains(tempQuestion, "-") {
				iiPlus := strings.Index(tempQuestion, "+")
				iiMin := strings.Index(tempQuestion, "-")

				if iiPlus < iiMin && iiPlus > 0 {
					secondNumber, _ = strconv.Atoi(tempQuestion[0:iiPlus])
					tempQuestion = tempQuestion[iiPlus:]
				} else {
					secondNumber, _ = strconv.Atoi(tempQuestion[0:iiMin])
					tempQuestion = tempQuestion[iiMin:]
				}
			} else {
				secondNumber, _ = strconv.Atoi(tempQuestion)
				tempQuestion = ""
			}

			tempResult := Plus(firstNumber, secondNumber)
			result = fmt.Sprintf("%d%s", tempResult, tempQuestion)
		} else {
			firstNumber, _ := strconv.Atoi(question[0:iMin]) // 2
			tempQuestion := question[iMin+1:]
			secondNumber := 0
			if strings.Contains(tempQuestion, "+") || strings.Contains(tempQuestion, "-") {
				iiPlus := strings.Index(tempQuestion, "+")
				iiMin := strings.Index(tempQuestion, "-")

				if iiPlus < iiMin && iiPlus > 0 {
					secondNumber, _ = strconv.Atoi(tempQuestion[0:iiPlus])
					tempQuestion = tempQuestion[iiPlus:]
				} else {
					secondNumber, _ = strconv.Atoi(tempQuestion[0:iiMin])
					tempQuestion = tempQuestion[iiMin:]
				}
			} else {
				secondNumber, _ = strconv.Atoi(tempQuestion)
				tempQuestion = ""
			}

			tempResult := Min(firstNumber, secondNumber)
			result = fmt.Sprintf("%d%s", tempResult, tempQuestion)
		}
	}

	if strings.Contains(result, "+") || strings.Contains(result, "-") {
		result = Calculate(result)
	}

	return result
}
