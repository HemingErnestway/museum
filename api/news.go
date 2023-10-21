package api

import (
	"encoding/json"
	"fmt"
	"museum/engine"
	"museum/entity"
	"museum/storage"
	"net/http"
)

func (h *Handler) NewsCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var news entity.News
	if err := decoder.Decode(&news); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.NewsCreate(news))
}

func (h *Handler) NewsReadSingle(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.NewsReadSingle(id))
}

func (h *Handler) NewsRead(ctx *engine.Context) {
	ctx.Print(storage.NewsRead())
}

func (h *Handler) NewsUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.News
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.NewsUpdate(newFields, id))
}

func (h *Handler) NewsDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.NewsDelete(id))
}
