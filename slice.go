package main
import "fmt"

func main() {
    nums := []int{1,2,3,4,5}
    nums = append(nums,6,7)
    fmt.Println("Slice:",nums,"length:",len(nums),"capacity:",cap(nums))
    
    //using make function to create integer slice with length 3 and capacity 5
    numbers := make([]int, 3, 5)
    fmt.Println("Slice:",numbers,"length:",len(numbers),"capacity:",cap(numbers))
}   
