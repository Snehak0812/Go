package main
import "fmt"

func main() {
    //Declaration
    //var nums[5] int
    //nums[0] = 5
    //fmt.Println(nums[0])
    
    //Declaration and Initialization
    nums := [5]int{} //bydefault value is 0
    fmt.Println("Array:",nums)
     
    names := [5]string{"sk","sk","ic"}
    fmt.Println("Array:",names)
    fmt.Printf("Array: %q\n",names)
}
