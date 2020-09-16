package Controllers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	// iconv "github.com/djimenez/iconv-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"

	"local.packages/DB"
	"local.packages/Models"
)

func ImportData() echo.HandlerFunc {
	return func(c echo.Context) error {
		upload_file, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		src, err := upload_file.Open()
		if err != nil {
			panic(err)
		}
		defer src.Close()

		dst_file, err := os.Create("./DDL/full.csv")
		if err != nil {
			panic(err)
		}
		defer dst_file.Close()

		if _, err = io.Copy(dst_file, src); err != nil {
			panic(err)
		}

		mess := InsertDB()
		fmt.Println(mess)
		return DB.ReturnResponse(mess, c)
	}
	// return func(c echo.Context) error {
	// 	name := c.FormValue("name")
	// 	return c.String(http.StatusOK, "create name = "+name)
	// }
}

func InsertDB() string {
	db := DB.GormConnect()
	defer db.Close()

	fp, err := os.Open("./DDL/full.csv")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// converter, err := iconv.NewReader(fp, "sjis", "utf-8")
	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true
	count := 0
	for {
		record, err := reader.Read()
		// reader.FieldsPerRecord = -1 //csvのカラム数が揃っていなくても無視して処理を進める
		fmt.Println(record)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("break!!")
			return err.Error()
		}

		if (count == 0) || (count == 1) {
			count += 1
			fmt.Println(count)
			continue
		}

		if count == 2 {
			newTable := Models.Table{record[0], record[1], record[2], record[3]}
			db.Create(&newTable)
		}
		count += 1
	}
	return "成功"
}
