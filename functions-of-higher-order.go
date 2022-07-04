package main

import (
  "fmt"
  "strings"
)

func concatStringUsingFunctionOfHigerOrder(s string) func(string) func(string) string {
  s = fmt.Sprintf("Added by 1st fuction %s\n", strings.ToUpper(s))
  return func(t string) func(string) string {
    t = fmt.Sprintf("Added by 2nd fuction %s\n", strings.ToUpper(t))
    return func(u string) string {
      u = fmt.Sprintf("Added by 3rd fuction %s", strings.ToUpper(u))
      return strings.Title(s) + strings.Title(t) + strings.Title(u)
    }
  }
}

func main() {
	fmt.Println(concatStringUsingFunctionOfHigerOrder("functions of")("higher")("order"))
}
