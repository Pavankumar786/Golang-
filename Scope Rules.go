Example 1: Local Variables
package main
import "fmt"
func main() {
var a, b, c int 
a = 10
b = 20
c = a + b
fmt.Printf ("value of a = %d, b = %d and c = %d\n", a, b, c)
}

Output:
value of a = 10, b = 20 and c = 30

Example 2: Global Variables
package main
import "fmt"
var g int
 func main() {
 var a, b int
 a = 10
 b = 20
 g = a + b
fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)
}

Output:
value of a = 10, b = 20 and g = 30


Example 3: Formal Parameters
package main
import "fmt"
var a int = 20;
func main() {
var a int = 10
var b int = 20
var c int = 0
fmt.Printf("value of a in main() = %d\n",  a);
c = sum( a, b);
fmt.Printf("value of c in main() = %d\n",  c);
}
func sum(a, b int) int {
fmt.Printf("value of a in sum() = %d\n",  a);
fmt.Printf("value of b in sum() = %d\n",  b);
 return a + b;
}

Output:
value of a in main() = 10
value of a in sum() = 10
value of b in sum() = 20
value of c in main() = 30
package main 
 
Example 4:
import "fmt"
func main() { 
var myvariable1, myvariable2 int = 89, 45 
fmt.Printf("The value of myvariable1 is : %d\n",myvariable1)  
fmt.Printf("The value of myvariable2 is : %d\n",myvariable2)  
   
}

Output:
The value of myvariable1 is : 89
The value of myvariable2 is : 45


Example 5:
package main
import "fmt"
var g int = 20
func main() {
var g int = 10
fmt.Printf ("value of g = %d\n",  g)
}

Output:
value of g = 10

Example 6:
package main  
import "fmt"
var myvariable1 int = 100 
func main() { 
var myvariable2 int = 200 
fmt.Printf("The value of Global myvariable1 is : %d\n",myvariable1)  
fmt.Printf("The value of Local myvariable2 is : %d\n",myvariable2)  
display() 
  } 
func display() {  
fmt.Printf("The value of Global myvariable1 is : %d\n",myvariable1)  
} 
Output:
The value of Global myvariable1 is : 100
The value of Local myvariable2 is : 200
The value of Global myvariable1 is : 100

Example 7:
package main  
import "fmt"
var myvariable1 int = 100 
func main() { 
var myvariable1 int = 200 
fmt.Printf("The value of myvariable1 is : %d\n",myvariable1)  
} 

Output:
The value of myvariable1 is : 200

Example 8:
package main
import "fmt"
var g int
 func main() {
 var a, b int
 a = 10
 b = 20
 g = a + b
fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)
}

Output:
value of a = 10, b = 20 and g = 30
