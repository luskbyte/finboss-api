package handlers

import (
	"net/http"
	"strconv"
)

func pathID(r *http.Request) (uint, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
