//Strings package

package main 

import (
    "fmt"
    "strings"
    )

func main() {
    //Spltting Strings
    str := " apple,orange, banana, orange   "
    parts := strings.Split(str, ",")
    fmt.Println(parts)
    
    //counting occurrences
    count := strings.Count(str,"orange")
    fmt.Println("Count: ",count)
    
    //TrimSpace
    trimmed := strings.TrimSpace(str)
    fmt.Println("[",trimmed,"]")
    trimmedAll := strings.ReplaceAll(str, " ", "")
    fmt.Println("[",trimmedAll,"]")
    
    //Concatenation
    str1 := "Hello"
    str2 := "World"
    result := strings.Join([]string{str1,str2}," ")
    fmt.Println(result)
}
