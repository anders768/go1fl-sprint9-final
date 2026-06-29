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
		wholeSlice[i] = rand.Int()
	}
	return wholeSlice // возвращаем сгенерированный слайс
}

// Maximum returns the largest element of slice.
func Maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 { // обрабатываем нулевой размер слайса
		return -1 // 0 нельзя использовать, т.к. он может быть нормальным возвратом функции (теоретически)
	}

	var max int
	for _, v := range data { // собственно, ищем максимум
		if v > max {
			max = v
		}
	}
	return max
}

// MaxChunks returns the largest element of slice in chunks.
func MaxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 { // обрабатываем нулевой размер слайса
		return -1
	}

	var (
		sliceMax  = make([]int, CHUNKS) // слайс максимумов, общий для всех горутин
		remainder = len(data) % CHUNKS  // если есть остаток, значит, есть хвост
		lenChunk  = len(data) / CHUNKS  // длина среза для одной горутины (т.е. длина чанки), не считая хвоста
		chunk     []int                 // производный слайс, с которым будет работать горутина (срез от исходного)
		wg        sync.WaitGroup
	)

	wg.Add(CHUNKS) // количество горутин

	for i := range CHUNKS { // запускаем горутины
		if remainder != 0 && i == CHUNKS-1 { // если хвост есть, и горутина - последняя
			// формируем срез, с которым будет работать горутина
			chunk = data[lenChunk*i : lenChunk*i+lenChunk+remainder] // хвост не забыт
		} else { // в прочих случаях - простой порядок (без хвоста)
			chunk = data[lenChunk*i : lenChunk*i+lenChunk] // формируем срез, с которым будет работать горутина
		}

		go func(i int, chunk []int) {
			defer wg.Done() // откладываем уменьшение счётчика горутин

			sliceMax[i] = Maximum(chunk)
		}(i, chunk)

	}
	wg.Wait()
	return Maximum(sliceMax)
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
	if max == -1 { // обрабатываем нулевой размер слайса
		fmt.Println("Дубина, нужен слайс ненулевой длины! Всё, финиш!")
		return
	}

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	before = time.Now() // начальная отсечка выполнения функции
	max = MaxChunks(slice)
	elapsed = time.Since(before).Microseconds()
	if max == -1 { // обрабатываем нулевой размер слайса
		fmt.Println("Дубина, нужен слайс ненулевой длины! Всё, финиш!")
		return
	}

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", max, elapsed)
}
