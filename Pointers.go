Example 1: Pointers
package main
import "fmt"
func main() {
var a int = 30  
fmt.Printf("Address of a variable: %x\n", &a  )
}

Output:
Address of a variable: c42000e1f8

Example 2:
package main
import "fmt"
func main() {
var a int = 40   
var ip *int      
ip = &a  
fmt.Printf("Address of a variable: %x\n", &a  )
fmt.Printf("Address stored in ip variable: %x\n", ip )
fmt.Printf("Value of *ip variable: %d\n", *ip )
}

Output:
Address of a variable: c42000e1f8
Address stored in ip variable: c42000e1f8
Value of *ip variable: 40

Example 3:
package main
import "fmt"
func main() {
var  ptr *int
fmt.Printf("The value of ptr is : %x\n", ptr  )
}

Output:
The value of ptr is : 0

Example 4:
package main 
import "fmt"
func main() { 
x := 0xFF 
y := 0x9C 
fmt.Printf("Type of variable x is %T\n", x) 
fmt.Printf("Value of x in hexdecimal is %X\n", x) 
fmt.Printf("Value of x in decimal is %v\n", x) 
fmt.Printf("Type of variable y is %T\n", y) 
fmt.Printf("Value of y in hexdecimal is %X\n", y) 
fmt.Printf("Value of y in decimal is %v\n", y)     
} 

Output:
Type of variable x is int
Value of x in hexdecimal is FF
Value of x in decimal is 255
Type of variable y is int
Value of y in hexdecimal is 9C
Value of y in decimal is 156

Example :5
package main 
import "fmt"
func main() { 
var x int = 5748   
var p *int
 p = &x 
 fmt.Println("Value stored in x = ", x) 
fmt.Println("Address of x = ", &x) 
fmt.Println("Value stored in variable p = ", p) 
} 

Output:
Value stored in x =  5748
Address of x =  0xc42000e1f8
Value stored in variable p =  0xc42000e1f8

Example :6
package main 
import "fmt"
func main()
{
    var y = 458 
    var p = &y 
    fmt.Println("Value stored in y before changing = ", y)
    fmt.Println("Address of y = ", &y) 
    fmt.Println("Value stored in pointer variable p = ", p) 
    fmt.Println("Value stored in y(*p) Before Changing = ", *p)
    *p = 500 
     fmt.Println("Value stored in y(*p) after Changing = ",y) 
  
} 

Output:
Value stored in y before changing =  458
Address of y =  0xc42000e1f8
Value stored in pointer variable p =  0xc42000e1f8
Value stored in y(*p) Before Changing =  458
Value stored in y(*p) after Changing =  500

Example:7
package main
import ("fmt")
func change(val *int) {  
*val = 60
}
func main() {  
    a := 58
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}
Output:
value of a before function call is 58
value of a after function call is 60

Example:8
package main
import ("fmt")
func helloworld() *int {  
    i := 10
    return &i
}
func main() {  
    d := helloworld()
    fmt.Println("Value of d", *d)
}
Output:
Value of d 10

Example:9
package main
import ("fmt")
func main() {  
    b := 278
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++
    fmt.Println("new value of b is", b)
}
Output:
address of b is 0xc42008e168
value of b is 278
new value of b is 279

Example:10
package main
import ("fmt")
func modify(arr *[3]int) {  
    arr[0] = 90
}
func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}

Output:
[90 90 91]