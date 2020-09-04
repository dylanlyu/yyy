package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id    int
	Name  string
	Price int
}

var (
	products map[int]*Product
)

func main() {
	//建立Product庫
	products = make(map[int]*Product)
	//建立第一個物品
	product := Product{
		Name:  "衣服",
		Price: 80,
	}
	//新增一個
	fmt.Println("ADD-->", add(product))
	//健立第二個物品
	product = Product{
		Name:  "鞋子",
		Price: 100,
	}
	fmt.Println("ADD-->", add(product))

	//更改第二個物品的名字
	product = Product{
		Name:  "白白的鞋子",
		Price: 10000,
	}
	//更改
	fmt.Println("EDIT ID 2 -->", edit(2, product))
	//列出所有
	fmt.Println("SHOW ALL-->", showAll())
	//取得一個
	fmt.Println("SHOW ID 1-->", showOne(1))
	//刪除一個
	fmt.Println("DELETE ID 1--->", Del(1))

}

//新增
func add(product Product) string {
	// 先確定現在陣列裡面有多少個物件,然後+1 AUTO ID
	product.Id = len(products) + 1
	// 將處理好的ID放進陣列做ID及obj的ID
	products[product.Id] = &product
	// 序列化,輸入進來的資料
	p, _ := json.Marshal(product)
	// 回傳輸入進來的資料
	return string(p)
}

//編輯
func edit(id int, product Product) string {
	// 判斷輸入進來name是否為空
	if product.Name != "" {
		// 把輸入ID的Name,改成輸入進來Name
		products[id].Name = product.Name
	}
	// 判斷輸入進來Price是否為0
	if product.Price != 0 {
		// 把輸入ID的Price,改成輸入進來Price
		products[id].Price = product.Price
	}
	// 序列化,product庫的資料
	p, _ := json.Marshal(products[id])
	// 回傳product庫的資料
	return string(p)
}

//列出所有
func showAll() string {
	// 序列化,product庫所有資料
	p, _ := json.Marshal(products)
	// 回傳product庫所有資料
	return string(p)
}

//列出指定
func showOne(id int) string {
	// 序列化,product庫指定的資料
	p, _ := json.Marshal(products[id])
	// 回傳product庫指定的資料
	return string(p)
}

//刪除指定
func Del(id int) string {
	// 刪除product庫指定的資料
	delete(products, id)
	// 序列化,product庫所有資料
	p, _ := json.Marshal(products)
	// 回傳product庫所有資料
	return string(p)
}
