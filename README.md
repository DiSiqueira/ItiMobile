# ItiMobile ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/ItiMobile) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Stable-brightgreen.svg)

ItiMobile is a matcher to be used in the [Itinerary][itinerary].

Match if the request is made by a Mobile Device. Powered by [user_agent] made by [mssola]

[itinerary]: https://github.com/DiSiqueira/Itinerary
[user_agent]: https://github.com/mssola/user_agent
[mssola]: https://github.com/mssola

## Project Status

ItiMobile is stable. Pull Requests [are welcome](https://github.com/DiSiqueira/ItiMobile#social-coding)

## Installation

### Go Get

```bash
$ go get github.com/disiqueira/itibot
```

## Basic usage

```go
r := itinerary.NewRouter()

r.NewPath(HeaderHandler).AddMatcher(itimobile.New(true)) //Match if the request is made by a Mobile Device
```

## Full Example

```go
package main

import (
	"github.com/disiqueira/itimobile"
	"github.com/disiqueira/itinerary"
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are on a Mobile device! :)"))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(HeaderHandler).AddMatcher(itimobile.New(true))

	http.ListenAndServe(":8000", r)
}
```

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/ItiMobile/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone https://github.com/DiSiqueira/ItiMobile.git itibot
$ cd itiwildcard/
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/ItiMobile/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.