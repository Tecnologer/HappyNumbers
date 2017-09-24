# HappyNumbers

Los números felices se definen por el siguiente procedimiento: empezando con cualquier número entero positivo, se reemplaza el número por la suma de los cuadrados de sus dígitos, y se repite el proceso hasta que el número es igual a 1 o hasta que se entra en un bucle que no incluye el 1. Los números que al finalizar el proceso terminan con 1, son conocidos como números felices. Aquellos que no, son conocidos como números infelices (o tristes). 


#Ejemplo

<b>7 es un número feliz</b>, ya que: <br/>
72 = 49 <br/>
42 + 92 = 97 <br/>
92 + 72 = 130 <br/>
12 + 32 + 02 = 10 <br/>
12 + 02 = 1 <br/>

Fuente: <a href="https://es.wikipedia.org/wiki/N%C3%BAmero_feliz">Wikipedia</a>

#Ejecucion<br/>
`go run HappyNumbers.go`
