```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mutex sync.Mutex

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mutex.Lock()
                        count++
                        mutex.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Println(count) // Output should be 1000
}
```