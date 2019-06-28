package backend

import (
	"encoding/json"
	"net/http"
)

//Controller ...
type Controller struct {
	Repository Repository
}

// GetNewJoke GET /
func (c *Controller) GetNewJoke(w http.ResponseWriter, r *http.Request) {
	joke := c.Repository.GetJoke()
	data, _ := json.Marshal(joke)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
