package main

import (
	"fmt"
	"log"
	"net/http"
	"tutorial-go/handlers"
)

func main() {
	// Inicializamos algunas tareas para este ejemplo
	handlers.Tasks = append(handlers.Tasks, handlers.Task{ID: 1, Title: "Completar el proyecto Go", Completed: false})
	handlers.Tasks = append(handlers.Tasks, handlers.Task{ID: 2, Title: "Aprender más sobre Go", Completed: false})

	// Configuramos los manejadores para las rutas
	http.HandleFunc("/tasks", handlers.GetTasksHandler)
	http.HandleFunc("/tasks/create", handlers.CreateTaskHandler)
	http.HandleFunc("/tasks/", handlers.MarkTaskCompletedHandler)

	// Iniciamos el servidor en el puerto 8080
	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
