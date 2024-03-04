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
	//database.ExecuteSchemaSQL(db)

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
    SELECT Products.name, Products.product_id, Order_details.order_id, Order_details.count
    FROM Order_details
    JOIN Products ON Order_details.product_id = Products.product_id
    JOIN Stands ON Order_details.product_id = Stands.product_id
    WHERE Order_details.order_id = $1
`, num)
		if err != nil {
			log.Fatalf("Ошибка выполнения запроса к базе данных: %v\n", err)
		}
		defer rows.Close()

		// Обработка результатов запроса
		for rows.Next() {
			var productName string
			var productID, orderID, count int
			if err := rows.Scan(&productName, &productID, &orderID, &count); err != nil {
				log.Fatalf("Ошибка при сканировании строки результата: %v\n", err)
			}
			fmt.Println("productName", productName, "productID", productID, "orderID", orderID, "count", count)
			// Далее обрабатываем данные...
		}

	}
}
