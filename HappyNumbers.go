package main
import (
	"fmt"
	"strconv"
	"strings"
)


func main(){
	var number int
	fmt.Print("Escribe un numero entero mayor a cero: ")
	_, err := fmt.Scan(&number)
	
	for err != nil || number <= 0 {		
		fmt.Print("Se esperaba un numero entero mayor a cero. Intenta de nuevo: ")
		_, err = fmt.Scan(&number)
	}

	var result = "infeliz :("
	if(IsHappyNumber(number, 0)){
		result = "feliz :)"
	}
		
	fmt.Printf("El numero %d es un numero %s", number, result)
}

func IsHappyNumber(number, tries int) bool{
	var strNumber = strconv.Itoa(int(number))
	var strSlice = strings.Split(strNumber, "")

	var total int
	for _,value := range strSlice {
		intValue,_ := strconv.Atoi(value)
		total += intValue * intValue
	}

	if(total == 1){
		return true
	}

	if(tries == 10){
		return false
	}	
	
	return IsHappyNumber(total, tries + 1)
}