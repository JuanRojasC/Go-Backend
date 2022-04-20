# Testing

## Software Quality

Es el grado con el que un sistema, componente o proceso cumple los requerimientos especificados y las necesidades o expectativas del cliente o usuario. (IEEE)

### Quality Vision

Los requerimientos puden ser *funcionales* como *no funcionales*, cumplir los funcionales nos ayuda a saber si desde la perspectiva del cliente si el software tiene buena calidad, mientras que cumplir los no funcionales comprendemos la calidad desde una perspectiva mas ingenieril.

### Quality Dimensions

* **Funcionality:** Functional completeness
* **Performance:** Time-behaviour, resource utilisation, capacity
* Compatibility
* Usability
* **Reliability:** Maturity, availability, fault tolerance, recoverability
* **Security:** Confidentially, integrity
* Maintainability
* Portability

### Quality Code

#### Maintainability

Se analiza la calidad desde una perspectiva del código fuente del software, teniendo en cuenta métricas como la cyclomatic complexity, code coverage, duplicate code y la adhesion a buenas prácticas del lenguaje.

### Clean Code

El código limpio es simple y directo, y se lee como prosa bien escrita que no oscurece las intenciones del diseãdor, sino que está lleno de abstracciones claras. (Grady Booch). Pasa los tests, No tiene duplicidades, Expresa las ideas de diseño del sistema. Minimiza el número de entidades como clases, métodos y similares. (Ron Jeffries)

## Testing Software

Conjunto de procesos métodos y herramientas para identificar defectos en el software alcanzando un proceso de estabilidad del mismo.

### Black Box

El diseño o funcionamiento interno, no es visible para quien ejecuta la prueba, su objetivo principal es probar la funcionalidad del código, evaluando las prespuestas y reacciones del componente testeado ante distintos escenarios.

* validar que el software cumpla el requirimiento dado
* Test de Regresión: pruebas que se ejecutan luego de cualquier cambio en el código, para verificar que los cambios no hayan afectado ottras funcionalidades

### White Box

Quien ejecuta la prueba conoce y tiene la visibilidad sobre el código. Se centra en os detalles procedimentales del software. Se aplican para probar el flujo, seguridad y estrcutura del código, y asi poder detectar vulnerabilidades en este, comporbar la correcta implementación de cada método o función y validad que el flujo de datos se comporte de acuerdo a lo esperado (condicionales, iteraciones y respuestas).

* Test unitarios
* Tests de integración

## Tests Types

### Test Unitarios

Prueban un fragmento de código, aislandolo y comprobando que funciona a la perfección. Su cualidad es se pequeños y valdidar el comportamiento de un objeto y la lógica.

* Facilitar los cambios en el código
* Encontrar bugs
* Proveen documentación
* Mejoran el diseño y la calidad del código

### Test de Integración

Comprueban la comunicación entre distintos componentes o capas de la aplicación. El objetivo es comprobar que todos aquellos bloques de código que fueron probados de forma unitaria, interactúen y se comuniquen entre sí generando los resultados especados.

* Pueden probar la interacción con una o múltiples bases de datos.
* Asegurar que los microservicios operen como se espera

### Test Funcionales

Se basan en la entrada y salida del software, con el objetivo final que la respuesta del software ante cada escenario, coincida exactamente con el resultado esperado. El qué de lo que hay que cumplir, no el cómo.

## Software Principles

### SOLID

* **S:** Single responsability, un objeto debería tener una única responsabilidad
* **O:** Open/close, los objetos deben estar abiertos para su extensión, pero cerrados para su modificación
* **L:** Liskov substitution, los objetos deberían ser reemplazables por objetos hijos sin alterar el correcto funcionamiento del programa
* **I:** Interface segregation, muchas interfaces más especifícas son mejores que una interfaz de propósito general
* **D:** Dependency inversion desacoplar los objetos abstractos de sus implementaciones

## Testing Principles

### F.I.R.S.T

* **Fast:** Los test deben ser räpidos de correr
* **Isolated:** Los test deben ser cumplor con los 3A arrange, act, assert, no deben ser corridos en un determinado ordena para funcionar
* **Repeatable:** Los test deben ser repetibles, son deterministicos
* **Sel-Validating:** No debe ser requerida una inspección manual para validar los resultados
* **Thorough:** Completos, deben cubrir cada escenario, probar mutaciones, edge cases, exccepciones y errores, no solo es buscar un 100% del coverge

