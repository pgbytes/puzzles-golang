package main

import (
	"fmt"
	"log"
	"puzzles-golang/ds"
)

const sampleBalanced = "{[()()][]}"
const sampleBalancedImpure = "{[(3)(a+b)][]}"
const sampleUnbalanced = "{[()}"

func main() {
	fmt.Printf("%s -> %t\n", sampleBalanced, isBalanced(sampleBalanced))
	fmt.Printf("%s -> %t\n", sampleBalancedImpure, isBalanced(sampleBalancedImpure))
	fmt.Printf("%s -> %t\n", sampleUnbalanced, isBalanced(sampleUnbalanced))
}

func isBalanced(sample string) bool {
	braceStack := ds.NewStack()
	sampleRune := []rune(sample)
	for i := 0; i < len(sampleRune); i++ {
		brace := string(sampleRune[i])
		// Check for opening braces
		if brace == "{" || brace == "[" || brace == "(" {
			braceStack.Push(string(brace))
		}
		// Check for closing braces
		if brace == "}" || brace == "]" || brace == ")" {
			opening, err := braceStack.Pop()
			// Check if stack is empty.. hence no opening braces
			if err != nil {
				log.Fatalln(err)
				return false
			}
			if (opening == "{" && brace != "}") || (opening == "[" && brace != "]") || (opening == "(" && brace != ")") {
				return false
			}
		}
	}
	return braceStack.IsEmpty()
}
