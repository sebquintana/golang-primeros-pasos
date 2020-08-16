# Primeros pasos en Golang

## TP1
El objetivo de este TP es que practiques cómo iterar slices, mapas, definir métodos, correr tests, user interfaces  y más!

Para los ejercicios de este TP vamos a estar usando un dataset que contiene un conjunto de productos formados de la siguiente manera:
```
[
	["dia", "1", "234"]
]
```

Cada ejercicio tiene una funcion definida en `tp1/tp1.go` que vas a tener que implementar para que sus tests, definidos en `tp1/tp1_test.go`, pasen exitosamente. Estas son las consignas:
* **CalcularPrecios**: Dada una lista de ids de productos la función debería calcular cuál sería el precio total de todos los productos para cada uno de los supermercados. Basicamente, armar un "carrito" para cada super mercado que se encuentra e indicar cuanto saldria comprar esos items en ese super mercado,
* **Promedio**: Dado el id de un producto la función debería calcular cuál es el precio promedio de ese producto utilizando la data de todos los supermercados,
* **BuscarMasBarato**: Recibe el ID de un producto y debería retornar cuál es el super mercado que lo vende más barato y a cuánto lo vende.

La funcion `BuscarMasBarato` retorna una interfaz llamada `Producto`. Vas a tener que definir algun tipo que cumpla con la definicion de esa interfaz para poder resolver el ejercicio.

![](images/tp1.jpeg)

### ¿Cómo probar tu solución?
Cada una de las funciones definidas en `tp1.go` tiene una `func` escrita en `tp1_test.go` que actúa como *test* de esa función. Ya tenemos un conjunto de casos identificados y escritos en los tests que validan que tu codigo haga lo que corresponde.  

Para validar tus soluciones, podés ejecutar los siguientes comandos parándote en la carpeta de `tp1`:

```
# correr los tests de la funcion sumar. Aca podrian cambiar
# `CalcularPrecios` por el nombre de la funcion que estan probando en el momento
go test -run=TestProductos_CalcularPrecios

# Correr los tests de todas las funciones
go test
```

Si quieren tener más información de los tests, por ejemplo saber cuáles escenarios fallaron, pueden correr el comando con el flag `-v`:
```
go test -v -run=TestProductos_CalcularPrecios
```

## TP2
El objetivo de este TP es que aprendas como usar las capacidades de concurrencia que nos ofrece Go. Vas a usar channels y goroutines para poder ejecutar y coordinar eficientemente una suma de muchos numeros.

