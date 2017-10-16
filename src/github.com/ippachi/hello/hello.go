package main

import "fmt"

func sum(a *int) {
  *a++
}

func main() {
  a := 0
  sum(&a)
  fmt.Println(a)
}
