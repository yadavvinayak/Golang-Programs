package main

import (
    "fmt"
    "os"
)

func main() {
    path := "sample.txt"
    if err := os.WriteFile(path, []byte("Hello from Go file I/O!\n"), 0644); err != nil {
        fmt.Println("write error:", err)
        return
    }
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println("read error:", err)
        return
    }
    fmt.Print(string(content))
}
