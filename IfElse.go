package main
import (
    "fmt"
    "bufio"
    "os"
    )

func main() {
    fmt.Println("Enter two numbers")
    reader := bufio.NewReader(os.Stdin)
    a,_ := reader.ReadString('\n')
    b,_ := reader.ReadString('\n')
    if a > b {
        fmt.Println(a, "is greater than", b)
    }else if a == b {
        fmt.Println(a, "is equal to", b)
    }else {
        fmt.Println(a, "is smaller than", b)
    }
}
