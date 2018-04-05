package search

import (
	"github.com/sahilm/fuzzy"
	"ray/embrice/entity"
)

// FuzzySearch 模糊查询
//func FuzzySearch(data []string, pattern string) []*entity.Goods {
func FuzzySearch(goods []*entity.Goods, pattern string) []*entity.Goods {
	data := []string{}
	for k, v := range goods {
		data = append(data, v.Name)
	}

	matches := fuzzy.Find(pattern, data)

	newGoods := make([]*entity.Goods, 0)
	for _, match := range matches {
		newGoods.Name = append(newGoods.Name, match.Str)
	}

	return newGoods
}
