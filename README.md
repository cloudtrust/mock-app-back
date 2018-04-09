Mock App Back-end [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![GoDoc][godoc-img]][godoc] [![Go Report Card][report-img]][report] [![OpenTracing Badge][opentracing-img]][opentracing]
==================

## Install
Go installer can be found [here](https://golang.org/doc/install).
The following packages are needed:
```
go get github.com/modocache/gover
go get github.com/mattn/goveralls
go get golang.org/x/tools/cmd/cover
go get github.com/golang/dep/cmd/dep
```

## Build
```
cd cmd
go build
```

## Run
```
cd cmd
./cmd
```

## Run tests
```
cd pkg
go test
```

## Bibliography
* [Introduction to Go](https://talks.godoc.org/github.com/davecheney/introduction-to-go/introduction-to-go.slide)
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)
* [Go for Java Programmers](https://talks.golang.org/2015/go-for-java-programmers.slide)
* [What's in a name?](https://talks.golang.org/2014/names.slide)
* [Go @ Fratt's wiki](https://wiki.frattino.ch/doku.php?id=go)

[ci-img]: https://travis-ci.org/cloudtrust/mock-app-back.svg?branch=master
[ci]: https://travis-ci.org/cloudtrust/mock-app-back
[cov-img]: https://coveralls.io/repos/github/cloudtrust/mock-app-back/badge.svg?branch=master
[cov]: https://coveralls.io/github/cloudtrust/mock-app-back?branch=master
[godoc-img]: https://godoc.org/github.com/cloudtrust/mock-app-back?status.svg
[godoc]: https://godoc.org/github.com/cloudtrust/mock-app-back
[report-img]: https://goreportcard.com/badge/github.com/cloudtrust/mock-app-back
[report]: https://goreportcard.com/report/github.com/cloudtrust/mock-app-back
[opentracing-img]: https://img.shields.io/badge/OpenTracing-enabled-blue.svg
[opentracing]: http://opentracing.io