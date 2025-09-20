package app

import (
	"fmt"
	"log"
	"os"

	"information-protect/internal/config"
	"information-protect/internal/converter"
	"information-protect/internal/generator"
	"information-protect/internal/model"
	"information-protect/internal/service/calculator"
)

func Run(cfg config.AppConfig) {
	// Создаем калькулятор
	calc := calculator.NewCalculator()

	var num1, num2 model.BigDigit
	var err error

	outDir := "bin/out"
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

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

		// Сохраняем сгенерированные числа в файлы
		if err := converter.ToFile(outDir+"/num1.txt", num1); err != nil {
			log.Printf("Ошибка записи num1.txt: %v", err)
		}
		if err := converter.ToFile(outDir+"/num2.txt", num2); err != nil {
			log.Printf("Ошибка записи num2.txt: %v", err)
		}
		fmt.Println("Сгенерированные числа сохранены в файлы num1.txt и num2.txt")
	}

	// Канал для результатов операций
	type result struct {
		name string
		val  model.BigDigit
		err  error
	}
	resCh := make(chan result)

	// операции
	operations := []struct {
		name string
		fn   func() model.BigDigit
	}{
		{"sum.txt", func() model.BigDigit { return calc.Sum(num1, num2) }},
		{"sub.txt", func() model.BigDigit { return calc.Sub(num1, num2) }},
		{"mul.txt", func() model.BigDigit { return calc.Mul(num1, num2) }},
		{"pow.txt", func() model.BigDigit { return calc.Pow(num1, int(cfg.Pow)) }},
		{"quot.txt", func() model.BigDigit {
			q, _ := calc.Div(num1, num2)
			return q
		}},
		{"rem.txt", func() model.BigDigit {
			_, r := calc.Div(num1, num2)
			return r
		}},
		{"gcd.txt", func() model.BigDigit { return calc.GCD(num1, num2) }},
		{"lcm.txt", func() model.BigDigit { return calc.LCM(num1, num2) }},
	}

	// Запускаем операции в горутинах
	for _, op := range operations {
		go func(opName string, fn func() model.BigDigit) {
			defer func() {
				if r := recover(); r != nil {
					resCh <- result{opName, model.BigDigit{}, fmt.Errorf("%v", r)}
				}
			}()
			resCh <- result{opName, fn(), nil}
		}(op.name, op.fn)
	}

	// Сохраняем результаты из канала
	for range operations {
		r := <-resCh
		if r.err != nil {
			log.Printf("Ошибка при вычислении %s: %v", r.name, r.err)
			continue
		}
		path := outDir + "/" + r.name
		if err := converter.ToFile(path, r.val); err != nil {
			log.Printf("Ошибка записи %s: %v", r.name, err)
			continue
		}
		fmt.Printf("Сохранено %s\n", path)
	}

	fmt.Println("Все операции выполнены и сохранены в", outDir)
}
