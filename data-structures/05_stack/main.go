package main

import "fmt"

type Stack struct{ items []int }

func (s *Stack) Push(v int) { s.items = append(s.items, v) }
func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 { return 0, false }
    i := len(s.items)-1
    v := s.items[i]
    s.items = s.items[:i]
    return v, true
}

func main() {
    var s Stack
    s.Push(10); s.Push(20)
    v, ok := s.Pop()
    fmt.Println(v, ok)
}
