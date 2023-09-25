package main

import (
	"github.com/sirupsen/logrus"
	"github.com/xpuls-com/xpuls-ml/cmd"
)

var logger = logrus.New().WithField("service", "xpuls-ml")

//func main() {
//
//	router, err := routes.NewRouter()
//	if err != nil {
//		fmt.Printf(err.Error())
//	}
//
//	err = router.Run(fmt.Sprintf(":%d", 5000))
//	if err != nil {
//		fmt.Printf(err.Error())
//		return
//	}
//}

func main() {
	cmd.Execute()
}
