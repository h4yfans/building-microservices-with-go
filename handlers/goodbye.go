package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Good Bye World"))
	//g.l.Println("Good Bye World")
	//
	//d, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(rw, "Ooops", http.StatusBadRequest)
	//	return
	//}
	//
	//fmt.Fprintf(rw, "Hello %s", d)
}
