package main
import "fmt"
type Books struct {
title string
author string
subject string
book_id int
}
func main() {
   var Book1 Books   
   var Book2 Books    
   Book1.title = "Go Programming"
   Book1.author = "Pavan Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407

   Book2.title = "Telecom Billing"
   Book2.author = "Zara Ali"
   Book2.subject = "Telecom Billing Tutorial"
   Book2.book_id = 6495700
 
    fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}
Output:
Book 1 title      : Go Programming
Book 1 author     : Pavan Kumar
Book 1 subject    : Go Programming Tutorial
Book 1 book_id    : 6495407
Book 2 title      : Telecom Billing
Book 2 author     : Zara Ali
Book 2 subject    : Telecom Billing Tutorial
Book 2 book_id    : 6495700


Example 2:
package main
import "fmt"
type Books struct {
title string
author string
subject string
book_id int
}
func main() {
   var Book1 Books   
   var Book2 Books    
   Book1.title = "Go Programming"
   Book1.author = "Keerthi"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 259793

   Book2.title = "The Go Basics"
   Book2.author = "John Caarter"
   Book2.subject = "Study tonight.com"
   Book2.book_id = 45679513
 
    fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}


Output:
Book 1 title : Go Programming
Book 1 author : Keerthi
Book 1 subject : Go Programming Tutorial
Book 1 book_id : 259793
Book 2 title : The Go Basics
Book 2 author : John Caarter
Book 2 subject : Study tonight.com
Book 2 book_id : 45679513

Example 3:
package main 
import "fmt"
// Defining a struct type 
type Address struct { 
    Name    string 
    city    string 
    Pincode int
} 
func main() { 
var a Address  
fmt.Println(a) 
a1 := Address{"Pavan", "Rayachoti", 516269} 
fmt.Println("Address1: ", a1) 
a2 := Address{Name: "Anikaa", city: "Ballia", Pincode: 277001} 
fmt.Println("Address2: ", a2) 
a3 := Address{Name: "Delhi"} 
fmt.Println("Address3: ", a3) 
} 
Output:

{  0}
Address1:  {Pavan Rayachoti 516269}
Address2:  {Anikaa Ballia 277001}
Address3:  {Delhi  0}


Example 4:
package main 
import "fmt"
// Defining a struct type 
type Address struct { 
 Name    string 
city    string 
Pincode int

} 
func main() { 
var a Address  
fmt.Println(a) 
a1 := Address{"Keerthi", "Kadapa", 516269,} 
fmt.Println("Address1: ", a1) 
a2 := Address{Name: "Kumar", city: "Chitoor", Pincode: 277001} 
fmt.Println("Address2: ", a2) 
a3 := Address{Name: "AndhraPradesh"} 
fmt.Println("Address3: ", a3) 
} 

Output:
{  0}
Address1:  {Keerthi Kadapa 516269}
Address2:  {Kumar Chitoor 277001}
Address3:  {AndhraPradesh  0}

Example 5:
package main 
import "fmt"
  
// defining the struct 
type Car struct { 
Name, Model, Color string 
WeightInKg   float64 
} 
func main() { 
 c := Car{Name: "Ferrari", Model: "GTC4",Color: "Red", WeightInKg: 1920} 
 fmt.Println("Car Name: ", c.Name) 
 fmt.Println("Car Color: ", c.Color) 
 c.Color = "Black"
 fmt.Println("Car: ", c) 
} 
Output:
Car Name:  Ferrari
Car Color:  Red
Car:  {Ferrari GTC4 Black 1920}


Example 6:Pointers to structures
package main 
import "fmt"
// defining a structure 
type Employee struct { 
firstName, lastName string 
age, salary int
} 
  
func main() { 
emp8 := &Employee{"Sam", "Anderson", 55, 6000} 
 fmt.Println("First Name:", (*emp8).firstName) 
fmt.Println("Age:", (*emp8).age) 
} 
Output:

First Name: Sam
Age: 55

Example 7:
package main
import "fmt"
// defining a structure
type Employee struct {
firstName, lastName string
age, salary  int
}
func main() {
emp := &Employee{"Pavan", "Kumar", 24, 10000}
fmt.Println("First Name:", (*emp).firstName)
fmt.Println("Age:", (*emp).age)
fmt.Println("salary:", (*emp).salary)
}

Output:
First Name: Pavan
Age: 24
salary: 10000

Example 8:
package main
import "fmt"
// defining a structure
type Employee struct {
firstName, lastName string
age, salary , number  int
}
func main() {
emp := &Employee{"Pavan", "Kumar", 24, 10000,9703043048}
fmt.Println("First Name:", (*emp).firstName)
fmt.Println("Age:", (*emp).age)
fmt.Println("salary:", (*emp).salary)
fmt.Println("number:", (*emp).number)
}

Output:
First Name: Pavan
Age: 24
salary: 10000
number: 9703043048

Example 9:
package main
import "fmt"
// defining a structure
type Employee struct {
firstName, secondName,lastName string
age, salary , number  int
}
func main() {
emp := &Employee{"Pavan", "Kumar","Keerthi", 24, 10000,9703043048}
fmt.Println("First Name:", (*emp).firstName)
fmt.Println("Second Name:", (*emp).secondName)
fmt.Println("Last Name:", (*emp).lastName)
fmt.Println("Age:", (*emp).age)
fmt.Println("salary:", (*emp).salary)
fmt.Println("number:", (*emp).number)
}

Output:
First Name: Pavan
Second Name: Kumar
Last Name: Keerthi
Age: 24
salary: 10000
number: 9703043048

Example 10:
package main 
import "fmt"
  
// defining the struct 
type Car struct { 
Name, Model, Color string 
WeightInKg   float64 
} 
func main() { 
 c := Car{Name: "Audi", Model: "GTC4",Color: "Red", WeightInKg: 1500} 
 fmt.Println("Car Name: ", c.Name) 
 fmt.Println("Car Color: ", c.Color) 
 c.Color = "White"
 fmt.Println("Car: ", c) 
} 
Output:
Car Name:  Audi
Car Color:  Red
Car:  {Audi GTC4 Black 1500}