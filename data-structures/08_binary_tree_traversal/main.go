package main

import "fmt"

type Node struct {
    Value int
    Left, Right *Node
}

func inorder(node *Node) {
    if node == nil { return }
    inorder(node.Left)
    fmt.Print(node.Value, " ")
    inorder(node.Right)
}

func main() {
    root := &Node{Value: 2, Left: &Node{Value: 1}, Right: &Node{Value: 3}}
    inorder(root)
    fmt.Println()
}
