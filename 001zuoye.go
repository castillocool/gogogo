//需要wraperror，可以让调用方知道哪句sql报错，若发现查询不到则记录查询sql，并继续执行下面的查询任务，以此类推，等待所有查询sql执行完毕后查看哪些sql返回失败。

package main

import (
	"errors"
	"log"
	"os"
)

func init() {
	file := "./" + "message" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[geektime]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {
	sqlerr := errors.New("these sql select nothing ")
	if sqlerr != nil {
		log.Println("Error: %w", sqlerr)
		return
	}

}
