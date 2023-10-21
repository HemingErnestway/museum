package api

import (
	"encoding/json"
	"fmt"
	"museum/engine"
	"museum/entity"
	"museum/storage"
	"net/http"
)

func (h *Handler) ExhibitCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var exhibit entity.Exhibit
	if err := decoder.Decode(&exhibit); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.ExhibitCreate(exhibit))
}

func (h *Handler) ExhibitRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ExhibitRead(id))
}

func (h *Handler) ExhibitsRead(ctx *engine.Context) {
	ctx.Print(storage.ExhibitsRead())
}

func (h *Handler) ExhibitUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Exhibit
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ExhibitUpdate(newFields, id))
}

func (h *Handler) ExhibitDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ExhibitDelete(id))
}
