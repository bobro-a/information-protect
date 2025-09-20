package main

import (
	"flag"
	"fmt"

	"information-protect/internal/app"
)

func main() {
	// Параметры запуска
	numFile1 := flag.String("file1", "", "Путь к первому числу")
	numFile2 := flag.String("file2", "", "Путь ко второму числу")
	size := flag.Int("size", 1, "Размер генерируемого числа, если файлы не заданы")
	pow := flag.Int64("pow", 9, "POW калькулятора")

	flag.Parse()

	cfg := app.AppConfig{
		NumFile1: *numFile1,
		NumFile2: *numFile2,
		NumSize:  *size,
		Pow:      *pow,
	}

	fmt.Println("Запуск приложения...")
	app.Run(cfg)
}
