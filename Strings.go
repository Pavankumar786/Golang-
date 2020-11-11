Example 1;
package main
import "fmt"
func main() {
var greeting =  "Hello world!"
fmt.Printf("normal string: ")
fmt.Printf("%s", greeting)
fmt.Printf("\n")
fmt.Printf("hex bytes: ")
for i := 0; i < len(greeting); i++ {
 fmt.Printf("%x ", greeting[i])
 }
fmt.Printf("\n")
const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
fmt.Printf("quoted string: ")
fmt.Printf("%+q", sampleText)
fmt.Printf("\n")  
}

Output:
normal string: Hello world!
hex bytes: 48 65 6c 6c 6f 20 77 6f 72 6c 64 21 
quoted string: "\xbd\xb2=\xbc \u2318"

Example 2:
package main
import "fmt"
func main() {
var greeting =  "Hello Pavankumar!"
fmt.Printf("normal string: ")
fmt.Printf("%s", greeting)
fmt.Printf("\n")
fmt.Printf("hex bytes: ")
for i := 0; i < len(greeting); i++ {
 fmt.Printf("%x ", greeting[i])
 }
fmt.Printf("\n")
const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
fmt.Printf("quoted string: ")
fmt.Printf("%+q", sampleText)
fmt.Printf("\n")  
}

Output:
normal string: Hello Pavankumar!
hex bytes: 48 65 6c 6c 6f 20 50 61 76 61 6e 6b 75 6d 61 72 21 
quoted string: "\xbd\xb2=\xbc \u2318"

Example 3:String Length
package main
import "fmt"
func main() {
var greeting =  "Hello world!"
fmt.Printf("String Length is: ")
fmt.Println(len(greeting))  
}

Output:
String Length is : 12


Example 4:
package main
import "fmt"
func main() {
var greeting =  "Hello world!"
fmt.Printf("String Length is: ")
fmt.Println(len(greeting))  
}

Output:
String Length is : 12

Example 5:Concatenating String
package main
import ("fmt" "math" )"fmt" "strings")
func main() {
greetings :=  []string{"Hello","world!"}   
fmt.Println(strings.Join(greetings, " "))
}

Output:
Hello world!


Example 6:
package main

import ("fmt" "math" )"fmt" "strings")

func main() {
   greetings :=  []string{"Hello","Pavan!"}   
   fmt.Println(strings.Join(greetings, " "))
}

Output:
Hello Pavan!


Example 7:
package main

import ("fmt" "math" )"fmt" "strings")

func main() {
   greetings :=  []string{"Hello","world!"}   
   fmt.Println(strings.Join(greetings, " "))
}

Output:
Hello World!