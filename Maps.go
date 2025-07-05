package main
import "fmt"

func main() {
    studentMarks := make(map[string]int)
    studentMarks["Sneha"] = 90
    studentMarks["Sanjay"] = 80
    studentMarks["Ishaan"] = 85
    //fmt.Println("Sneha's marks:",studentMarks["Sneha"])
    
    //Addition
    studentMarks["Raunak"] = 84
    
    //Modification
    studentMarks["Sneha"] = 95
    
    //Deletion
    delete(studentMarks,"Ishaan")
    
    //Iterating
    for name, marks := range studentMarks {
        fmt.Printf("Name: %s, Marks: %d\n",name, marks)
    }
    
    //Checking if a value exists
    marks,exists := studentMarks["Sanjay"]
    fmt.Println("Sanjay marks exists: ",exists)
    fmt.Println("Sanjay marks: ",marks)
    
    //More concise way to create and initialize a map in single line
    
    students := map[string]int {
        "Alice" : 30,
        "Bob" : 40,
    }
    for name, marks := range students {
        fmt.Printf("Name: %s, Marks: %d\n",name, marks)
    }
}
