package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tutorial-go/models" // Asegúrate de reemplazar "nombre_proyecto" por el nombre real de tu módulo
)

// Lista de tareas simuladas para este ejemplo
var tasks []models.Task

// getTaskHandler maneja las solicitudes para obtener todas las tareas
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// createTaskHandler maneja las solicitudes para crear una nueva tarea
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error al decodificar la tarea", http.StatusBadRequest)
		return
	}

	// Simulamos la generación del ID
	task.ID = len(tasks) + 1

	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// markTaskCompletedHandler maneja las solicitudes para marcar una tarea como completada
func MarkTaskCompletedHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenemos el ID de la tarea de la ruta
	// Ejemplo de ruta: /tasks/3 (para marcar la tarea con ID 3 como completada)
	// Aquí extraemos el "3" de la ruta
	// Si el ID no es un número válido, responderemos con un error
	var id int
	_, err := fmt.Sscanf(r.URL.Path, "/tasks/%d", &id)
	if err != nil {
		http.Error(w, "ID de tarea inválido", http.StatusBadRequest)
		return
	}

	// Buscamos la tarea con el ID correspondiente en la lista de tareas
	// Si no encontramos la tarea, responderemos con un error
	var foundTask models.Task
	for i, t := range tasks {
		if t.ID == id {
			foundTask = tasks[i]
			foundTask.Completed = true
			tasks[i] = foundTask
			break
		}
	}

	if foundTask.ID == 0 {
		http.Error(w, "Tarea no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundTask)
}