Para realizar la suma de dos numeros ya tenes una funcion a tu disposicion `sumFunc` que recibe dos argumentos como parametro y retorna la suma. El problema es que `sumFunc` tarda una cantidad de tiempo indeterminado (pero no mayor a medio segundo) en retornarte la suma, entonces hacer una suma secuencial de muchos numeros nos puede llegar a tardar mucho tiempo. Para poder hacerla mas eficiente **es necesario usar goroutines y channels** que paralelicen estas operaciones y reduzcan el tiempo que toma sumar todo. Una forma de encarar esto seria usando semáforos y mantener una variable global de suma que cada goroutine incremente, el problema es que esto no esta acorde a uno de los [proverbios mas importantes de Go](https://go-proverbs.github.io/):
> [Don't communicate by sharing memory, share memory by communicating](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s).

Ese proverbio simplemente indica que usar operadores de exclusión mutua y variables compartidas no es lo mejor que se puede hacer. No es una mala herramienta, y a veces es totalmente necesaria, pero para este caso se puede encarar la solución usando canales para la comuniación y reducir el overhead de tener que proteger la memoria manualmente.

La consiga es:
* **SumarLista**: Dada una lista de enteros la funcion deberia calcular la suma de todos ellos y retornar un error en caso de que no se reciban numeros.

### ¿Cómo probar tu solución?
Al igual que en el TP1, definimos tests para la función `SumarLista` que estan escritos en `tp2/tp2_test.go`. Una vez que termines tu solución podés correr el siguiente comando estan parado/a en la carpeta `tp2`:
```
go test -run=TestSumarLista
```

## TP3
El objetivo de este TP es que aprendas a usar la primitiva de `select` que te permite hacer un `switch` no deterministico sobre varios canales para poder recibir mensajes.

La consigna consiste en implementar una calculadora que va a recibir sus operaciones en 4 canales distintos, 1 para cada tipo de operación. Por ahora solo vas a implementar suma, resta, multiplicación y división. También vas a recibir un canal de corte que te va a indicar que ya deberías terminar de procesar. Los resultados los vas a escribir a otro canal que tenes que crear vos, ese canal es el que vas a retornar en la función. Una vez que detectaste a través del canal de corte que no hay mas operaciones deberías cerrar el canal de resultados para indicarle a los usuarios de la función que ya no hay mas resultados que enviar.  
El procesamiento tiene que ser asíncrono, por ende la iteración sobre los canales de operaciones debería correr en una goroutine aparte. Los tests que validan el comportamiento de tu función tienen un timeout de 1 segundo, eso quiere decir que si después de un segundo no se cerro el canal de resultados el test va a fallar. Son operaciones simples que no toman nada de tiempo en ejecutar, 1 segundo es suficiente.

La función a implementar es `Calcular` y se encuentra en `tp3/tp3.go`.

### ¿Cómo probar tu solución?
Al igual que en el TP1 y TP2, definimos tests para la función `Calcular` que estan escritos en `tp3/tp3_test.go`. Una vez que termines tu solución podés correr el siguiente comando estando parado/a en la carpeta `tp3`:
```
go test -run=TestCalcular
```

## TP4
Este TP va a ser distinto a los demás. Ahora vas a tener que implementar un programa completo y no vas a tener tests disponibles que validen la solución. La idea de este TP es que aprendas a escribir web servers en Go desde cero y a comunicarte con APIs externas.

La consigna consiste en implementar algo similar a la función `CalcularPrecios` del primer TP, solo que ahora la funcionalidad va a tener que estar expuesta a través de una API y la información de los productos la vas a obtener haciendo HTTP requests a un servicio que armamos. Tu servicio debe exponer un endpoint que acepte en el request una lista de IDs de productos y una lista de supermercados. La respuesta tiene que ser la suma de precios para cada uno de los supermercados que se pidió en formato de JSON. Si algún producto no se encuentra en algún de supermercado el servicio tiene que responder con un mensaje de error. 

El servicio que retorna los productos se encuentra en `https://productos-p6pdsjmljq-uc.a.run.app` y el formato de las URLs es: `/<tienda>/productos/<id>`, por ejemplo:
```sh
$ curl https://productos-p6pdsjmljq-uc.a.run.app/dia/productos/1
{"tienda":"dia","id":1,"precio":7887}
```

Los supermercados y productos que están disponibles en ese servicio son los mismos que el TP1 y se encuentran en `tp1/productos.json`. Para validar que se esta calculando correctamente el total podes usar los tests escritos en `tp1/tp1_test.go`. En particular la función `TestProductos_CalcularPrecios`, contiene el monto de cuanto sale la suma de los productos con ID 1 y 2 en cada uno de los supermercados disponibles.

El código tiene que estar escrito dentro de la carpeta `tp4` y pueden utilizar paquetes para estructurar su proyecto de manera mas clara. Bonus points si la solución es muy eficiente y utiliza goroutines para obtener los productos en paralelo! :)

## TP FINAL
La idea de este TP es que tengas que desarrollar algo bastante desafiante. Te va a llevar algo de tiempo y vas a tener que investigar bastante y leer mucho. Seguramente incluso tengas que usar paquetes de terceros, etc.
Lo que te podemos garantizar es que después de hacer el TP vas a tener un muy buen dominio del lenguaje :)
La propuesta es que desarrolles alguna de los siguientes 3 proyectos:
  1. Desarrollá el juego [Snake](http://www.menet.umn.edu/~dockt036/snake.html). Idealmente lo hacés en modo terminal. Esto te va a obligar a usar paquetes, goroutines y aplicar lógica.
  2. Desarrollá un programa que reciba funciones polinómicas, tipo `x^2+2x+8`, y las grafique. Para puntos extra podrías calcular las raíces y graficarlas y además dar la posibilidad de hacer zoom in y zoom out.
  3. Desarrollá un loadbalancer http. La idea es que tenga un archivo de configuración (podría ser JSON o cualquier otra cosa) con nodos de backend y puerto. Cuando un request llega a alguno de los puertos configurados, se elige alguno de los nodos de backend que corresponda y se envía el request y por supuesto se debe retorna el response también. **Prestá atención a los headers**. Puntos extra si das distintos algoritmos de loadbalancing (random, round-robin, etc). Y más puntos extra si ofrecés algún tipo de rate limiting.
  
Obviamente el que desee consagrarse como **The Ultimate Golang Hacker** puede hacer más de 1 proyecto o todos :scream:
