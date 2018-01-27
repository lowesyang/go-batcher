# go-batcher
Simply create batch handler in go

## Install
`go get github.com/yyh1102/go-batcher`

## Usage
### Batch
```Go
batch:=NewBatch("TEST", 10, 5*time.Second, func(batch []interface{}) {
    // Do your task with batch
})
// Push single data into batch
batch.Push(1)

// Push batch of data into batch
arr:=[]interface{}{1,2,3}
batch.Batch(arr)
```

#### Batcher (manage batches)
```Go
......
batcher:=NewBatcher()
batcher.AddBatch(batch)
batcher.GetBatch("TEST")    // "Batch { name:"TEST", maxCapacity: 10, timeout:5s }"
```