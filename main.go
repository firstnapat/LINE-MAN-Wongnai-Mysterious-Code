package main

import (
	"encoding/base64"
  "fmt"
  "strings"
)

func Reverse(s string) string {
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}
func main() {
  sd, err := base64.StdEncoding.DecodeString("aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Decoded text: %s\n", sd)   
    fmt.Printf("Decoded text: %s\n", Reverse(string(sd)))   
    decodeText := strings.ReplaceAll(Reverse(string(sd)), ":", " ")
    fmt.Println("Decoded text:",decodeText)
}
  
