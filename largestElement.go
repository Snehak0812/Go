package main

import "fmt"

func largestElement(nums []int) int{
    largest := nums[0] 
    for i := 0; i < len(nums); i++ {
        if nums[i] > largest {
            largest = nums[i]
        }
        continue
    }
    return largest
    
}
func main() {
    nums := []int{5,4,2,8,19}
    largest := largestElement(nums)
    fmt.Println(largest)
}
