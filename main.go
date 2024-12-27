package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	// Create a reader for standard input
	reader := bufio.NewReader(os.Stdin)

	// Prompt for input
	fmt.Print("Enter the Range(int): ")

	// Read input and trim spaces
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)

	// Convert input to uint64
	n, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Start timing
	start := time.Now()

	// Choose appropriate method based on input size
	var primes []uint64
	if n <= 10000000 { // 10^7
		primes = SieveOfEratosthenes(n)
	} else {
		primes = Segmented_SOE(n)
	}

	// Calculate duration
	duration := time.Since(start)

	// Create filename with range
	filename := fmt.Sprintf("primes_%d.csv", n)

	// Create and write to file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Write each prime number to the file
	for _, prime := range primes {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", prime))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Flush the writer to ensure all data is written
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Print execution time and success message
	fmt.Println("Time taken:", duration)
	fmt.Printf("File '%s' has been successfully created\n", filename)
}

// This is without segmenting.
func SieveOfEratosthenes(n uint64) []uint64 {
	cores := runtime.NumCPU()
	next := make(chan bool, cores)
	var nums = make([]bool, n/2+1)
	m := uint64(math.Sqrt(float64(n)))

	for i := uint64(3); i <= m; i = i + 2 {
		if nums[i/2] == false {
			go goFill(nums, i, n, next)
			next <- true
		}
	}

	for i := 0; i < cores; i++ {
		next <- true
	}

	var ps []uint64
	if n >= 2 {
		ps = append(ps, 2)
	}
	for i := uint64(3); i <= n; i = i + 2 {
		if nums[i/2] == false {
			ps = append(ps, i)
		}
	}
	return ps
}

func fill(nums []bool, i uint64, max uint64) {
	// a := 3 * i
	iteration := 0
	a := i * i
	for a <= max {
		iteration++
		nums[a/2] = true
		a = a + 2*i
	}
}

func goFill(nums []bool, i uint64, max uint64, next chan bool) {
	fill(nums, i, max)
	<-next
}

// Segmented Sieve
var csegPool sync.Pool

func fillSegments(n uint64, basePrimes []uint64, allPrimes *[]uint64, segSize uint64, segNum uint64, next chan bool, nextTurn []chan bool) {
	cseg := (csegPool.Get()).([]bool)
	for i := uint64(0); i < segSize; i++ {
		cseg[i] = false
	}

	segEnd := segSize * (segNum + 1)

	for i := 0; i < len(basePrimes); i++ {
		p := basePrimes[i]
		pSquare := p * p

		if pSquare > segEnd {
			continue
		}

		jMax := segSize * (segNum + 1) / basePrimes[i]

		startJ := basePrimes[i] - 1
		if startJ < (segSize*segNum)/basePrimes[i] {
			startJ = (segSize * segNum) / basePrimes[i]
		}

		for j := startJ; j < jMax; j++ {
			sn := (j + 1) * basePrimes[i]
			cseg[sn-segSize*segNum-1] = true
		}
	}

	if segNum > 1 {
		<-nextTurn[segNum]
	}

	for i := uint64(0); i < segSize; i++ {
		if !cseg[i] && segSize*segNum+i+1 <= n {
			*allPrimes = append(*allPrimes, segSize*segNum+i+1)
		}
	}

	<-next
	if int(segNum)+1 < len(nextTurn) {
		nextTurn[segNum+1] <- true
	}

	csegPool.Put(cseg)
}

func Segmented_SOE(n uint64) (allPrimes []uint64) {
	allPrimes = make([]uint64, 0, n/uint64(math.Log(float64(n))-1))

	segSize := uint64(math.Sqrt(float64(n)))

	csegPool.New = func() interface{} {
		return make([]bool, segSize)
	}

	basePrimes := SieveOfEratosthenes(segSize)
	allPrimes = append(allPrimes, basePrimes...)

	cores := runtime.NumCPU()
	next := make(chan bool, cores)
	var nextTurn []chan bool
	nextTurn = make([]chan bool, n/segSize+1)
	for i := uint64(0); i < n/segSize+1; i++ {
		nextTurn[i] = make(chan bool)
	}
	for segNum := uint64(1); segNum <= n/segSize-1; segNum++ {
		go fillSegments(n, basePrimes, &allPrimes, segSize, segNum, next, nextTurn)
		next <- true
	}
	for i := 0; i < cores; i++ {
		next <- true
	}

	return allPrimes
}
