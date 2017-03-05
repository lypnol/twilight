package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func startWebApp(m *Map) {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	index, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index.Execute(w, packMap(m))
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(packMap(m))
	})

	log.Println("Web server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type packed struct {
	X, Y   int
	Humans []cell
	Vamps  []cell
	Wolfs  []cell
}

func packMap(m *Map) packed {
	p := packed{X: m.Columns*80 + 1, Y: m.Rows*80 + 1}
	for _, i := range m.humans {
		p.Humans = append(p.Humans, scale(m.cells[i]))
	}
	for _, i := range m.monster[wolf-1] {
		p.Wolfs = append(p.Wolfs, scale(m.cells[i]))
	}
	for _, i := range m.monster[vamp-1] {
		p.Vamps = append(p.Vamps, scale(m.cells[i]))
	}
	return p
}

func scale(c cell) cell {
	c.X = c.X * 80
	c.Y = c.Y * 80
	return c
}