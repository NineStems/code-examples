# simple-http-calculator
example of simple http calculator  
пример простого веб калькулятора на golang 
## Примечание  
Задача похоже была взята с какого-то форума и засылась всем кандидатам без предварительной вычистки текста 
## Задание
Разработать "калькулятор", работающий по HTTP
## Методы
GET /
### Параметры (get-параметры)
- expr - строка с выражением. Например: ?expr=(2+2)2
### Результаты
- строка с результатом или ошибка
## Примечания
- должен учитываться приоритет операторов, в том числе скобки
- должны учитываться унарные операции (-2, +2-2)
- разделить для чисел с плавающей точкой - точка
- запрещено использовать сторонние решения. Только стандартная библиотека
- необходимо релизовать сложение, вычитание, произведение и деление - всё
## Подсказки
- не нужно писать свой парсер, он уже есть в стандартной библиотеке: гуглить evaluate formula in
go (первый результат на stack overflow)
- в Go есть возможность сопоставления типов, гуглить golang switch type
Задание звучит сложно, но по факту - очень простое.
У меня программа вышла в 92 строки (скриншот). Сложного на самом деле нет ничего.

