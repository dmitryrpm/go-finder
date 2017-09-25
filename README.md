**Задание**

Процессу на stdin приходят строки, содержащие URL или названия файлов. 
Что именно приходит на stdin определяется с помощью параметра командной строки -type. 
Например, -type file или -type url.
Каждый такой URL нужно запросить, каждый файл нужно прочитать, и посчитать кол-во вхождений строки "Go" в ответе.

В конце работы приложение выводит на экран общее кол-во найденных строк "Go" во всех источниках данных, например:

```bash
$ echo -e 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run 1.go -type url
Count for https://golang.org: 9
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 27
```

```bash
$ echo -e '/etc/passwd\n/etc/hosts' | go run 1.go - type file
Count for /etc/passwd: 0
Count for /etc/hosts: 0
Total: 0
```

Каждый источник данных должен начать обрабатываться сразу после вычитывания и параллельно с вычитыванием следующего. 
Источники должны обрабатываться параллельно, но не более k=5 одновременно. 

Обработчики данных не должны порождать лишних горутин, т.е. если k=1000 а обрабатываемых источников нет, 
не должно создаваться 1000 горутин.

Нужно обойтись без глобальных переменных и использовать только стандартные библиотеки. 

Код должен быть написан так, чтобы его можно было легко тестировать.
Формат предоставления решения: ссылка на github.

**Реализация**

Запускаем 5 воркеров и 3 задания (загрузка по URL)
```bash
>> make run 

echo 'https://golang.org\nhttps://golang.org\nhttps://golang.org' | go run cmd/run/main.go -type url 
Count for https://golang.org: 9
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 27
```

Запускаем 1 воркер и 2 заданий (загрузка из файла) 
```bash
>> make run_file flags="-k=1"

echo '/etc/passwd\n/etc/hosts' | go run cmd/run/main.go -type file -k=1
Count for /etc/passwd: 0
Count for /etc/hosts: 0
Total: 0
```

Тесты 
```bash
>> make test flags="-cover"
go test ./... -cover -tags=integration
ok      github.com/dmitryrpm/go-finder/cmd/run  0.003s  coverage: 64.3% of statements
ok      github.com/dmitryrpm/go-finder/config   0.008s  coverage: 100.0% of statements
ok      github.com/dmitryrpm/go-finder/finder   0.007s  coverage: 97.5% of statements

```
