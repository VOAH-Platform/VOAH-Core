package directory

import (
	"os"
	"sync"

	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/utils/logger"
)

func InitDirectory(wait *sync.WaitGroup) {
	serverConf := configs.Env.Server

	if _, err := os.Stat(serverConf.DataDir + "/user-profiles"); os.IsNotExist(err) {
		err := os.Mkdir(serverConf.DataDir+"/user-profiles", 0755)
		if err != nil {
			log := logger.Logger
			log.Fatal(err)
		}
	}
	defer wait.Done()
}
