package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Task representa una tarea en la lista de tareas
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Lista de tareas simuladas para este ejemplo
var Tasks []Task

// getTasksHandler maneja las solicitudes para obtener todas las tareas
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tasks)
}

// createTaskHandler maneja las solicitudes para crear una nueva tarea
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error al decodificar la tarea", http.StatusBadRequest)
		return
	}

	// Simulamos la generación del ID
	task.ID = len(Tasks) + 1

	Tasks = append(Tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// markTaskCompletedHandler maneja las solicitudes de tareas completadas
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
	var foundTask Task
	for i, t := range Tasks {
		if t.ID == id {
			foundTask = Tasks[i]
			foundTask.Completed = true
			Tasks[i] = foundTask
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
