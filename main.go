package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// GenerateRandomElements generates random elements.
func GenerateRandomElements(size int) []int {
	// ваш код здесь
	if size == 0 { // обрабатываем нулевой размер слайса
		return nil
	}
	wholeSlice := make([]int, size) // создаём болванку слайса
	for i := range size {           // заполняем его (псевдо)случайными положительными значениями в количестве size элементов
		wholeSlice[i] = rand.Int() + 1 // чтоб точно положительными (т.к. ТЗ требует)
	}
	return wholeSlice // возвращаем сгенерированный слайс
}

// Maximum returns the largest element of slice.
func Maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 { // обрабатываем нулевой размер слайса
		return 0
	}
	var max int
	for _, v := range data {
		if v > 0 {
			if v > max {
				max = v
			}
		} else { // обрабатываем наличие в слайсе неположительных элементов
			return 0
		}
	}
	return max
}

// maxChunks returns the largest element of slice in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	var (
		maxSlice = make([]int, CHUNKS) // слайс максимумов, общий для всех горутин
		lenChunk = len(data) / CHUNKS  // длина среза для одной горутины (а если нацело не делится?)
		wg       sync.WaitGroup
	)

	wg.Add(CHUNKS)          // количество горутин
	for i := range CHUNKS { // запускаем горутины
		chunk := data[lenChunk*i : lenChunk*i+lenChunk] // формируем срез, с которым будет работать горутина
		go func(i int, chunk []int) {
			defer wg.Done() // откладываем уменьшение счётчика горутин

			var max int
			for _, v := range chunk {
				if v > max {
					max = v
				}
			}
			maxSlice[i] = max
			fmt.Printf("Горутина %d обработала %d элементов\n", i, len(chunk)) // проверка (поиск "хвоста")
		}(i, chunk)

	}
	wg.Wait()
	return Maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	slice := GenerateRandomElements(SIZE) // получаем сгенеренный слайс
	if slice == nil {                     // обрабатываем нулевой размер слайса
		fmt.Println("Дубина, нужен слайс ненулевой длины! Всё, финиш!")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	before := time.Now() // начальная отсечка выполнения функции
	max := Maximum(slice)
	elapsed := time.Since(before).Microseconds()
	if max == 0 { // обрабатываем нулевой размер слайса
		fmt.Println("Дубина, нужен слайс ненулевой длины и с положительными элементами! Всё, финиш!")
		return
	}

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	before = time.Now() // начальная отсечка выполнения функции
	max = maxChunks(slice)
	elapsed = time.Since(before).Microseconds()
	if max == 0 { // обрабатываем нулевой размер слайса
		fmt.Println("Дубина, нужен слайс ненулевой длины! Всё, финиш!")
		return
	}

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)
}
