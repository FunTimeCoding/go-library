package sweep

import "time"

func RunLoop(harbor string) {
	for {
		time.Sleep(10 * time.Minute)
		Run(harbor)
	}
}
