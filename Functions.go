package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)
func sum(a,b int) (result int) {
    result = a + b
    return   
}
func main() {
   
  fmt.Println("Enter two numbers")
  reader := bufio.NewReader(os.Stdin)
  
  num1,_ := reader.ReadString('\n')
  num1 = strings.TrimSpace(num1)
  num3,_ := strconv.Atoi(num1)
  
  num2,_ := reader.ReadString('\n')
  num2 = strings.TrimSpace(num2)
  num4,_ := strconv.Atoi(num2)
  
  sum2 := sum(num3, num4)
  fmt.Println("Sum:",sum2)
}
