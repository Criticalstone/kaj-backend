package products

import (
  "strconv"
  "net/http"
  "encoding/json"
  "errors"
  "fmt"
  "database/sql"
)

func getDummyData() []Product {
  data := make([]Product, 10)
  for i := 0; i < 10; i++ {
    data[i] = Product{ID: i + 1, Name: "Random product " + strconv.Itoa(i + 1)}
  }
  return data
}

func selectFromDB(query string) []Product {
  db, err := sql.Open("mysql",
    "root:kakoregoda@tcp(127.0.0.1:3306)/products")
  if (err != nil) {
    fmt.Println(err)
  }

  var product Product
  var products []Product
  rows, err := db.Query(query)
  if err != nil {
  	fmt.Println(err)
  }

  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&product.ID, &product.Name)
    if err != nil {
      fmt.Println(err)
    }
    products = append(products, product)
  }
  err = rows.Err()
  if err != nil {
    fmt.Println(err)
  }

  defer db.Close()
  return products
}

func printJSON(product interface{}, w http.ResponseWriter) {
  response, err := json.MarshalIndent(product, "", "\t")
  if (err != nil) {
    http.Error(w, "Something went wrong, try again.", http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func getProductsByID(products []Product, id int) (Product, error) {
  var product Product
  found := false
  // Super simple but crap search
  for _,p := range products {
    if (p.ID == id) {
      product = p
      found = true
      break
    }
  }

  if (!found) {
    return Product{}, errors.New("Product not found")
  }
  return product, nil
}
