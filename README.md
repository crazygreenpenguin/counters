# counters
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/crazygreenpenguin/counters)
[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/crazygreenpenguin/graphite/master/LICENSE)
[![Build Status](https://travis-ci.org/crazygreenpenguin/counters.svg?branch=master)](https://travis-ci.org/crazygreenpenguin/counters)
[![Go Report Card](https://goreportcard.com/badge/github.com/crazygreenpenguin/counters)](https://goreportcard.com/report/github.com/crazygreenpenguin/counters)


Simple lightweight counters library for implementation  different application stats and counters primitives.
It's lock-free library, CAS instruction and atomic operation used instead this.

# Features
 - lightweight
 - thread safe
 - simple
 - lock-free

# Installation
 Use go-get to install
 ```
 go get github.com/crazygreenpenguin/counters
 ```
# External dependencies
This project has no external dependencies other,
 than the Go standard atomic library.

# Benchmarks result
```
   goos: darwin
   goarch: amd64
   pkg: counters
   BenchmarkMaxTime_Get-4         	1000000000	         0.281 ns/op
   BenchmarkMaxTime_Reset-4       	238555690	         5.02 ns/op
   BenchmarkMaxTime_Store-4       	134522019	         8.92 ns/op
   BenchmarkSequence_Set-4        	238822569	         5.02 ns/op
   BenchmarkSequence_Get-4        	1000000000	         0.557 ns/op
   BenchmarkSequence_Next-4       	134435308	         8.93 ns/op
   BenchmarkUint64Counter_Get-4   	1000000000	         0.279 ns/op
   BenchmarkUint64Counter_Set-4   	238969924	         5.02 ns/op
   BenchmarkUint64Counter_Inc-4   	238941595	         5.02 ns/op
   BenchmarkUint64Counter_Add-4   	238906304	         5.02 ns/op
```

#Use cases

 - Uint64Counter - it's simple counter. May used for request count for example.
 - Sequence - generate sequence uint64 uniq number. Guarantees uniq number for each Next().
 - MaxTime - using for find max request time in interval, for example.