package handlers

import (
	"encoding/json"
	"example/go-database-example/pkg/models"
	"log"
	"net/http"
)

func (h handler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	results, err := h.DB.Query("SELECT * FROM articles;")
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var articles = make([]models.Article, 0)
	for results.Next() {
		var article models.Article
		err := results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			log.Println("failed to scan", err)
			return
		}

		articles = append(articles, article)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(articles)
}