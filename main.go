package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"test-task/database"
)

func main() {
	//_ = database.InitDB()
	db := database.InitDB()
	database.ExecuteSchemaSQL(db)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Нет аргументов командной строки.")
		return
	}
	arg := args[1]
	values := strings.Split(arg, ",")
	//var arguments []int
	for _, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Ошибка преобразования аргумента %s в число: %v\n", value, err)
			continue
		}
		//arguments = append(arguments, num)
		fmt.Println("Прочитанное значение:", num)
		rows, err := db.Query(`
    SELECT Stands.stand_id, Stand_name.name, Products.name, Products.product_id, Order_details.order_id, Order_details.count, Stands.parent_id
    FROM Order_details
    JOIN Products ON Order_details.product_id = Products.product_id
    JOIN Stands ON Order_details.product_id = Stands.product_id
	JOIN Stand_name ON Stands.stand_id = Stand_name.stand_id
    WHERE Order_details.order_id = $1
`, num)
		if err != nil {
			log.Fatalf("Ошибка выполнения запроса к базе данных: %v\n", err)
		}
		defer rows.Close()

		// Обработка результатов запроса
		for rows.Next() {
			var productName string
			var standsID, productID, orderID, count, parentId int
			var standsName string
			if err := rows.Scan(&standsID, &standsName, &productName, &productID, &orderID, &count, &parentId); err != nil {
				log.Fatalf("Ошибка при сканировании строки результата: %v\n", err)
			}
			fmt.Println("standsID", standsID, "standsName", standsName, "productName", productName, "productID", productID, "orderID", orderID, "count", count, "parentId", parentId)

		}

	}
}
