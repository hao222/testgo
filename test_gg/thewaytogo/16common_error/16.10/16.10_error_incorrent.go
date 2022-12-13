package main

import (
	"errors"
	"io"
	"net/http"
)

// todo 大部分代码都是致力于错误的检测, 通常解决此问题的好办法是尽可能以闭包的形式封装你的错误检测

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	err := func() error {
		if req.Method != "GET" {
			return errors.New("expected GET")
		}
		if input := parseInput(req); input != "command" {
			return errors.New("malformed command")
		}
		// 可以在此进行其他的错误检测
	}()

	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, err)
		return
	}
	doSomething() ...
