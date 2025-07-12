package main 
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    )
var gradeBook map[string]float64
var count = 0
func addStudent() {
   fmt.Println("How many student details you want to add")
   reader := bufio.NewReader(os.Stdin)
   totalStud,err1 := reader.ReadString('\n')
   totalStud = strings.TrimSpace(totalStud)
   totalStudents,err2 := strconv.Atoi(totalStud)
   if err1 != nil || err2 != nil {
    fmt.Println("Invalid Input")
    fmt.Println("\n")
    return
   }
   
   for i:= 0; i < totalStudents; i++ {
        
        fmt.Printf("Enter name for student %d\n",i+1)
        name,err3:= reader.ReadString('\n')
        name = strings.TrimSpace(name)
        fmt.Printf("Enter marks for student %d\n",i+1)
        marks,err4 := reader.ReadString('\n')
        marks = strings.TrimSpace(marks)
        marksFloat,err5 := strconv.ParseFloat(marks,64)
        gradeBook[name] = marksFloat
        
        
        if err3 != nil || err4 != nil || err5 != nil {
            fmt.Println("Invalid Input")
            fmt.Println("\n")
            return
        }
   } 
   fmt.Printf("\n")
   fmt.Println("-------------------------------------")
   fmt.Println("Student details added successfully!!!")
   fmt.Println("-------------------------------------")
   fmt.Printf("\n")
}

func viewStudent() {
   
   fmt.Println("----------------GradeBook-----------------")
   fmt.Printf("| %-15s | %-10s |\n", "Name", "Marks")
   fmt.Println("------------------------------------------")
   for name,marks:= range gradeBook {
       count++
        fmt.Printf("| %-15s | %-10.2f | \n",name,marks)
   } 
   fmt.Println("------------------------------------------")
   fmt.Printf("\n")
   
}
func updateStudent() {
    fmt.Println("Enter name of student you want to modify")
    reader := bufio.NewReader(os.Stdin)
    name,err1:= reader.ReadString('\n')
    name = strings.TrimSpace(name)
    if err1 != nil {
        fmt.Println("Invalid Input")
        fmt.Println("\n")
        return
   }
   _,exist := gradeBook[name]
   if exist == false {
       fmt.Println("------------------------------------------")
       fmt.Println("Student doesn't exist!!")
       fmt.Println("------------------------------------------")
       fmt.Println("\n")
       return
   }    
   fmt.Printf("Enter new marks for %s\n",name)
        newMarks,_ := reader.ReadString('\n')
        newMarks = strings.TrimSpace(newMarks)
        marksFloat,_ := strconv.ParseFloat(newMarks,64)
        gradeBook[name] = marksFloat
        
    fmt.Printf("\n")
   fmt.Println("-------------------------------------")
   fmt.Println("Student details updated successfully!!!")
   fmt.Println("-------------------------------------")
   fmt.Printf("\n")
        
}
func deleteStudent() {
    fmt.Println("Enter name of student you want to delete")
    reader := bufio.NewReader(os.Stdin)
    name,err1:= reader.ReadString('\n')
    name = strings.TrimSpace(name)
    if err1 != nil {
        fmt.Println("Invalid Input")
        fmt.Println("\n")
        return
   }
   _,exist := gradeBook[name]
   if exist == false {
       fmt.Println("------------------------------------------")
       fmt.Println("Student doesn't exist!!")
       fmt.Println("------------------------------------------")
       fmt.Println("\n")
       return
   }  
   delete(gradeBook, name)
   fmt.Printf("\n")
   fmt.Println("-------------------------------------")
   fmt.Println("Student deleted successfully!!!")
   fmt.Println("-------------------------------------")
   fmt.Printf("\n")
    
}

func getAverage() {
    var sum float64
    var avg float64
    for _,marks := range gradeBook {
        sum += marks
        avg = sum/float64(count)
    }
    fmt.Println("-------------------")
    fmt.Println("Average marks:",avg)
    fmt.Println("-------------------")
}
func main() {
    gradeBook = make(map[string]float64) 
   for {
       fmt.Println("1.Add Student")
       fmt.Println("2.Update")
       fmt.Println("3.Delete")
       fmt.Println("4.View")
       fmt.Println("5.Get Average Marks")
       fmt.Println("6.Get Highest Marks")
       fmt.Println("7.Get Lowest Marks")
       
       reader := bufio.NewReader(os.Stdin)
       ch,_ := reader.ReadString('\n')
       ch = strings.TrimSpace(ch)
       choice,_ := strconv.Atoi(ch)
       
       switch choice {
           case 1: addStudent()
           case 2: updateStudent()
           case 3: deleteStudent()
           case 4: viewStudent()
           case 5: getAverage()
           //case 6: getMaxMarks()
           //case 7: getMinMarks()
           default: fmt.Println("Invalid choice")
       }
      
   }
    
   
  
}
