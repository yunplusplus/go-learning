package enums

import (
	"go.uber.org/zap"
	"testing"
)
import util "github.com/yunplusplus/go-learning/utils"

func TestWeekends(t *testing.T) {
	util.Logger.Info("weekends", zap.Int("monday", Monday))
	util.Logger.Info("weekends", zap.Int("tuesday", Tuesday))
	util.Logger.Info("weekends", zap.Int("Wednesday", Wednesday))
	util.Logger.Info("weekends", zap.Int("Thursday", Thursday))
	util.Logger.Info("weekends", zap.Int("Friday", Friday))
	util.Logger.Info("weekends", zap.Int("Saturday", Saturday))
	util.Logger.Info("weekends", zap.Int("Sunday", Sunday))
	util.MainLogger.Info("weekends", zap.Int("Sunday", Sunday))
}
