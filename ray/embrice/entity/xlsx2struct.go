package entity

type Goods struct {
	Name      string  `json:"name"`      // 商品名称
	Price     float64 `json:"price"`     // 商品价格
	Inventory int     `json:"inventory"` // 商品库存
}

type GoodsList struct {
	List []Goods `json:"list"` // 商品列表
}

func (this *Goods) GetName() string {
	return this.Name
}

func (this *Goods) GetPrice() float64 {
	return this.Price
}

func (this *Goods) GetInventory() int {
	return this.Inventory
}
