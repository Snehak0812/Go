package main
import ( "fmt"
        "bufio"
        "os"
       )

func main () {
    //var name string
    //var age int
    fmt.Println("Enter name and age")
    //STANDARD INPUT(STDIN)
    
    //fmt.Scan(&name, &age)
    //I/P: sneha kumari 24
    //O/P: sneha 0
    
    //fmt.Scanf("%s %d",&name, &age)
    //I/P: sneha
    //O/P: sneha 0
    //I/P: sneha kumari 20
    //O/P: sneha 0

    //fmt.Scanln(&name, &age)
    //I/P: sneha kumari 23
    //O/P: sneha 0
    
    reader := bufio.NewReader(os.Stdin)
    name1,_ := reader.ReadString('\n')
    age1,_ := reader.ReadString('\n')
    fmt.Println(name1, age1)
}
