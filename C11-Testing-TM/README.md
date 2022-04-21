# Double Test

Es un tipo de test unitario, con su uso, se busca porbar porciones de código que tengan dependencias de otros componentes, creando pequeñas piezas que suplan dicha dependencia.

## Types

* **Dummy:** Algo que se utiliza para satisfacer dependencias, pero su uso en ejecucuión es completamente irrelevante. Se usa en aquellos casos donde el objeto requerido no interviene en la ejecución de nuestro algoritmo.
  * *Si la dependencia solo es necesaria para la creación del objeto a probar, hacemos un dummy ya que este no es lo testeable, pero es obligatorio para porder llevar a cabo el test, un ejemplo seria un dummy para un campo al momento de crear una estructura, pero ese campo no sera usado o ejecutado en el test*

* **Stub:** Es similar al dummy, con la diferencia de que este retornará calores concretos, para guiar la ejecución del código por un camino determinado.
  * *En este caso la depedencia si sera utilizada en la ejecución, por ya no es un dummy pues debera retornar algo concreto*

* **Spy:** Nos permite comprobar o asegurar de haber llamado a un método para dar el test como válido, basicamente nos informa cualdo algo sucede.
  * *Se podria generar una strcut con campo boleanos que nos permiten determinar si un metodo fue o no llamado de acuerdo a su valor, y el metodo testado modificamos el valor*

* **Mock:** Se utiliza para comprobar tpdp el funcionamiento interno del método o código que se está testeando, es una mezcla entre un stub y un spy.
  * *El mock verifica el resultado esperado y que el método encargado de generarlo haya sido efectivamente llamado* [GitHub](https://gist.github.com/IgnacioFalco/d6d340363201db76076b9c3b29b2e4c8)

* **Fake:** Son objetos que tienen implementación funcional, pero no se usará en producción, suelen ser acercamientos simplificados al comportamiento real del componente.
  * *Simular una base de datos con un map, que puda permitirnos relizar operaciones como si de la BD se tratara*

## Consideraciones

* Un spy es un tipo de stub, pero con la responsabilidad adicional de guardar información
* Un mock es una clase de Spy, pero en el que las validaciones se hacen sobre la información guardada en el Mock
* Un fake se distingue porque contiene lógica de negocia, y devuelve disintas respuestas de acuerdo al escenario
* Los tests deben ser independientes, cada test contendra su arrange necesario
