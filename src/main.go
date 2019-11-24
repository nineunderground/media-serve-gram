/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   24-Nov-2019
 * @Project: media-serve-gram
 * @Filename: main.go
 * @Copyright: Iñaki Rodriguez
 */
package main

import (
  "github.com/schollz/progressbar"
  "time"
  "fmt"
  "os"
  //"./utils"
)

func main() {
  testProgressBar()
  //telegram.RunBot()
  os.Exit(0)
}

func testProgressBar() {
  bar := progressbar.New(100)
  for i := 0; i < 100; i++ {
      bar.Add(1)
      time.Sleep(3 * time.Millisecond)
  }
  fmt.Println("")
  fmt.Println("APP LOADED!")
  VARIABLE := os.Getenv("TELEGRAM_TOKEN")
  
  fmt.Println("VARIABLE:")
  fmt.Println(VARIABLE)
}
