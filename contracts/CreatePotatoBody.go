package contracts

// CreatePotatoBody represents the input structure for creating a potato
type CreatePotatoBody struct {
	Img        string  `json:"img"`
	Price      uint    `json:"price"`
	Title      string  `json:"title"`
	Rate       float32 `json:"rating"`
	Types      []uint  `json:"types"`
	Sizes      []uint  `json:"sizes"`
	Categories []uint  `json:"categories"`
}
