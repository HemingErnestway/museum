package api

import (
	"museum/engine"
	"strconv"
	"strings"
)

type Handler struct {
}

func GetIdFromContext(ctx *engine.Context) uint32 {
	path := strings.Split(ctx.Request.URL.Path, "/")
	idString := path[len(path)-1]
	idUint32, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	return uint32(idUint32)
}
