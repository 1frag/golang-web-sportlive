package handlers

import (
	"common"
	"html/template"
	"net/http"
)

func HandlerA(w http.ResponseWriter, r *http.Request) {
	// стартовая страничка
	t, _ := template.ParseFiles("pages/case-a.html")
	_ = t.Execute(w, nil)
}

func HandlerB(w http.ResponseWriter, r *http.Request) {
	// Страничка заполнение даты и команд поединка
	t, _ := template.ParseFiles("pages/case-b.html")
	_ = t.Execute(w, nil)
}

func HandlerC(w http.ResponseWriter, r *http.Request) {
	// Страничка выбора ивентов
	t, _ := template.ParseFiles("pages/case-c.html")
	_ = t.Execute(w, nil)
}

func HandlerD(w http.ResponseWriter, r *http.Request) {
	// Страничка заполнения деталей ивента
	t, _ := template.ParseFiles("pages/case-d.html")
	_ = t.Execute(w, nil)
}

func HandlerH(w http.ResponseWriter, r *http.Request) {
	// Страничка просмотра сыгранных игр
	t, _ := template.ParseFiles("pages/case-h.html")
	_ = t.Execute(w, nil)
}

func HandlerG(w http.ResponseWriter, r *http.Request) {
	// Страничка просмотра событий игры
	type Resp struct {
		Id int64
	}
	_ = r.ParseForm()
	id := common.GetInt64(r, "id")
	t, _ := template.ParseFiles("pages/case-g.html")
	_ = t.Execute(w, Resp{Id: id})
}
