Example 1:Interfaces

package main
import ("fmt" "math")
type Shape interface {
area() float64
}
type Circle struct {
x,y,radius float64
}
type Rectangle struct {
width, height float64
}
func(circle Circle) area() float64 {
return math.Pi * circle.radius * circle.radius
}
func(rect Rectangle) area() float64 {
return rect.width * rect.height
}
func getArea(shape Shape) float64 {
return shape.area()
}
func main() {
circle := Circle{x:0,y:0,radius:5}
rectangle := Rectangle {width:10, height:5}
fmt.Printf("Circle area: %f\n",getArea(circle))
fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
}

Output:
Circle area: 78.539816
Rectangle area: 50.000000


Example 2:
package main  
import ("fmt")  
type vehicle interface {  
accelerate()  
}  
func foo(v vehicle)  {  
fmt.Println(v)  
}  
type car struct {  
model string  
color string  
}  
func (c car) accelerate()  {  
fmt.Println("Accelrating")  
}  
type toyota struct {  
model string  
color string  
speed int  
}  
func (t toyota) accelerate(){  
fmt.Println("I am toyota, I accelerate fast")  
}  
func main() {  
c1 := car{"suzuki","blue"}  
t1:= toyota{"Toyota","Red", 100}  
c1.accelerate()  
t1.accelerate()  
foo(c1)  
foo(t1)  
}  


Output:
Accelrating
I am toyota, I accelerate fast
{suzuki blue}
{Toyota Red 100}


Example 3:
package main
import ("fmt")
import	("math")
type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
width, height float64
}
type circle struct {
radius float64
}
func (r rect) area() float64 {
return r.width * r.height
}
func (r rect) perim() float64 {
return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
return 2 * math.Pi * c.radius
}
func measure(g geometry) {
fmt.Println(g)
fmt.Println(g.area())
fmt.Println(g.perim())
}
func main() {
r := rect{width: 3, height: 4}
c := circle{radius: 5}
measure(r)
measure(c)
}


Output:
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793


Example 4:
package main 
import "fmt"
type tank interface { 
Tarea() float64 
Volume() float64 
} 
type myvalue struct { 
radius float64 
height float64 
} 
func (m myvalue) Tarea() float64 { 
return 2*m.radius*m.height + 2*3.14*m.radius*m.radius 
} 
func (m myvalue) Volume() float64 { 
return 3.14 * m.radius * m.radius * m.height 
} 
func main() { 
var t tank 
t = myvalue{10, 14} 
fmt.Println("Area of tank :", t.Tarea()) 
fmt.Println("Volume of tank:", t.Volume()) 
}
 
Output:
Area of tank : 908
Volume of tank: 4396

Example 5:
package main 
import "fmt"
type tank interface { 
Tarea() float64 
Volume() float64 
} 
type myvalue struct { 
radius float64 
height float64 
} 
func (m myvalue) Tarea() float64 { 
return 2*m.radius*m.height + 2*3.14*m.radius*m.radius 
} 
func (m myvalue) Volume() float64 { 
return 3.14 * m.radius * m.radius * m.height 
} 
func main() { 
var t tank 
t = myvalue{23, 89} 
fmt.Println("Area of tank :", t.Tarea()) 
fmt.Println("Volume of tank:", t.Volume()) 
} 

Output:
Area of tank : 7416.12
Volume of tank: 147834.34


Example 6:
package main 
import "fmt"
func myfun(a interface{}) { 
value, ok := a.(float64) 
fmt.Println(value, ok) 
} 
func main() { 
var a1 interface { 
} = 98.09 
myfun(a1) 
var a2 interface { 
} = "welcome"
myfun(a2) 
} 

Output:
98.09 true
0 false


Example 7:
package main 
import "fmt"
func myfun(a interface{}) { 
switch a.(type) { 
 case int: 
fmt.Println("Type: int, Value:", a.(int)) 
 case string: 
 fmt.Println("\nType: string, Value: ", a.(string)) 
 case float64: 
 fmt.Println("\nType: float64, Value: ", a.(float64)) 
 default: 
 fmt.Println("\nType not found") 
} 
}  
func main() { 
  
    myfun("welcome") 
    myfun(67.9) 
    myfun(true) 
} 

Output:
Type: string, welcome
Type: float64, Value:  67.9
Type not found

Example 8:
package main 
import "fmt"
type tank interface { 
Tarea() float64 
Volume() float64 
} 
func main() { 
var t tank 
fmt.Println("Value of the tank interface is: ", t) 
fmt.Printf("Type of the tank interface is: %T ", t) 
} 


Output:
Value of the tank interface is:  <nil>
Type of the tank interface is: <nil> 


