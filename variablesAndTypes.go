package main
import "fmt"

func main() {
  //Variables with different data types
  var accNo int = 10908765
  var accBalance float64 = 54000
  var name string= "Sneha"
  //const - Values cannot be changed once assigned
  const branch string = "Bengaluru"
 
  //Print the values of variables
  fmt.Println("Hi", name,"!")
  fmt.Println("Your account number:",accNo)
  fmt.Println("Your account balance:",accBalance)
  fmt.Println("Your branch:",branch)

  //We can use the := notation to declare and initialize a variable in single line
  contactNo := 87890
  fmt.Println("Your contact no:",contactNo)
}
