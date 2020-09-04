package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id int
	Name string
	Price int
}

var (
	products map[int]*Product
)

func main() {
	//建立Product庫
	products = make(map[int]*Product)
	//建立CRUD操作物件
	//建立第一個物品
	product := Product{
		Name: "衣服",
		Price: 80,
	}
	//新增一個
	fmt.Println("ADD-->",add(product))
	//健立第二個物品
	product = Product{
		Name: "鞋子",
		Price: 100,
	}
	fmt.Println("ADD-->",add(product))

	//更改第二個物品的名字
	product = Product{
		Name: "白白的鞋子",
		Price: 10000,
	}
	//更改
	fmt.Println("EDIT ID 2 -->",edit(2,product))
	//列出所有
	fmt.Println("SHOW ALL-->",showAll())
	//取得一個
	fmt.Println("SHOW ID 1-->",showOne(1))
	//刪除一個
	fmt.Println("DELETE ID 1--->",Del(1))

}
//新增
func add(product Product) string {
	product.Id = len(products) + 1
	products[product.Id] = &product
	p,_ :=json.Marshal(product)
	return string(p)
}
//編輯
func edit(id int,product Product) string {
	if id != 0 {
		if product.Name !="" {
			products[id].Name = product.Name
		}

		if product.Price != 0 {
			products[id].Price = product.Price
		}
	}
	p,_ :=json.Marshal(products[id])
	return string(p)
}
//列出所有
func showAll() string {
	p,_ :=json.Marshal(products)
	return string(p)
}
//列出指定
func showOne(id int) string {
	p,_ :=json.Marshal(products[id])
	return string(p)
}
//刪除指定
func Del(id int) string {
	delete(products,id)
	p,_ :=json.Marshal(products)
	return string(p)
}