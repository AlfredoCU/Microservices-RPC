# Microservices (RPC)

## Objetivo

Implementar un *microservicio* usando la tecnología `RPC` de *Go* para guardar las calificaciones de los alumnos del profe Boites.

## Requerimientos

El servidor deberá de exponer usando `RPC` las siguientes funciones:

- Agregar la calificación de un alumno por materia.
- Obtener el promedio del alumno.
- Obtener el promedio de todos los alumnos.
- Obtener el promedio por materia.
- Retornar un mensaje de error cuando ya exista la calificación de una materia para un alumno en específico.

La información deberá de ser almacenada en estructuras de datos (slice, map, listas, etc) sin usar un gestor de base de datos. Así como construir las estructuras necesarias.

El cliente deberá de contar con un menú con opciones para:

- **Agregar calificación de una materia**: Pedir el nombre del alumno, materia y su correspondiente calificación, para posteriormente invocar por `RPC` la función para *Agregar la calificación de un alumno por materia.*
- **Mostrar el promedio de un Alumno**: Pedir nombre del alumno, invocar por `RPC` la función *Obtener el promedio del alumno* e imprimir el promedio del alumno.
- **Mostrar el promedio general**: Invocar por `RPC` la función *Obtener el promedio de todos los alumnos* e imprimir el promedio general.
- **Mostrar el promedio de una materia** (pedir materia): Pedir el nombre de la materia,  invocar por `RPC` la función para *Obtener el promedio por materia* e imprimir el promedio de la materia.
