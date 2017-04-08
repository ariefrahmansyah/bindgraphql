# bindgraphql
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ariefrahmansyah/bindgraphql)
[![CircleCI](https://circleci.com/gh/ariefrahmansyah/bindgraphql/tree/master.png?style=shield)](https://circleci.com/gh/ariefrahmansyah/bindgraphql/tree/master)
[![Coverage Status](https://coveralls.io/repos/github/ariefrahmansyah/bindgraphql/badge.svg?branch=master)](https://coveralls.io/github/ariefrahmansyah/bindgraphql?branch=master)
[![GoReportCard](https://goreportcard.com/badge/github.com/ariefrahmansyah/bindgraphql)](https://goreportcard.com/report/github.com/ariefrahmansyah/bindgraphql)

You have RESTful API. You have your struct, it use JSON tag. Then, one of your developer friend introduce you to GraphQL.

If you want to migrate your API to GraphQL without so much pain, this library is for you.

```go
import "github.com/ariefrahmansyah/bindgraphql"

graphObj, err := bind.NewObject("YourObj", yourObj)
```

# Example
See [example.go](https://github.com/ariefrahmansyah/bindgraphql/blob/master/example/example.go).

# Contributor
* [Ahmad Muzakki](https://gist.github.com/ahmadmuzakki29/73fc1e21bf7ae087a9ac53299032f09c)
