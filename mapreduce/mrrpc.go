package mapreduce

import (
	"os"
	"strconv"
)

func masterSock() string {
	s := "/var/tmp/mr"

	s += strconv.Itoa(os.Getuid())
	return s
}
