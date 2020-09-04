package main

import "fmt"

type Product struct {
	Id int
	name string
	price int
}

func main() {
	//建立Product庫
	var products []Product
	//建立CRUD操作物件
	var product Product
	//建立第一個物品
	product = Product{
		Id: 1,
		name: "衣服",
		price: 80,
	}
	//新增一個
	products = add(product,products)
	//健立第二個物品
	product = Product{
		Id: 2,
		name: "鞋子",
		price: 100,
	}
	products = add(product,products)
	fmt.Println(products[1].name)
	//更改第二個物品的名字
	product = Product{
		Id: 2,
		name: "白白的鞋子",
		price: 10000,
	}
	//更改
	products = edit(product,products)
	fmt.Println(products[1].name)

}

func add(product Product,products []Product) []Product {
	products = append(products,product)
	return products
}

func edit(product Product,products []Product) []Product {
	arrId := 0
	for i := 0; i < len(products); i++ {
		if product.Id == products[i].Id {
			arrId = i
			continue
		}
	}

	if arrId != 0 {
		if product.name != "" {
			products[arrId].name = product.name
		}

		if product.price != 0 {
			products[arrId].price = product.price
		}
	}
	return products
}