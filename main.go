package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"test-task/database"
)

type Answer struct {
	orderID     int
	productName string
	productID   int
	count       int
	childs      []string
}

type key struct {
	id   int
	name string
}

func main() {

	db := database.InitDB()
	//database.ExecuteSchemaSQL(db)

	args := os.Args

	var arguments []int
	resulst := make(map[key][]Answer)

	if len(args) < 2 {
		fmt.Println("Нет аргументов командной строки.")
		return
	}
	arg := args[1]
	values := strings.Split(arg, ",")

	for _, value := range values {
		num, err := strconv.Atoi(value)
		//Ошибка преобразования аргумента в число (вместо номера заказа какая-то бурда)
		if err != nil {
			continue
		}

		rows, err := db.Query(`
			SELECT Stands.stand_id, 
			Stand_name.name AS stand_name, 
			Products.name AS product_name, 
			Products.product_id, 
			Order_details.order_id, 
			Order_details.count
			FROM Order_details
			JOIN Products ON Order_details.product_id = Products.product_id
			JOIN Stands ON Order_details.product_id = Stands.product_id
			JOIN Stand_name ON Stands.stand_id = Stand_name.stand_id
			WHERE Order_details.order_id = $1
			AND Stands.parent = TRUE;
		`, num)

		if err != nil {
			log.Fatalf("Ошибка выполнения запроса к базе данных: %v\n", err)
		}

		// Обработка случая, когда номер заказа не найден
		if !rows.Next() {
			continue
		}

		arguments = append(arguments, num)

		defer rows.Close()

		// Обработка результатов запроса
		for rows.Next() {
			var productName string
			var standsID, productID, orderID, count int
			var standsName string

			if err := rows.Scan(&standsID, &standsName, &productName, &productID, &orderID, &count); err != nil {
				log.Fatalf("Ошибка при сканировании строки результата: %v\n", err)
			}

			keys := key{standsID, standsName}

			rows2, err := db.Query(`
				SELECT Stands.stand_id, Stand_name.name
				FROM Stands 
				JOIN Stand_name ON Stands.stand_id = Stand_name.stand_id 
				WHERE product_id = $1 AND parent = FALSE`, productID)

			if err != nil {
				log.Fatal(err)
			}
			defer rows2.Close()

			var childs []string

			for rows2.Next() {
				var child string
				if err := rows2.Scan(&standsID, &child); err != nil {
					log.Fatalf("Ошибка при сканировании строки результата: %v\n", err)
				}
				childs = append(childs, child)
			}

			answ := Answer{orderID, productName, productID, count, childs}
			resulst[keys] = append(resulst[keys], answ)
		}
	}
	var stands []key

	for key, _ := range resulst {
		stands = append(stands, key)
	}

	sort.Slice(stands, func(i, j int) bool { return stands[i].name < stands[j].name })

	fmt.Println("=+=+=+=")
	fmt.Println("Страница сборки заказов", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arguments)), ","), "[]"))
	fmt.Println()

	for _, name := range stands {
		fmt.Println("===Стеллаж", name.name)
		for _, answ := range resulst[name] {
			fmt.Printf("%s, (id=%d)\n", answ.productName, answ.productID)
			fmt.Printf("заказ %d, %d шт\n", answ.orderID, answ.count)
			fmt.Println()
		}
	}

}
