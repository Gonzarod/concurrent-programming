package main

import "fmt"

func main() {

	myArray:=[7]int{1,89,82,654,221,45,78}

	//Primera forma
	//i: index, j: value
	for i,v :=  range myArray {
		fmt.Printf("El valor es %d en el indice %d \n",v,i)
	}
	fmt.Println()
	//Segunda forma
	//i: index, j: value
	for _,v :=  range myArray {
		fmt.Printf("El valor es %d \n",v)
	}
	fmt.Println()
	//Tercera forma
	//i: index, j: value
	i:=5
	for {
		if i==5{
			fmt.Println("Finshed")
			break
		}
		i++

	}
	fmt.Println("Numero de intentos i:",i)
	fmt.Println()

	j:=0
	for i>0{
		if j==5{
			break
		}
		j++
	}
	fmt.Println("Numero de intentos j:",j)
}
