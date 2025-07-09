//Data Conversion

package main

import (
    "fmt"
    "strconv"
    )

func main() {
    var num int = 30
    fmt.Printf("%d\n",num)
    var floatNum float32 =float32(num)  //IntegerToFloat
    fmt.Printf("%.3f\n",floatNum)
    
    var num2 int = 40
    fmt.Printf("%d %T\n",num2,num2)
    str2 := strconv.Itoa(num2)   //IntegerToString
    fmt.Printf("%s %T\n",str2,str2)
    
    num3,_ := strconv.Atoi(str2) //StringToInteger
    fmt.Printf("%d %T\n",num3,num3)
    
    var str string = "3.14"
    num4, _ := strconv.ParseFloat(str,64) //StringToFloat
    fmt.Printf("%.3f %T",num4,num4)
}
