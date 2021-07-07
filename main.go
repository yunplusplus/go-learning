package main

import (
	"encoding/json"
	dubbogost "github.com/dubbogo/gost/encoding/json"
	"go.uber.org/zap"
)

func main() {
	var logger, _ = zap.NewProduction()
	logger.Info("Success", zap.String("name", "赵云天"))
	path := "./user.json"
	var hessianPair, _, _ = dubbogost.File2Interface(path)
	var content, _ = json.Marshal(hessianPair)
	logger.Info("hessianSerialize", zap.String("hessianPair",string(content)))

}
