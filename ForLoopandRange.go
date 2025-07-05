package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    counter := 0
    for {
        fmt.Println("Infinite loop")
        counter++
        if counter == 3 {
            break
        }
    }
    //We use range keyword for looping over elements of a collection like slices, arrays, maps and strings
    
    nums := []int{1,2,3,4,5}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n",index,value)
    }
    
    msg := "Hello Sneha!"
    for index, char := range msg {
        fmt.Printf("Index: %d, Char:%c\n", index, char)
    }
    
}
