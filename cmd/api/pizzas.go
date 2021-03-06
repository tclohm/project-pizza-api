package main

import (
	"fmt"
	"net/http"
	"strconv"
	"errors"

	"github.com/tclohm/project-pizza/internal/data"
	"github.com/tclohm/project-pizza/internal/validator"

	"github.com/gorilla/mux"
)

func (app *application) createPizzaHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name 				string 	`json:"name"`
		Style 				string 	`json:"style"`
		Price 				float32 `json:"price"`
		Description 		string 	`json:"description"`
		Cheesiness 			float32 `json:"cheesiness"`
		Flavor 				float32 `json:"flavor"`
		Sauciness 			float32 `json:"sauciness"`
		Saltiness 			float32 `json:"saltiness"`
		Charness 			float32 `json:"charness"`
		ImageId				int64 	`json:"image_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pizza := &data.Pizza{
		Name: 				input.Name,
		Style: 				input.Style,
		Price: 				input.Price,
		Description: 		input.Description,
		Cheesiness: 		input.Cheesiness,
		Flavor: 			input.Flavor,
		Sauciness: 			input.Sauciness,
		Saltiness: 			input.Saltiness,
		Charness: 			input.Charness,
		ImageId:			input.ImageId,
	}

	v := validator.New()

	if data.ValidatePizza(v, pizza); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Pizzas.Insert(pizza)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/pizzas/%d", pizza.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"pizza": pizza}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showPizzaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	pizza, err := app.models.Pizzas.Get(n)

	if err != nil {
		switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.notFoundResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
		}
		return
	}
	

	err = app.writeJSON(w, http.StatusOK, envelope{"pizza": pizza}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePizzaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	pizza, err := app.models.Pizzas.Get(n)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name 				*string 	`json:"name"`
		Style 				*string 	`json:"style"`
		Price  				*float32 	`json:"price"`
		Description 		*string 	`json:"description"`
		Cheesiness 			*float32 	`json:"cheesiness"`
		Flavor 				*float32 	`json:"flavor"`
		Sauciness 			*float32 	`json:"sauciness"`
		Saltiness 			*float32 	`json:"saltiness"`
		Charness 			*float32 	`json:"charness"`
		ImageId				*int64 	 	`json:image_id`

	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		pizza.Name = *input.Name
	}

	if input.Style != nil {
		pizza.Style = *input.Style
	}

	if input.Price != nil {
		pizza.Price = *input.Price
	}

	if input.Description != nil {
		pizza.Description = *input.Description
	}

	if input.Cheesiness != nil {
		pizza.Cheesiness = *input.Cheesiness
	} 

	if input.Flavor != nil {
		pizza.Flavor = *input.Flavor
	} 	

	if input.Sauciness != nil {
		pizza.Sauciness = *input.Sauciness
	}

	if input.Saltiness != nil {
		pizza.Saltiness = *input.Saltiness
	}

	if input.Charness != nil {
		pizza.Charness = *input.Charness
	} 

	if input.ImageId != nil {
		pizza.ImageId = *input.ImageId
	}	


	v := validator.New()

	if data.ValidatePizza(v, pizza); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Pizzas.Update(pizza)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pizza": pizza}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}	
}

func (app *application) deletePizzaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Pizzas.Delete(n)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "pizza successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPizzasHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name 		string
		Style 		string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Style = app.readString(qs, "style", "")

	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "name", "style", "-id", "-name", "-style"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pizzas, metadata, err := app.models.Pizzas.GetAll(input.Name, input.Style, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pizzas": pizzas, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}