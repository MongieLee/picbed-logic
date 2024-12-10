package global

import (
	"picbed/utils"
)

func InitGlobalEnv() {
	InitViperConfig()
	InitMysqlDbConnector()
	utils.InitZeroLog()
	utils.InitValidatorTrans("zh")
}
