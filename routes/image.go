package routes

import (
	"encoding/json"
	"gtp/databases"
	"gtp/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func HandleProductImages(w http.ResponseWriter, r *http.Request) {
	db := databases.DbConn()

    productId := strings.TrimPrefix(r.URL.Path, "/images/")

	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var imageData models.Image
		err = json.Unmarshal(body, &imageData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		db = databases.DbConn()
		insForm, err := db.Prepare("INSERT INTO Images(product_id, image_url, description) VALUES(?,?,?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		_, err = insForm.Exec(productId, imageData.ImageUrl, imageData.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		http.Redirect(w, r, "/images/" + productId, http.StatusSeeOther)

	case http.MethodGet:
		db := databases.DbConn()
		result, err := db.Query("SELECT id, image_url, description FROM Images WHERE product_id=?", productId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		res := []models.Image{}
		for result.Next() {
			var id int
			var image_url, description string
			err = result.Scan(&id, &image_url, &description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			image := models.Image{
				Id: id,
				ImageUrl: image_url,
				Description: description,
			}
			res = append(res, image)
		}

		jsonRes, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	default:
		http.Error(w, "Wrong method", http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/products", http.StatusSeeOther)
}