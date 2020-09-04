package main

import "fmt"

type Product struct {
	Id int
	name string
	price int
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
		Id: 1,
		name: "衣服",
		price: 80,
	}
	//新增一個
	add(product)
	//健立第二個物品
	product = Product{
		Id: 2,
		name: "鞋子",
		price: 100,
	}
	add(product)

	//更改第二個物品的名字
	product = Product{
		Id: 2,
		name: "白白的鞋子",
		price: 10000,
	}
	//更改
	edit(product)
	//列出所有
	//showAll()
	//showOne(1)
	Del(1)
	showAll()


}
//新增
func add(product Product) {
	products[product.Id] = &product
}
//編輯
func edit(product Product){
	if product.Id != 0 {
		if product.name !="" {
			products[product.Id].name = product.name
		}

		if product.price != 0 {
			products[product.Id].price = product.price
		}

	}
}
//列出所有
func showAll() {
	for _, product := range products {
		fmt.Println(product.name)
	}
}
//列出指定
func showOne(id int)  {
	fmt.Println(products[id].Id)
	fmt.Println(products[id].name)
	fmt.Println(products[id].price)
}
//刪除指定
func Del(id int)  {
	delete(products,id)
}