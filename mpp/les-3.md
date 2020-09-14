# Go

Go is erg vergelijkbaar met c zonder de `terror` en gemaakt door mensen bij google

start altijd een go bestand met package main (zegt dat het executable wordt)

Hierna volgen de import statements, onderandere altijd "fmt"(format) die handled printen bijv.

En er is dan altijd een main functie die wordt uitgevoerd
```go
package main

import (
    "ftm"
)

func main() {
    fmt.Println('Hello, World!'
}
```

variabelen kan je instellen als
```go
    var name int;
    name = 1;
```
of als 
```go
    name := 1;
```

Example code:
```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
    variables()
}

func variables() {
    // eerste optie
    var x int
    x = 42

    // Tweede optie
    y := "Pepijn"

    fmt.Println(x)
    fmt.Println(y)
}

// dit kan met (x int, y int)
func namedReturns(x, y int) (z int) {
    // kan dus dit, want go weet dat je z returnt
    z = x+y
    return
   // had ook return x+y; kunnen doen
}

```

## arrays, loopen en if

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello WOrld!")

    i := 1;

    //for loop
    for i < 10 {
        fmt.Println(i)
        i++
    }

    //while loop
    for {
        fmt.Println("smthn")
    }

    // if condities
    if 10 > 5 {
        fmt.Println("yay")
    } else {
        fmt.Println("whut")
    }

    // arrays
    // array met 10 ints
    var array [10]int
    fmt.Println("empty aray: ",array)

    array[2] = 10;
    fmt.Println(array, len(array))

    //multidimentional
    var twoDim [3][3]int
    twoDim[0][2] = 8

    fmt.Println(twoDim)
}
```

## errorhandlin

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("hello")

    file, err := os.Open("filename.txt")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(file)
}
```

Tip: Oefen misschien [gobyexample](gobyexample.com/) tot en met errors en misschien defer en HTTP
