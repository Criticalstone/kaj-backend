package products

import (
  "github.com/gorilla/mux"

  "strconv"
  "net/http"
)

// GET: /products/
func Index(w http.ResponseWriter, r *http.Request) {
  printJSON(getDummyData(), w)
}

// GET: /products/1
func Show(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])

  // This shouldn't happen since the mux only accepts numbers to this route
  if (err != nil) {
    http.Error(w, "Invalid Id, please try again.", http.StatusBadRequest)
    return
  }

  products := getDummyData()
  product, err := getProductsByID(products, id)

  if (err != nil) {
    http.NotFound(w, r)
    return
  }

  printJSON(product, w)
}
