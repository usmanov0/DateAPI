// Bu Task 2-etapdan ohirida bitta qolgan edi.
package main
import ("fmt"
        "math")

func kvadratEquation(a,b,c float64){
    var d float64
    
    d = b * b - 4*a*c
    
    if d > 0 {
        sol1 := (-b + math.Sqrt(d)) / (2 * a)
        sol2 := (-b - math.Sqrt(d)) / (2 * a)
        fmt.Printf("Solutions: %f and %f\n", sol1, sol2)
    }else if d == 0{
        sol := -b / (2 * a)
        fmt.Println(sol)
    }else {
        r1 := -b / (2 * a)
        r2 := math.Sqrt(math.Abs(d)) / (2 * a) 
        fmt.Println("First root", r1, "+", "i", r2)
        fmt.Println("Second root", r1, "-", "i", r2)
    }
}
func main() {
    var a,b,c float64
    a = 4
    b = 12
    c = 9
    
    fmt.Println("a =", a)
    fmt.Println("b =", b)
    fmt.Println("c =", c)
    kvadratEquation(a, b, c)
}
