EXAMPLE 1 :
package main
import "fmt"
func SimpleFunction() {
fmt.Println("Hello Pavankumar")
}
func main() {
SimpleFunction()
}

OUTPUT :
Hello Pavankumar

Example 2:
package main
import "fmt"
func add(x int, y int) int {
total := 0
total = x + y
return total
}
func main() {
sum := add(80,40)
fmt.Println(sum)
}

Output:
120

Example 3:
package main
import "fmt"
func main() {
var a int = 1000
var b int = 2000
var ret int
ret = max(a, b)
fmt.Printf( "Max value is : %d\n", ret )
}
func max(num1, num2 int) int {
var result int
if (num1 > num2) {
result = num1
} else {
result = num2
}
return result 
}

Output:
Max value is : 2000

Example 4:
package main
import "fmt"
func swap(x, y string) (string, string) {
return y, x
}
func main() {
a, b := swap("Pavankumar", "Keerthi")
fmt.Println(a, b)
}

OUTPUT :
Keerthi Pavankumar

Example 5:
package main
import "fmt"
func add(x int, y int) int {
return x + y
}
func main() {
fmt.Println(add(70, 13))
}

Output:
83

Example 6:
package main
import "fmt"
func plus(a int, b int) int {
return a + b
}
func plusPlus(a, b, c int) int {
return a + b + c
}
func main() {
res := plus(1, 2)
fmt.Println("1+2 =", res)
res = plusPlus(1, 2, 3)
fmt.Println("1+2+3+4+6 =", res)
}

Output:
1+2 = 3
1+2+3+4+6 = 16

Example 7:
package main
import "fmt"
func rectangle(l int, b int) (area int) {
var parameter int
parameter = 2 * (l + b)
fmt.Println("Parameter: ", parameter)
area = l * b
return 
}

func main() {
fmt.Println("Area: ", rectangle(30, 30))
}

Output:
Parameter:  180
Area:  900

Example 8:
package main

import "fmt"
type First func(int) int
type Second func(int) First
func squareSum(x int) Second {
return func(y int) First {
return func(z int) int {
return x*x + y*y + z*z
}
}
}
func main() {
fmt.Println(squareSum(5)(6)(4))
}

Output :
77

Example 9:
package main
import "fmt"
func sum(x, y int) int {
return x + y
}
func partialSum(x int) func(int) int {
return func(y int) int {
return sum(x, y)
}
}
func main() {
partial := partialSum(3)
fmt.Println(partial(7))
}  

Output :
10

Example 10:

package main
import "fmt"
func main() {
l := 30
b := 30
func() {
var area int
area = l * b
fmt.Println(area)
}()
}

Output :
900


Example 11:

package main
import "fmt"
func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
sum := add(100, 200)
fmt.Println(sum)
}

Output :
2000