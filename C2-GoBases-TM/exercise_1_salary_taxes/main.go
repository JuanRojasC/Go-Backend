package main

import . "fmt"

/* Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%. */

func main() {

	employee1 := 46700
	employee2 := 87460
	employee3 := 189000

	employees := []int{int(employee1), int(employee2), int(employee3)}

	for _, v := range employees {
		Printf("El empleado con salario %d debe pagar %0.2f en impuestos\n", v, salaryTaxes(float64(v)))
	}

}

func salaryTaxes(salary float64) float64 {
	basicTax := 0.17
	if salary > 150000 {
		return salary * (basicTax + 0.1)
	} else if salary > 50000 {
		return salary * basicTax
	} else {
		return 0
	}
}
