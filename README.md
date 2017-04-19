# go-filelock

[![GoDoc](https://godoc.org/github.com/zbiljic/go-filelock?status.svg)](https://godoc.org/github.com/zbiljic/go-filelock)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/zbiljic/go-filelock/master/LICENSE)

Package go-filelock provides a cross-process mutex based on file locks that works on windows and *nix platforms.

## Installation

```bash
go get github.com/zbiljic/go-filelock
```

## Example:

```go
import github.com/zbiljic/go-filelock

fl, err := filelock.New(filename)
if err != nil {
    panic(err)
}
var lock filelock.TryLockerSafe
lock, err = fl.Lock()
if err != nil {
    panic(err)
}
defer lock.Unlock()

...
```

See the [reference][] for more info.

[reference]: http://godoc.org/github.com/zbiljic/go-filelock

---

Copyright © 2017 Nemanja Zbiljić
