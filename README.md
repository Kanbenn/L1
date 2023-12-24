## Устные вопросы

1. я бы подумал, что оптимальным должен быть `strings.Builder`, но [вот здесь](https://gist.github.com/dtjm/c6ebc86abe7515c988ec) есть бенчмарки, которые показывают, что самый простое складывание через знак `+` очень круто работает, а также bytes.Buffer с `buf.Reset()`.
2. интерфейсы в Go это именованный набор сигнатур методов. применяется для реализации полиморфизма.
3. RWMutex содержит RLock, который не блокирует конкурентных Reader'ов.
4. при работе с unbuffered-channel горутина зависает до момента ответа с обратной стороны канала. Uber настоятельно [рекомендует](https://github.com/uber-go/guide/blob/master/style.md#channel-size-is-one-or-none) не увлекаться буферными каналами.
5. 0 байт. крутой лайфхак 🙂👍
6. перегрузки операторов в Go нету. до версии 1.18 не было и п. методов, но теперь есть дженерики.
7. рандомно.
8. make() задумана для выделения памяти под сложные типы (слайсы, мапы и каналы), new() такого не сумеет.
9. всего способа четыре, но Uber [утверждает](https://github.com/uber-go/guide/blob/master/style.md#initializing-maps), что юзабельны из них два: make и литерал.
10. коварный вопрос, спасибо! этот код выведет `1 1`, что  связано с такими неочевидными темами как `variable shadowing` и `copy by value`, в функцию update попадает лишь копия указателя `p`. чтобы код в update() повлиял на main(), вместо `p = &b` нужно `*p = b`.
11. мне эта программа [вывела](https://go.dev/play/p/dIImieVIJEz) кучу ошибок, всё из-за того же `copy by value semantics`. варианты решения: передать [wg указателем](https://go.dev/play/p/0lIzCrepwhK) либо просто [через closure](https://go.dev/play/p/g84tWdNmiyz).
12. выводит `0` из-за `variable shadowing`. `if` создаеёт свой собственный namespace.
13. ещё один коварный и очень интересный вопрос, большое спасибо! 🙏 программа выводит `[100 2 3 4 5]`. Слайс в Go это ссылочный тип, поэтому изменение элементов слайса внутри someAction() изменяет изначальный слайс в main() как если бы мы передали в функцию someAction() переменную по указателю. Но `append()` создаёт новый слайс, ссылка на который кладётся в переменную с тем же именем v (снова тема `variable shadowing`), которая "умирает" при завершении `someAction()`. Это задание очень схоже с заданием №10.
14. программа выведет `[b b a][a a]` по причинам указанным в ответе на предыдущее задание №13.
