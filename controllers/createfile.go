package Controllers

import (
	"os"
	"encoding/csv"
  
	"local.packages/DB"
	"local.packages/Models"
  
	"github.com/labstack/echo/v4"
	 _ "github.com/lib/pq"
  )

  func GetTabledate(tablename string)Models.Table{
	db := DB.GormConnect()
	defer db.Close()
	table := Models.Table{}
	db.First(&table,"table_name=?",tablename)
	return table
  }

  func WriteFile(table Models.Table) string{
	  file,err := os.Create("./DDL/sample.csv")
	  if err != nil{
		  panic(err)
	  }
	  defer file.Close()
	  file.Truncate(0)
	  writer := csv.NewWriter(file)
	  writer.Write([]string{"テーブル情報"})
	  writer.Write([]string{"テーブル名", "テーブル和名", "更新情報", "説明"})
	  writer.Write([]string{table.Table_name,"",table.Update_info,""})
	  writer.Flush()

	  return "生成完了"

  }

  func CreateFile() echo.HandlerFunc{
	  return func(c echo.Context) error{
		  table := GetTabledate("sample_table")
		  mes := WriteFile(table)
		  return DB.ReturnResponse(mes, c)
	  }
  }