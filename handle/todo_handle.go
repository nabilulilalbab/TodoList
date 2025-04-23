package handle

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/nabilulilallbab/TodoList/config"
	"github.com/nabilulilallbab/TodoList/entity"
	"github.com/nabilulilallbab/TodoList/library"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.ConnectDatabase()

func Index(w http.ResponseWriter, r *http.Request) error {
	var todos []entity.Todo
	if err := DB.Order("id desc").Find(&todos).Error; err != nil {
		log.Println("Error fetching todos:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	err := library.NewTemplate().ExecuteTemplate(w, "index.html", map[string]any{
		"Title": "Home Todo",
		"Todos": todos,
	})
	if err != nil {
		log.Println("Error rendering index template:", err)
	}
	return err
}

func CreateForm(w http.ResponseWriter, r *http.Request) error {
	log.Println("Rendering create form...")
	err := library.NewTemplate().ExecuteTemplate(w, "create.html", map[string]any{
		"Title": "Create form",
	})
	if err != nil {
		log.Println("Error rendering create form:", err)
	}
	return err
}

func CreateTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return errors.New("method not allowed")
	}

	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return nil
	}

	todo := entity.Todo{
		Title:  r.FormValue("title"),
		IsDone: false,
	}

	if err := DB.Create(&todo).Error; err != nil {
		log.Println("Error creating todo:", err)
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return nil
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func ToggleTodo(w http.ResponseWriter, r *http.Request) error {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid ID:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil
	}

	var todo entity.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		log.Println("Todo not found:", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return nil
	}

	todo.IsDone = !todo.IsDone
	if err := DB.Save(&todo).Error; err != nil {
		log.Println("Failed to update todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) error {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid ID:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil
	}

	var todo entity.Todo
	if err := DB.Delete(&todo, id).Error; err != nil {
		log.Println("Failed to delete todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func EditForm(w http.ResponseWriter, r *http.Request) error {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid ID:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil
	}

	var todo entity.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		log.Println("Todo not found:", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return nil
	}

	err = library.NewTemplate().ExecuteTemplate(w, "edit.html", map[string]any{
		"Title": "Edit Todo",
		"Todo":  todo,
	})
	if err != nil {
		log.Println("Error rendering edit form:", err)
	}
	return err
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return errors.New("method not allowed")
	}

	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return nil
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid ID:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil
	}

	var todo entity.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		log.Println("Todo not found:", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return nil
	}

	todo.Title = r.FormValue("title")
	if err := DB.Save(&todo).Error; err != nil {
		log.Println("Failed to update todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
