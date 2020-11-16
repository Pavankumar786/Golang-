Example 1: Go Arrays
package main
import "fmt"
func main() {
var n [10]int 
var i,j int
for i = 0; i < 10; i++ {
n[i] = i + 100 
}
for j = 0; j < 10; j++ {
fmt.Printf("Element[%d] = %d\n", j, n[j] )
}
}

Output:
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109


Example 2:
package main
import "fmt"
func main() {
var n [15]int 
var i,j int
for i = 0; i < 15; i++ {
n[i] = i + 100 
}
for j = 0; j < 15; j++ {
fmt.Printf("Element[%d] = %d\n", j, n[j] )
}
}

output:
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
Element[10] = 110
Element[11] = 111
Element[12] = 112
Element[13] = 113
Element[14] = 114


Example 3:
package main
import "fmt"
func main() {
var n [25]int 
var i,j int
for i = 0; i < 25; i++ {
n[i] = i + 100 
}
for j = 0; j < 25; j++ {
fmt.Printf("Element[%d] = %d\n", j, n[j] )
}
}

Output:
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
Element[10] = 110
Element[11] = 111
Element[12] = 112
Element[13] = 113
Element[14] = 114
Element[15] = 115
Element[16] = 116
Element[17] = 117
Element[18] = 118
Element[19] = 119
Element[20] = 120
Element[21] = 121
Element[22] = 122
Element[23] = 123
Element[24] = 124


Example 4:Multidimensional array
package main
import "fmt"
func main() {
var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
var i, j int
for  i = 0; i < 5; i++ {
for j = 0; j < 2; j++ {
fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
}
}
}

Output:
a[0][0]: 0
a[0][1]: 0
a[1][0]: 1
a[1][1]: 2
a[2][0]: 2
a[2][1]: 4
a[3][0]: 3
a[3][1]: 6
a[4][0]: 4
a[4][1]: 8

Example 5:
package main
import "fmt"
func main() {
var  balance = []int {1000, 2, 3, 17, 50}
var avg float32
avg = getAverage( balance, 5 ) ;
fmt.Printf( "Average value is: %f ", avg );
}
func getAverage(arr []int, size int) float32 {
var i,sum int
var avg float32  
for i = 0; i < size;i++ {
sum += arr[i]
}
avg = float32(sum / size)
return avg;
}

Output:
Average value is: 214.400000

Example 6:
package main
import "fmt"
func main() {
var  balance = []int {1000000, 2, 3, 17, 50}
var avg float32
avg = getAverage( balance, 5 ) ;
fmt.Printf( "Average value is: %f ", avg );
}
func getAverage(arr []int, size int) float32 {
var i,sum int
var avg float32  
for i = 0; i < size;i++ {
sum += arr[i]
}
avg = float32(sum / size)
return avg;
}

Output:
Average value is: 200014.000000 


Example 7:
package main
import "fmt"
func main() {
var n [10]int 
var i,j int
for i = 0; i < 10; i++ {
n[i] = i + 100 
}
for j = 0; j < 10; j++ {
fmt.Printf("Element[%d] = %d\n", j, n[j] )
}
}

Output:
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109



Example 8:
package main 
import "fmt"
func main() { 
var myarr[3]string
myarr[0] = "Pavan"
myarr[1] = "kumar"
myarr[2] = "Keerthi"
fmt.Println("Elements of Array:") 
fmt.Println("Element 1: ", myarr[0]) 
fmt.Println("Element 2: ", myarr[1]) 
fmt.Println("Element 3: ", myarr[2]) 
} 

Output:
Elements of Array:
Element 1:  Pavan
Element 2:  kumar
Element 3:  Keerthi


Example 9:
package main 
import "fmt"
func main() { 
arr:= [3][3]string{{"C#", "C", "Python"},{"Java", "Scala", "Perl"}, {"C++", "Go", "HTML"},} 
fmt.Println("Elements of Array 1") 
for x:= 0; x < 3; x++{ 
for y:= 0; y < 3; y++{ 
fmt.Println(arr[x][y]) 
} 
} 
var arr1 [2][2] int
arr1[0][0] = 100 
arr1[0][1] = 200 
arr1[1][0] = 300 
arr1[1][1] = 400 
fmt.Println("Elements of array 2") 
for p:= 0; p<2; p++{ 
for q:= 0; q<2; q++{ 
fmt.Println(arr1[p][q]) 
} 
} 
} 
Output:
Elements of Array 1
C#
C
Python
Java
Scala
Perl
C++
Go
HTML
Elements of array 2
100
200
300
400

Example 10:
package main 
import "fmt"
func main() { 
var myarr[2]int 
fmt.Println("Elements of the Array :", myarr) 
} 
Output:
Elements of the Array : [0 0]