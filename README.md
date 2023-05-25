# CRUD básico en Go

Este es un ejemplo de una aplicación básica en Go que implementa un CRUD (Crear, Leer, Actualizar y Eliminar) utilizando una base de datos MySQL y un servidor HTTP.

## Funcionalidades

La aplicación implementa las siguientes funcionalidades:

- **Crear**: Permite agregar un nuevo empleado a la base de datos.

- **Leer**: Muestra una tabla en la página de inicio que lista todos los empleados registrados en la base de datos.

- **Actualizar**: No implementado en este ejemplo, pero se puede agregar para permitir la edición de los datos de un empleado existente.

- **Eliminar**: No implementado en este ejemplo, pero se puede agregar para permitir la eliminación de un empleado existente.

## Requisitos previos

Antes de ejecutar la aplicación, asegúrate de tener los siguientes requisitos previos:

- Go instalado en tu sistema. Puedes descargarlo desde: https://golang.org/dl/

- Una base de datos MySQL instalada y configurada. Asegúrate de tener el nombre de la base de datos, el nombre de usuario y la contraseña disponibles.

## Configuración

Antes de ejecutar la aplicación, debes configurar la conexión a la base de datos. Sigue estos pasos:

1. Abre el archivo `main.go` en un editor de texto.

2. Busca la función `conexionBD` y actualiza los valores de las variables `Usuario`, `Contrasena` y `Nombre` con tu información de conexión a la base de datos.

3. Guarda los cambios en el archivo.

## Ejecución

Sigue estos pasos para ejecutar la aplicación:

1. Abre una terminal y navega hasta el directorio del proyecto.

2. Ejecuta el siguiente comando para compilar el código fuente:

   ```bash
   go build



