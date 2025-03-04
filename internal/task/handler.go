package task

import (
	"context"
	"fmt"

	"github.com/mrzhov/course-app/internal/web/tasks"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetList()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:          &tsk.ID,
			Title:       tsk.Title,
			Description: &tsk.Description,
			Completed:   &tsk.Completed,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	body := request.Body

	// if body.Title == nil {
	// 	err := echo.NewHTTPError(http.StatusBadRequest, "Field title is required")
	// 	return nil, err
	// }

	fmt.Println(*body)

	// task := Task{
	// 	Title:       *body.Title,
	// 	Description: *body.Description,
	// 	Completed:   *body.Completed,
	// }

	// err := h.service.Create(&task)
	// if err != nil {
	// 	return nil, err
	// }

	// response := tasks.PostTasks201JSONResponse{
	// 	Id:          &task.ID,
	// 	Title:       &task.Title,
	// 	Description: &task.Description,
	// 	Completed:   &task.Completed,
	// }

	response := tasks.PostTasks201JSONResponse{}

	return response, nil
}

// type CreateBody struct {
// 	Title  *string `json:"title"`
// 	IsDone bool    `json:"is_done"`
// }

// func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
// 	var body CreateBody
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if body.Title == nil {
// 		err := errors.New("error: field \"title\" is required")
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var task Task
// 	task.Title = *body.Title
// 	task.IsDone = body.IsDone
// 	err := h.service.Create(&task)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(task)
// }

// func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
// 	id, idErr := strconv.Atoi(mux.Vars(r)["id"])

// 	if idErr != nil {
// 		http.Error(w, idErr.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	task, taskErr := h.service.GetById(uint(id))
// 	if taskErr != nil {
// 		http.Error(w, taskErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(task)
// }

// type PatchBody struct {
// 	Title  *string `json:"title"`
// 	IsDone *bool   `json:"is_done"`
// }

// func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
// 	id, idErr := strconv.Atoi(mux.Vars(r)["id"])
// 	if idErr != nil {
// 		http.Error(w, idErr.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var body PatchBody
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if body.Title == nil && body.IsDone == nil {
// 		err := errors.New("error: at least one field must be specified")
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	task, taskErr := h.service.GetById(uint(id))
// 	if taskErr != nil {
// 		http.Error(w, taskErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if body.Title != nil {
// 		task.Title = *body.Title
// 	}
// 	if body.IsDone != nil {
// 		task.IsDone = *body.IsDone
// 	}

// 	patchErr := h.service.Patch(&task)

// 	if patchErr != nil {
// 		http.Error(w, patchErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(task)
// }

// type ResponseMessage struct {
// 	Message string `json:"message"`
// }

// func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
// 	id, idErr := strconv.Atoi(mux.Vars(r)["id"])

// 	if idErr != nil {
// 		http.Error(w, idErr.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	task, taskErr := h.service.GetById(uint(id))
// 	if taskErr != nil {
// 		http.Error(w, taskErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	deleteErr := h.service.Delete(&task)

// 	if deleteErr != nil {
// 		http.Error(w, deleteErr.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(ResponseMessage{Message: "Deleted successfully!"})
// }
