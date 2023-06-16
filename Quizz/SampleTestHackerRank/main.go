package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(text string) string {
	uLang := strings.Split(text, " ")[1]
	allowedL := "C:CPP:JAVA:PYTHON:PERL:PHP:RUBY:CSHARP:HASKELL:CLOJURE:BASH:SCALA: ERLANG:CLISP:LUA:BRAINFUCK:JAVASCRIPT:GO:D:OCAML:R:PASCAL:SBCL:DART: GROOVY:OBJECTIVEC"
	allowedArr := strings.Split(allowedL, ":")
	for _, item := range allowedArr {
		if uLang == strings.Trim(item, " ") {
			return "VALID"
		}
	}
	return "INVALID"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		text := scanner.Text()
		fmt.Println(solution(text))
	}
}
