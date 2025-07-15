// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import ( 
    "fmt"
    "time"
    )


func main() {
  currentTime := time.Now()
  formattedTime := currentTime.Format("2006-01-02 03:04 PM Monday")
  fmt.Println(formattedTime)
  
  //The time.Parse function is used to parse a formatted string representing a time and convert it into a time
  
  dateStr := "2025-07-14"
  parsedTime, _ := time.Parse("2006-01-02 03:04 PM ",dateStr)
  fmt.Println("Parsed Time:", parsedTime) 
  
  //We can add or subtract time
  
  //adding 1 day
  newTime := currentTime.Add(24*time.Hour)
  fmt.Println("New time:",newTime)
  
}
