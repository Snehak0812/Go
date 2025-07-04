package main
import "fmt"

func main() {
    day := 2
    switch day{
        case 1: fmt.Println("Mon")
        case 2: fmt.Println("Tue")
        case 3: fmt.Println("Wed")
        default: fmt.Println("Invalid day")
    }
    month := "Mar"
    switch month{
        case "Jan","Feb","Mar": fmt.Println("1st Quarter")
        case "Apr","May","Jun": fmt.Println("2nd Quarter")
        default: fmt.Println("Invalid month")
    
    }
}
