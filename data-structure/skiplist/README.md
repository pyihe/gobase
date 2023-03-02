# go-skipList
skip list written in go

# usage example
```go
go get github.com/pyihe/go-skipList
```

```go
package main

import (
	"fmt"
	"github.com/pyihe/go-skipList"
)

func main() {
    ss := go_skipList.NewSkipList()
    ss.InsertByEle("k1", 10, nil)
    ss.InsertByEleArray("k2", 10.1, "this is k2", "k3", 1.1, nil)
    nodes, err := ss.GetElementByRank(1)
    if err != nil {
        //handle err
    }
    //output: mem: k2, score: 10.1, data: this is k2
    fmt.Printf("mem: %s, score: %v, data: %v\n", nodes[0].Name(), nodes[0].Score(), nodes[0].Data())
}
```