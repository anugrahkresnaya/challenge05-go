package main

import (
	"fmt"
  "sync"
)

func randomCondition() {
  var wg sync.WaitGroup

  cobaInterface := []string{"coba1", "coba2", "coba3"}
  bisaInterface := []string{"bisa1", "bisa2", "bisa3"}
  for i := 0; i <= 3; i++ {
    wg.Add(2)
    go func(i int) {
      fmt.Println(cobaInterface, i + 1)
      wg.Done()
    }(i)
    go func(i int) {
      fmt.Println(bisaInterface, i + 1)
      wg.Done()
    }(i)
  }
  wg.Wait()
}

func structuredCondition() {
  var wg sync.WaitGroup
  var mu sync.Mutex

  cobaInterface := []string{"coba1", "coba2", "coba3"}
  bisaInterface := []string{"bisa1", "bisa2", "bisa3"}

  for i := 0; i <= 3; i++ {
    wg.Add(2)
    go func(i int) {
      fmt.Println(cobaInterface, i + 1)
    }

    func(i int) {
      
    }
  }
}

func main() {
  fmt.Println("random condition:")
  randomCondition()
}
