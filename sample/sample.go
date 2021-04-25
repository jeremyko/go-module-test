package main

import (
    "fmt"
    "github.com/jeremyko/my_mod" //테스트 해 볼 모듈을 import
)

func main() {
    message := my_mod.MyModTest()
    fmt.Println(message)
}

