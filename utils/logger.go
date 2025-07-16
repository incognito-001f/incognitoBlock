package utils

import (
    "log"
    "os"
    "time"
)

var logger *log.Logger

func init() {
    file, _ := os.OpenFile("node.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    logger = log.New(file, "", log.LstdFlags)
}

func Log(msg string) {
    logger.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05")
}