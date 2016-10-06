package products

import (
  "strconv"
  "net/http"
  "encoding/json"
  "errors"
)

func getDummyData() []Product {
  data := make([]Product, 10)
  for i := 0; i < 10; i++ {
    data[i] = Product{ID: i + 1, Name: "Random product " + strconv.Itoa(i + 1)}
  }
  return data
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
