package handler

import "net/http"

type Home struct {
}

func (*Home) PostUrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("wadaw"))
}

func HomeH() *Home {
	return &Home{}
}
