package main
import "fmt"

func div(a, b float64) (float64,error){
    if b == 0 {
        return 0,fmt.Errorf("denominator must not be zero")
    }
    return a/b, nil
}

func main() {
    data,err := div(10,5)
    if err != nil {
        fmt.Print("Error:",err)
        return
    }
    
    fmt.Println("Result: ",data)
    fmt.Printf("Result is %.4f\n",data)
}
