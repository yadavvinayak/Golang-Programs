package main

import "fmt"

type Queue struct{ items []int }

func (q *Queue) Enqueue(v int) { q.items = append(q.items, v) }
func (q *Queue) Dequeue() (int, bool) {
    if len(q.items) == 0 { return 0, false }
    v := q.items[0]
    q.items = q.items[1:]
    return v, true
}

func main() {
    var q Queue
    q.Enqueue(10); q.Enqueue(20)
    v, ok := q.Dequeue()
    fmt.Println(v, ok)
}
