# Information Protect Calculator

Калькулятор для работы с большими числами (`BigDigit`) с поддержкой следующих операций:

* Сложение (`Sum`)
* Вычитание (`Sub`)
* Умножение (`Mul`)
* Деление (`Div`)
* Возведение в степень (`Pow`)
* НОД (`GCD`)
* НОК (`LCM`)

---

## Структура проекта

```
cmd/                       # точка входа для сборки бинарника
internal/app/              # основная логика приложения: файлы или генерация чисел
internal/converter/        # чтение и запись BigDigit из файлов
internal/generator/        # генерация случайных больших чисел
internal/model/            # структура BigDigit
internal/service/calculator/ # реализация операций калькулятора
internal/utils/            # вспомогательные функции
```

---

## Сборка и запуск

```bash
make build       # сборка бинарника в bin/information-protect
make run         # сборка и запуск
make clean       # удаление бинарников из bin
```

---

## Генерация случайных чисел

Для генерации случайных больших чисел и вычислений:

```bash
make build &&
./bin/information-protect -size=10 -pow=9
```

* `-size` — количество цифр в случайных числах
* `-pow` — степень для операции возведения в степень

---

## Использование файлов с числами

Создание файлов с числами:

```bash
make files
```

Запуск программы с файлами:

```bash
make build &&
./bin/information-protect -file1=in/num1.txt -file2=in/num2.txt -pow=9
```

Результаты сохраняются в папку `out/`:

```
sum.txt, sub.txt, mul.txt, pow.txt, quot.txt, rem.txt, gcd.txt, lcm.txt
```

---

## Требования

* Go 1.21+


## TODO

- не работает со знаками плюс/минус, нужно с этими кейсами отдебажить