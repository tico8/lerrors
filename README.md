# lerrors
Package lerrors provides labeled error for handling.
labeled error is based on [Go1.13 errors](https://blog.golang.org/go1.13-errors).

`go get github.com/tico8/lerrors`


```go

import (
    "errors"
    le "github.com/tico8/lerrors"
)

var ErrNotFound = errors.New("not found")

func main() {
	book, err := Find("golang")
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("err: %+v", err) // error details
		// ....
	}

}

func Find(id string) (*Book, error) {
	b, err := otherpkg.Find(id)
	if err != nil && otherpkg.IsErrNotFound(err) {
		return nil, le.Wrap(err, ErrNotFound)
	}

	return b, nil
}
```
