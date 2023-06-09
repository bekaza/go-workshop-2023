# First Step

create `wire.go`
```go
func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
```

change func main() to
```go
func main() {
	event := InitializeEvent()
	event.Start()
}
```

put on top of file `wire.go`
```
//+build wireinject
```

Then install wire
```
go install github.com/google/wire/cmd/wire@latest
```

Run command
```
wire
```
