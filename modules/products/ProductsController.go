package products

import (
  "github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
  "strconv"
  "net/http"
)

// GET: /products/
func Index(w http.ResponseWriter, r *http.Request) {
  products, _ := selectFromDB("select * from products")

  for _,p := range products {
    printJSON(p, w)
  }
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

  product, err := selectFromDB("select * from products where id=" + strconv.Itoa(id))

  if (err != nil) {
    http.Error(w, "No product with that id", http.StatusBadRequest)
    return
  }

  printJSON(product[0], w)
}
