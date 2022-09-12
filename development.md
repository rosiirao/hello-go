# Development

## API design and use

### Assert errors for behavior, not type

```go
func isTimeout(err error) bool {
  type timeout interface {
    Timeout() bool
  }
  te, ok = err.(timeout)
  return ok && te.Timeout()
}
```

## protocol buffers

### compile

```bash
$ protoc api/v1/​*​.proto ​\​
      --go_out=. ​\​
      --go_opt=paths=source_relative ​\​
      --proto_path=.
```

Code it in Makefile compile section.
