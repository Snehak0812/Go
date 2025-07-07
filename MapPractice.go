package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    )

var studentData map[string]string 

func add() {
    fmt.Println("Enter Id and Name\n")
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("ID: ")
    id,_ := reader.ReadString('\n')
    id = strings.TrimSpace(id) 
    
    fmt.Print("Name: ")
    name,_ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
    
    studentData[id] = name
    
}

func modify() {
    
    fmt.Print("Enter ID you want to modify: ")
    reader := bufio.NewReader(os.Stdin)
    id,_ := reader.ReadString('\n')
    id = strings.TrimSpace(id) 
    
    fmt.Print("Enter new name: ")
    reader2 := bufio.NewReader(os.Stdin)
    name,_ := reader2.ReadString('\n')
    name = strings.TrimSpace(name)
    
    studentData[id] = name
    fmt.Println("Name updated!!")
}

func deleteData(id string) {
    
    delete(studentData,id)
    fmt.Println("Deleted successfully!!")
}
func exists(id string) bool{
    _,exists := studentData[id]
    return exists
}
func display() {
    for id, name := range studentData {
        fmt.Printf("Id: %s,Name: %s\n",id, name)
    }
}

func main() {
    studentData = make(map[string]string)
    for {
    fmt.Println("1 to Add")
    fmt.Println("2 to Modify")
    fmt.Println("3 to Delete")
    fmt.Println("4 to Check exists or not")
    fmt.Println("5 to Display")
    reader := bufio.NewReader(os.Stdin)
    
    choice,_ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)
    num, _:= strconv.Atoi(choice)
    
    switch num{
        case 1: fmt.Print("hi")
                add()
        case 2: modify()
        case 3: fmt.Print("Enter ID you want to delete")
                reader := bufio.NewReader(os.Stdin)
                id,_ := reader.ReadString('\n')
                id = strings.TrimSpace(id) 
                deleteData(id)
                
        case 4: fmt.Print("Enter ID you want to check")
                reader := bufio.NewReader(os.Stdin)
                id,_ := reader.ReadString('\n')
                id = strings.TrimSpace(id)
    
                if exists(id) {
                    fmt.Println("Id exists")
                } else {
                    fmt.Println("Id does not exist ")
                }
        case 5: display()
    }
    }
    
    
    
} 
