package utils

import "github.com/Knetic/govaluate"

func ParseBoolExpression(express string) (result bool) {
	expression, _ := govaluate.NewEvaluableExpression(express)
	res, _ := expression.Evaluate(nil)
	result, _ = res.(bool)
	return
}
