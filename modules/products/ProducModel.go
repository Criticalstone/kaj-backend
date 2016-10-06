package products

type Product struct {
  ID        int     `json:"id"`
  Name    string  `json:"name"`
}

func (p *Product) GetName() string {
  return p.Name
}
