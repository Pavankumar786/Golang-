Example 1: Calculating Factorial using recursion
package main
import "fmt"
func factorial(i int)int {
if(i <= 1) {
 return 1
}
return i * factorial(i - 1)
}
func main() { 
 var i int = 15
fmt.Printf("Factorial of %d is %d", i, factorial(i))
}

Output:
Factorial of 15 is 1307674368000

Exxample 2:fibonacci series using Recursion
package main
import "fmt"
func fibonaci(i int) (ret int) {
if i == 0 {
return 0
 }
if i == 1 {
return 1
}
return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
var i int
for i = 0; i < 10; i++ {
 fmt.Printf("%d ", fibonaci(i))
}
}


Output:
0 1 1 2 3 5 8 13 21 34 


Example 3:
package main
import "fmt"
func fibonaci(i int) (ret int) {
if i == 0 {
return 0
 }
if i == 1 {
return 1
}
return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
var i int
for i = 0; i < 30; i++ {
 fmt.Printf("%d ", fibonaci(i))
}
}


Output:
0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368 75025 121393 196418 317811 514229 


Example 4:
package main
import "fmt"
func factorial(i int)int {
if(i <= 1) {
 return 1
}
return i * factorial(i - 1)
}
func main() { 
 var i int = 98
fmt.Printf("Factorial of %d is %d", i, factorial(i))
}

Output:
Factorial of 98 is 0


