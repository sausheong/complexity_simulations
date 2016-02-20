package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	files := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", files))

	http.HandleFunc("/volunteers/show", volunteerShow)
	http.HandleFunc("/volunteers", volunteers)

	http.HandleFunc("/segregation/show", segregationShow)
	http.HandleFunc("/segregation/start", segregationStart)
	http.HandleFunc("/segregation", segregation)

	http.HandleFunc("/culture/show", cultureShow)
	http.HandleFunc("/culture/start", cultureStart)
	http.HandleFunc("/culture", culture)
	http.HandleFunc("/culture/stats/show", cultureStatsShow)
	http.HandleFunc("/culture/stats", cultureStats)
	http.ListenAndServe(":8080", nil)
}

// Volunteer's Dilemma
func volunteerShow(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/volunteers.html")
	t.Execute(w, nil)
}

// Emit JSON for Volunteer's Dilemma
func volunteers(w http.ResponseWriter, r *http.Request) {
	cost, _ := strconv.ParseFloat(r.FormValue("cost"), 64)
	overall, _ := strconv.ParseFloat(r.FormValue("overall"), 64)
	maxn, _ := strconv.Atoi(r.FormValue("maxn"))
	output, _ := json.Marshal(volunteerExp(cost, overall, maxn))
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// Racial Segregation
func segregationShow(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/segregation.html")
	t.Execute(w, WIDTH)
}

// Initialize the grid
func segregationStart(w http.ResponseWriter, r *http.Request) {
	neighbourliness, _ := strconv.Atoi(r.FormValue("n"))
	races, _ := strconv.Atoi(r.FormValue("races"))
	vacancy, _ := strconv.Atoi(r.FormValue("vacancy"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	populate(neighbourliness, races, vacancy, limit)
	output, _ := json.Marshal(grid)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// show grid after each tick; emit JSON for segregation
func segregation(w http.ResponseWriter, r *http.Request) {
	tick()
	output, _ := json.Marshal(grid)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// Dissemination of Culture
func cultureShow(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/culture.html")
	t.Execute(w, WIDTH)
}

// Initialize the cultureGrid
func cultureStart(w http.ResponseWriter, r *http.Request) {
	culturePopulate()
	output, _ := json.Marshal(cultureGrid)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// Emit JSON data for culture
func culture(w http.ResponseWriter, r *http.Request) {
	cultureTick()
	output, _ := json.Marshal(cultureGrid)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// Show graph page for culture
func cultureStatsShow(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/culture_stats.html")
	t.Execute(w, nil)
}

// Emit JSON for culture graph
func cultureStats(w http.ResponseWriter, r *http.Request) {
	count := 0
	stats := make([][]int, len(fdistances))
	for n, _ := range fdistances {
		stats[count] = []int{count, fdistances[n], changes[n], uniques[n]}
		count++
	}

	output, _ := json.Marshal(stats)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
