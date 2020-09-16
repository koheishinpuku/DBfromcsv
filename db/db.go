package DB

import(
  "fmt"
  "os"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "net/http"
  "github.com/labstack/echo/v4"
  "encoding/json"
  "local.packages/Models"
)

func ReturnResponse(mes string, c echo.Context) error {
	message :=  Models.Message{}
	json.Unmarshal([]byte(`{"Mess": "`+mes+`"}`), &message)
	return c.JSON(http.StatusOK, message)
  }

func GormConnect() *gorm.DB{
	var isDev bool = os.Getenv("GO_ENV") == "development"
	var user = os.Getenv("DB_USER")
	var pwd = os.Getenv("DB_PWD")
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var database = os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, database)
  	db, err := gorm.Open("mysql", connString)

	if err != nil{
		panic(err.Error())
	}
	db.LogMode(isDev)
	return db

  // defer db.Close()
}
