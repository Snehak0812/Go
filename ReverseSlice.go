package main

import "fmt" 

func main() {
    nums := []int{1,2,3,4,5}
    for i,j := 0,len(nums)-1; i < j; i,j = i+1,j-1 {
        nums[i],nums[j] = nums[j],nums[i]
    }
    fmt.Println(nums)
    
    str := []string{"book","a","is","this"}
    for a,b := 0,len(str)-1; a < b; a,b = a+1,b-1 {
        str[a],str[b] = str[b],str[a]
    } 
    fmt.Println(str)
}
