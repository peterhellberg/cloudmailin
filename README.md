# cloudmailin

Go package used to handle the CloudMailin [JSON Hash Email Message Format](http://docs.cloudmailin.com/http_post_formats/json/)

[![Build Status](https://travis-ci.org/peterhellberg/cloudmailin.svg?branch=master)](https://travis-ci.org/peterhellberg/cloudmailin)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/cloudmailin)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/cloudmailin#license-mit)

## Installation

```bash
go get -u github.com/peterhellberg/cloudmailin
```

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/peterhellberg/cloudmailin"
)

var port = getenv("PORT", "5454")

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("POST a message here.\n\n" +
			"curl -X POST -d @example.json http://0.0.0.0:" + port))

		return
	}

	if msg, err := cloudmailin.Decode(r.Body); err == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		json.NewEncoder(w).Encode(&map[string]string{
			"to":      msg.Headers.To,
			"subject": msg.Headers.Subject,
		})
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Listening on http://0.0.0.0:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
```

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
