package main

import "fmt"

func balanced(text string) bool {
    stack := []rune{}
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
    for _, ch := range text {
        switch ch {
        case '(', '[', '{':
            stack = append(stack, ch)
        case ')', ']', '}':
            if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] { return false }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

func main() {
    fmt.Println(balanced("{[()]}"), balanced("{[(])}"))
}
