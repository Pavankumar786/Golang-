Example 1:
package main  
import (  
   "os"  
   "log"  
   "io/ioutil"  
   "fmt"  
)  
func main() {  
   file, err := os.Create("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   file.WriteString("Hi... there")  
   file.Close()  
   stream, err:= ioutil.ReadFile("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   readString := string(stream)  
   fmt.Println(readString)  
}  
Output:

Hi... there

Example 2:
package main  
import (  
   "os"  
   "log"  
   "io/ioutil"  
   "fmt"  
)  
func main() {  
   file, err := os.Create("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   file.WriteString("Hi... This is Pavankumar")  
   file.Close()  
   stream, err:= ioutil.ReadFile("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   readString := string(stream)  
   fmt.Println(readString)  
}  
Output:

Hi...This is pavankumar

Example 3:
package main  
import (  
   "os"  
   "log"  
   "io/ioutil"  
   "fmt"  
)  
func main() {  
   file, err := os.Create("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   file.WriteString("Hi... John")  
   file.Close()  
   stream, err:= ioutil.ReadFile("file.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   readString := string(stream)  
   fmt.Println(readString)  
}  
Output:

Hi...John