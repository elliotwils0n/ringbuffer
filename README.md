# ringbuffer
Ring Buffer implementation based on Go's slices.

[![Go Docs][docs-badge]][docs-url]
[![MIT licensed][mit-badge]][mit-url]
[![CI Status][actions-badge]][actions-url]
[![Go Report Card][report-badge]][report-url]

[docs-badge]: https://godoc.org/github.com/elliotwils0n/ringbuffer?status.svg
[docs-url]: https://godoc.org/github.com/elliotwils0n/ringbuffer
[mit-badge]: https://img.shields.io/badge/license-MIT-blue.svg
[mit-url]: LICENSE
[actions-badge]: https://github.com/elliotwils0n/ringbuffer/workflows/CI/badge.svg
[actions-url]: https://github.com/elliotwils0n/ringbuffer/actions?query=workflow%3ACI+branch%3Amaster
[report-badge]: https://goreportcard.com/badge/github.com/elliotwils0n/ringbuffer
[report-url]: https://goreportcard.com/report/github.com/elliotwils0n/ringbuffer

## Usage

### Import module
```shell
go get -u github.com/elliotwils0n/ringbuffer
```

### Import package
```go
import (
    "github.com/elliotwils0n/ringbuffer"
)
```

### Init Ring Buffer
With or without initial capacity (defaults to 32)
```go
rb := ringbuffer.New[int]()
```
```go
rb := ringbuffer.NewWithCapacity[int](10);
```

### Push, pop and peek elements with Ring Buffer
Push back
```go
rb.PushBack(123)
```
Push front
```go
rb.PushFront(123)
```

Pop front, error returned on empty Ring Buffer
```go
element, err := rb.PopFront()
```
Pop back, error returned on empty Ring Buffer
```go
element, err := rb.PopBack()
```

Peek front/back, error returned on empty Ring Buffer
```go
head_element, err := rb.PeekFront()
tail_element, err := rb.PeekBack()
```
