package main

import "regexp"

func main() {
	compile, err := regexp.Compile("[0-1]")
	if err != nil {
		return
	}

	matchString := compile.MatchString("中国12313jad我99")

	println(matchString)

}
