package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

func (head *Node) Append(value int) *Node {
    node := &Node{Value: value}
    if head == nil { return node }
    current := head
    for current.Next != nil { current = current.Next }
    current.Next = node
    return head
}

func main() {
    var head *Node
    head = head.Append(10); head = head.Append(20); head = head.Append(30)
    for node := head; node != nil; node = node.Next { fmt.Print(node.Value, " ") }
    fmt.Println()
}
