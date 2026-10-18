package main

import "fmt"

type Notifier interface{ Notify(message string) }

type ConsoleNotifier struct{}
func (ConsoleNotifier) Notify(message string) { fmt.Println("Notification:", message) }

type Service struct{ notifier Notifier }
func (s Service) CompleteTask() { s.notifier.Notify("task completed") }

func main() {
    service := Service{notifier: ConsoleNotifier{}}
    service.CompleteTask()
}
