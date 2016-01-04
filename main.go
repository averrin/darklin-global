package main

import (
  "time"
  "fmt"
  "sync"
)

func main() {
  var wg sync.WaitGroup
  ticker := time.NewTicker(time.Second)
  wg.Add(1)
  fmt.Println("Ticker started")
  go func() {
    defer wg.Done()
    for t := range ticker.C {
      fmt.Println("Tick at", t)
    }
  }()
  wg.Wait()
  ticker.Stop()
  fmt.Println("Ticker stopped")
}
