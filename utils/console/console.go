package console

import (
	"fmt"

	logging "github.com/sirupsen/logrus"
)

func Config(data ...interface{}) {
	logging.SetFormatter(&logging.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}
func ConfigText(data ...interface{}) string {
	Config()
	str := ""
	for index := 0; index < len(data); index++ {
		str1 := fmt.Sprintf("%v", data[index])
		if len(data) == 0 || len(data) == 1 {
			str += "" + str1
		} else {
			str += "" + str1 + " "
		}
	}
	return str
}
func Info(data ...interface{}) {
	str := ConfigText(data...)
	logging.Info(str)
}
func Warning(data ...interface{}) {
	str := ConfigText(data...)
	logging.Warning(str)
}
func Error(data ...interface{}) {
	str := ConfigText(data...)
	logging.Error(str)
}
func Fatal(data error) {
	logging.Fatal(data)
}
