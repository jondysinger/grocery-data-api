package api

import (
	"encoding/json"
	"net/http"

	"github.com/jondysinger/grocery-data/api/pkg/models"
)

// Writes the given data as a json response
func (app *App) writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// Writes an error as a json response with the given status code
func (app *App) errorJson(w http.ResponseWriter, err error, statusCode int) error {
	var payload models.JsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJson(w, statusCode, payload)
}
