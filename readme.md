# Indice
- [Variables](#variables)
  * [Int](#int)
  * [Uint](#uint)
  * [Float](#float)
  * [String](#string)
- [Operadores Logicos](#operadores-logicos)
  * [Or](#or)
  * [And](#and)


# Variables

## Int
|Tipo|Tamaño|Rango|
|----|------|-----|
|int8|8 bits|-128 hasta 127|
|int16|16 bits|-2<sup>15</sup> hasta 2<sup>15</sup> -1|
|int32|32 bits|-2<sup>31</sup> hasta 2<sup>31</sup> -1|
|int64|64 bits|-2<sup>63</sup> hasta 2<sup>63</sup> -1|
|int|Depende plataforma|Depende plataforma|


## Uint
|Tipo|Tamaño|Rango|
|----|------|-----|
|uint8|8 bits|0 hasta 255|
|uint16|16 bits|0 hasta 2<sup>16</sup> -1|
|uint32|32 bits|0 hasta 2<sup>32</sup> -1|
|uint64|64 bits|0 hasta 2<sup>64</sup> -1|
|uint|Depende plataforma|Depende plataforma|


## Float

|Tipo|Tamaño|
|----|------|
|float32|32 bits|
|float64|64 bits|

## String
Los **strings** son cadenas de bytes

# Operadores Logicos

## Or
|A|B|A or B|
|----|------|-----|
|false|false|false|
|false|true|true|
|true|false|true|
|true|true|true|

## And
|A|B|A and B|
|----|------|-----|
|false|false|false|
|false|true|false|
|true|false|false|
|true|true|true|