# HappyNumbers

Los números felices se definen por el siguiente procedimiento: empezando con cualquier número entero positivo, se reemplaza el número por la suma de los cuadrados de sus dígitos, y se repite el proceso hasta que el número es igual a 1 o hasta que se entra en un bucle que no incluye el 1. Los números que al finalizar el proceso terminan con 1, son conocidos como números felices. Aquellos que no, son conocidos como números infelices (o tristes). 


### Ejemplo
<b>7 es un número feliz</b>, ya que: <br/>
7<sup>2</sup> = 49 <br/>
4<sup>2</sup> + 9<sup>2</sup> = 97 <br/>
9<sup>2</sup> + 7<sup>2</sup> = 130 <br/>
1<sup>2</sup> + 3<sup>2</sup> + 02 = 10 <br/>
1<sup>2</sup> + 0<sup>2</sup> = 1 <br/>

Fuente: <a href="https://es.wikipedia.org/wiki/N%C3%BAmero_feliz">Wikipedia</a>

### Ejecucion
`go run HappyNumbers.go`
