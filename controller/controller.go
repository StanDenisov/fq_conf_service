package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/StanDenisov/utils/confstruct"

	"github.com/labstack/echo"
)

//GetConf - return conf by app_name
func GetConf(c echo.Context) error {
	conf := confstruct.ConfStruct{}
	serviceName := c.Param("service_name")
	if serviceName == "favicon.ico" {
		return c.String(http.StatusOK, "ok")
	}
	fmt.Println("name is :", serviceName)
	appPath := fmt.Sprintf("configurations/%s.json", serviceName)
	jsonFile, err := os.Open(appPath)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &conf); err != nil {
		return c.String(http.StatusNoContent, "unmarshal filed")
	}
	defer jsonFile.Close()
	fmt.Println(conf)
	return c.JSON(http.StatusOK, conf)
}
