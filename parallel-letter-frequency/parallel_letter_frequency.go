package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	var freqMapChannel = make(chan FreqMap)
	var freqMapFinal = make(FreqMap)

	wg.Add(len(l))
	for _, s := range l {
		go func(s string) {
			defer wg.Done()
			freqMapChannel <- Frequency(s)
		}(s)
	}

	go func() {
		wg.Wait()
		close(freqMapChannel)
	}()

	for m := range freqMapChannel {
		for k, v := range m {
			freqMapFinal[k] += v
		}
	}
	return freqMapFinal
}
