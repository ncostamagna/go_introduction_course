# Indice
- [Variables](#variables)
  * [Int](#int)
  * [Uint](#uint)
  * [Float](#float)
  * [String](#string)
- [Operadores Logicos](#operadores-logicos)
  * [Or](#or)
  * [And](#and)
- [Vetores](#vetores)
  * [Array](#array)
  * [Slice](#slice)
- [Maps](#maps)
- [For](#for)

# Variables
Podemos definir una variable indicando el tipo de dato
```go
var int myVar 
myVar = 64
```
O podemos utilizar el dos puntos igual (:=) para instanciar y asignar al mismo tiempo, va a tomar el tipo de dato del valor que le pasamos
```go
myVar := 64 
```
## Int

```go
var int64 myVar 
myVar = 64
```

|Tipo|Tamaño|Rango|
|----|------|-----|
|int8|8 bits|-128 hasta 127|
|int16|16 bits|-2<sup>15</sup> hasta 2<sup>15</sup> -1|
|int32|32 bits|-2<sup>31</sup> hasta 2<sup>31</sup> -1|
|int64|64 bits|-2<sup>63</sup> hasta 2<sup>63</sup> -1|
|int|Depende plataforma|Depende plataforma|


## Uint

```go
var uint64 myVar 
myVar = 64
```

|Tipo|Tamaño|Rango|
|----|------|-----|
|uint8|8 bits|0 hasta 255|
|uint16|16 bits|0 hasta 2<sup>16</sup> -1|
|uint32|32 bits|0 hasta 2<sup>32</sup> -1|
|uint64|64 bits|0 hasta 2<sup>64</sup> -1|
|uint|Depende plataforma|Depende plataforma|


## Float

```go
var float64 myVar 
myVar = 64.55
```

|Tipo|Tamaño|
|----|------|
|float32|32 bits|
|float64|64 bits|

## String
Los **strings** son cadenas de bytes

```go
var string myVar 
myVar = "My string"
```

# Operadores Logicos

## Or
|A|B|A or B|
|----|------|-----|
|false|false|false|
|false|true|true|
|true|false|true|
|true|true|true|

```go
myValue1 := 3
myValue2 := 5

fmt.Println(myValue1 == myValue2 || myValue1 < myValue2) // (false or true) = true
```

## And
|A|B|A and B|
|----|------|-----|
|false|false|false|
|false|true|false|
|true|false|false|
|true|true|true|

```go
myValue1 := 3
myValue2 := 5

fmt.Println(myValue1 == myValue2 && myValue1 < myValue2) // (false and true) = false
```

# Vetores

## Array
Tamaño fijo

```go
	var myArrayVar [6]int
	fmt.Println(myArrayVar)

	myArrayVar1 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar1)

	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9
	fmt.Println(myArrayVar)
```

## Slice
Tamaño dinamico
```go
var slice1 []int

slice1 = append(slice1, 10, 20, 30, 40, 50)
fmt.Printf("size: %d, value %v\n", len(slice1), slice1)
```

# Maps

```go
// Declaracion e instancia de la variable
m1 := make(map[int]string)

// Asignamos valores
m1[1] = "A"
m1[2] = "B"
m1[3] = "C"

// Mostramos el valor de la key 1
fmt.Println(m1[1])
```

# For

```go
sum := 0
for i := 0; i < 10; i++ {
	sum++
}


arr := []int{1, 2, 3, 1, 2, 3}
for i, v := range arr {
	fmt.Println("Index: ",i ," - Value: ", v)
}

map2 := map[string]float64{
	"A": 12.3,
	"Z": 23.1,
	"C": 34,
}
for key, value := range map2 {
	fmt.Println("Key:", key, "Value:", value)
}
```

# Functions

```go
// x, y are int parameters & return int value
func MiFunction(x, y int) int {
	return x + y
}
```