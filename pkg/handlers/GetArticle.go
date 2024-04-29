package handlers

import (
	"encoding/json"
	"example/go-database-example/pkg/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	queryStmt := `SELECT * FROM articles WHERE id = $1;`
	results, err := h.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("failed to execute query", err)
		return
	}

	var article models.Article
	for results.Next() {
		err = results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)

}