package entity

type Goods struct {
	Name  string  `json:"name"`  // 商品名称
	Price float64 `json:"price"` // 商品价格
}

type GoodsList struct {
	List []Goods `json:"list"` // 商品列表
}
