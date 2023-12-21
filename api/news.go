package api

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	//news := NewsContent{
	//	News: storage.NewsRead(),
	//}

	news := storage.NewsRead()

	ctx.Response.Header().Set("Content-Type", "text/html")
	//dataMarshal, _ := json.Marshal(news.News)

	t := template.Must(template.New("templ").Parse(templ))
	//var m []map[string]interface{}
	//
	//if err := json.Unmarshal([]byte(dataMarshal), &m); err != nil {
	//	panic(err)
	//}

	if err := t.Execute(ctx.Response, news); err != nil {
		panic(err)
	}
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

type NewsContent struct {
	News []entity.News
}

const templ = `
{{ range . }}
<div class="row">
    <div class="col-md-4">
        <img src="{{ .ImgPath }}" alt="" class="info-image img-fluid">
    </div>
    <div class="col-md-8">
        <h3>{{ .Header }}</h3>
        <p class="datetime">{{ .DateTime }}</p>
        <p>{{ .Content }}</p>
    </div>
</div>
{{ end }}
`
