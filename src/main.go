package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"standardLibTest/src/rpcexe"
)

func fib(num int) {
	a, b := 0, 1
	for b < num {
		fmt.Print(b, ",")
		a, b = b, a+b
	}
}

func readFile() {
	file, e := os.Open("src/hello")
	if e != nil {
		fmt.Println(e)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))

}

func main() {
	rpcexe.GenerateService()
}
