package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/a-17-1-penggunaan-map", func(c echo.Context) error {

		chicken := make(map[string]int) // Menggunakan make() untuk membuat map

		chicken["januari"] = 50
		chicken["februari"] = 40

		fmt.Println("januari", chicken["januari"]) // januari 50
		fmt.Println("mei", chicken["mei"])         // mei 0

		return c.String(http.StatusOK, "Success") // Menambahkan return statement
	})

	e.GET("/a-17-2-inisialisasi-nilai-map", func(c echo.Context) error {

		// cara horizontal
		chicken1 := map[string]int{"januari": 50, "februari": 40}

		// cara vertikal
		chicken2 := map[string]int{
			"januari":  50,
			"februari": 40,
		}

		fmt.Println("januari", chicken1["januari"])   // januari 50
		fmt.Println("februari", chicken1["februari"]) // februari 40

		fmt.Println("januari", chicken2["januari"])   // januari 50
		fmt.Println("februari", chicken2["februari"]) // februari 40

		return c.String(http.StatusOK, "Success")
	})

	e.GET("/a-17-3-iterasi-item-map-Menggunakan-for-range", func(e echo.Context) error {

		chicken := map[string]int{
			"januari":  50,
			"februari": 40,
			"maret":    34,
			"april":    67,
		}

		for key, val := range chicken {
			fmt.Println(key, "\t:", val)
		}

		return e.String(http.StatusOK, "Success")
	})

	e.GET("/a-17-4-menghapus-item-map", func(f echo.Context) error {

		var chicken = map[string]int{"januari": 50, "februari": 40}

		fmt.Println(len(chicken)) // 2
		fmt.Println(chicken)

		delete(chicken, "januari")

		fmt.Println(len(chicken)) // 1
		fmt.Println(chicken)

		return f.String(http.StatusOK, "Success")
	})

	e.GET("/a-17-5-deteksi-keberadaan-item-dengan-key-tertentu", func(g echo.Context) error {

		var chicken = map[string]int{"januari": 50, "februari": 40}
		var value, isExist = chicken["mei"]

		if isExist {
			fmt.Println(value)
		} else {
			fmt.Println("item is not exists")
		}

		return g.String(http.StatusOK, "Success")
	})

	e.GET("a-17-6-kombinasi-slice&map", func(h echo.Context) error {

		var chickens = []map[string]string{
			map[string]string{"name": "chicken blue", "gender": "male"},
			map[string]string{"name": "chicken red", "gender": "male"},
			map[string]string{"name": "chicken yellow", "gender": "female"},
		}

		for _, chicken := range chickens {
			fmt.Println(chicken["gender"], chicken["name"])
		}
		return h.String(http.StatusOK, "Success")
	})
	e.Start(":9000")
}
