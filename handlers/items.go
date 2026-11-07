package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KaustubhBuilds/acetlisto/stores"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
)

type ItemHandler struct {
	store stores.ItemStore
}

func NewItemHandler(store stores.ItemStore) *ItemHandler {
	return &ItemHandler{store: store}
}

func (i *ItemHandler) listItems(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("handlers")
	_, span := tracer.Start(r.Context(), "list-items")
	defer span.End()

	items, err := i.store.GetAllItems()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not load items")
		return
	}
	respondWithJSON(w, http.StatusOK, items)
}

func (i *ItemHandler) createItem(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("handlers")
	_, span := tracer.Start(r.Context(), "create-item")
	defer span.End()

	var request stores.CreateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer deferredClose(r.Body)

	item, err := i.store.CreateItem(request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not store item")
		return
	}
	respondWithJSON(w, http.StatusCreated, item)
}

func (i *ItemHandler) getItem(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("handlers")
	_, span := tracer.Start(r.Context(), "get-item")
	defer span.End()

	vars := mux.Vars(r)
	item, err := i.store.GetItem(vars["ID"])
	if err != nil {
		respondWithError(w, http.StatusNotFound, "")
		return
	}
	respondWithJSON(w, http.StatusOK, item)
}

func (i *ItemHandler) updateItem(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("handlers")
	_, span := tracer.Start(r.Context(), "update-item")
	defer span.End()

	var request stores.CreateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer deferredClose(r.Body)

	vars := mux.Vars(r)
	item, err := i.store.UpdateItem(vars["ID"], request)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "")
		return
	}
	respondWithJSON(w, http.StatusOK, item)
}

func (i *ItemHandler) deleteItem(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("handlers")
	_, span := tracer.Start(r.Context(), "delete-item")
	defer span.End()

	vars := mux.Vars(r)
	if err := i.store.DeleteItem(vars["ID"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, "")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{})
}
