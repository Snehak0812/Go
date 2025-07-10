// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    )

func main() {
  
  fmt.Println("Enter two numbers")
  reader := bufio.NewReader(os.Stdin)
  operand1,_ := reader.ReadString('\n')
  operand1 = strings.TrimSpace(operand1)
  a,_ := strconv.Atoi(operand1)
  
  operand2,_ := reader.ReadString('\n')
  operand2 = strings.TrimSpace(operand2)
  b,_ := strconv.Atoi(operand2)
  
  fmt.Println("Enter operation")
  op,_ := reader.ReadString('\n')
  op = strings.TrimSpace(op)
  
  switch op{
      case "+" : fmt.Println("Sum:",a+b)
      case "-" : if a >= b {
                    fmt.Println("Diff:",a-b)
                 }else {
                     fmt.Println("Diff:",b-a)
                 }
      
      case "*" : fmt.Println("Prod:",a*b)
      case "/" : if b == 0 {
			        fmt.Println("Cannot divide by zero")
		         } else {
			     fmt.Println("Div:", a/b)
		         }
      default : fmt.Println("Invalid Input")
  }
  
}
