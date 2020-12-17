# Go-ReJSON - a golang client for ReJSON (a JSON data type for Redis)
Go-ReJSON is a [Go](https://golang.org/) client for [ReJSON](https://github.com/RedisLabsModules/rejson) Redis Module. 

[![GoDoc](https://godoc.org/github.com/kaiaulu/go-rejson?status.svg)](https://godoc.org/github.com/kaiaulu/go-rejson)
[![Build Status](https://travis-ci.org/kaiaulu/go-rejson.svg?branch=master)](https://travis-ci.org/kaiaulu/go-rejson)
[![codecov](https://coveralls.io/repos/github/kaiaulu/go-rejson/badge.svg?branch=master)](https://coveralls.io/github/kaiaulu/go-rejson?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaiaulu/go-rejson)](https://goreportcard.com/report/github.com/kaiaulu/go-rejson)

> ReJSON is a Redis module that implements ECMA-404 The JSON Data Interchange Standard as a native data type. It allows storing, updating and fetching JSON values from Redis keys (documents).

## Installation
	go get github.com/kaiaulu/go-rejson

## Example usage

```golang
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/kaiaulu/go-rejson"
)

// Name - student name
type Name struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

// Student - student object
type Student struct {
	Name Name `json:"name,omitempty"`
	Rank int  `json:"rank,omitempty"`
}

func Example_JSONSet(rh *rejson.Handler) {

	student := Student{
		Name: Name{
			"Mark",
			"S",
			"Pronto",
		},
		Rank: 1,
	}
	res, err := rh.JSONSet("student", ".", student)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return
	}

	if res.(string) == "OK" {
		fmt.Printf("Success: %s\n", res)
	} else {
		fmt.Println("Failed to Set: ")
	}

	studentJSON, err := rh.JSONGet("student", ".")
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}

	readStudent := Student{}
	err = json.Unmarshal(studentJSON., &readStudent)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
		return
	}

	fmt.Printf("Student read from redis : %#v\n", readStudent)
}

func main() {
	var addr = flag.String("Server", "localhost:6379", "Redis server address")

	rh := rejson.NewReJSONHandler()
	flag.Parse()

	cli := redis.NewClient(&redis.Options{Addr: *addr})
	defer func() {
		if err := cli.FlushAll(context.Background()).Err(); err != nil {
			log.Fatalf("goredis - failed to flush: %v", err)
		}
		if err := cli.Close(); err != nil {
			log.Fatalf("goredis - failed to communicate to redis-server: %v", err)
		}
	}()
	rh.SetRedisClient(cli)
	fmt.Println("\nExecuting Example_JSONSET for Redigo Client")
	Example_JSONSet(rh)
}
```
