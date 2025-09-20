package app

import (
	"fmt"
	"log"
	"os"

	"information-protect/internal/converter"
	"information-protect/internal/generator"
	"information-protect/internal/model"
	"information-protect/internal/service/calculator"
)

// AppConfig хранит настройки приложения
type AppConfig struct {
	NumFile1 string
	NumFile2 string
	NumSize  int // количество цифр, если будем генерировать
	Pow      int64
}

// Run выполняет работу приложения
func Run(cfg AppConfig) {
	// Создаем калькулятор
	calc := calculator.NewCalculator(cfg.Pow)

	var num1, num2 model.BigDigit
	var err error

	// Проверяем, переданы ли файлы
	if cfg.NumFile1 != "" && cfg.NumFile2 != "" {
		num1, err = converter.FromFile(cfg.NumFile1)
		if err != nil {
			log.Fatalf("Ошибка чтения %s: %v", cfg.NumFile1, err)
		}
		num2, err = converter.FromFile(cfg.NumFile2)
		if err != nil {
			log.Fatalf("Ошибка чтения %s: %v", cfg.NumFile2, err)
		}
		fmt.Println("Числа загружены из файлов")
	} else {
		// Генерируем случайные числа
		num1 = generator.GenerateBigNumber(cfg.NumSize)
		num2 = generator.GenerateBigNumber(cfg.NumSize)
		fmt.Println("Числа сгенерированы случайным образом")
	}

	// Выполняем операции
	sum := calc.Sum(num1, num2)
	sub := calc.Sub(num1, num2)
	mul := calc.Mul(num1, num2)
	pow := calc.Pow(num1, 3)
	quot, rem := calc.Div(num1, num2)
	//gcd := calc.GCD(num1, num2)
	//lcm := calc.LCM(num1, num2)

	// Создаем папку для результатов
	outDir := "bin/out"
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Сохраняем результаты в файлы
	results := map[string]model.BigDigit{
		"sum.txt":  sum,
		"sub.txt":  sub,
		"mul.txt":  mul,
		"pow.txt":  pow,
		"quot.txt": quot,
		"rem.txt":  rem,
		//"gcd.txt":  gcd,
		//"lcm.txt":  lcm,
	}

	for name, val := range results {
		path := outDir + "/" + name
		if err := converter.ToFile(path, val); err != nil {
			log.Fatalf("Ошибка записи %s: %v", name, err)
		}
		fmt.Printf("Сохранено %s\n", path)
	}

	fmt.Println("Все операции выполнены и сохранены в", outDir)
}
