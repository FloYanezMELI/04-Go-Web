
Bootcamp Go
# Procesando peticiones
Go Web

---

### Objetivo
El objetivo de esta guía práctica es que podamos afianzar los conceptos sobre condicionales, vistos en el módulo de Go Web. Para esto vamos a plantear una serie de ejercicios simples e incrementales (trabajaremos y agregaremos complejidad a lo que tenemos que construir), lo que nos permitirá repasar los temas que estudiamos. 

### Forma de trabajo
Los ejercicios deben ser realizados en sus computadoras. Les recordamos crear una carpeta para cada clase y que ahí dentro tengan un archivo .go para cada ejercicio.


¿Are you ready? 😎👍

---

### 💡 Ejercicio 1 - Prueba de Ping
Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
1. El endpoint deberá ser de método GET
2. La respuesta de “pong” deberá ser enviada como texto, NO como JSON



### 💡 Ejercicio 2 - Manipulando el body

Vamos a crear un endpoint llamado **/saludo**. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hola + nombre + apellido”

1. El endpoint deberá ser de método POST
2. Se deberá usar el package JSON para resolver el ejercicio
3. La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
4. La estructura deberá ser como esta:
```json
{
	“nombre”: “Andrea”,
	“apellido”: “Rivas”
}
```
