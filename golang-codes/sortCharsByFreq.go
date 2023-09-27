// https://leetcode.com/problems/sort-characters-by-frequency

type Pair struct {
    Letter rune
    Frequency int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i].Frequency > h[j].Frequency
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    
    // store the element before popping out
    x := old[n-1]

    // slice off the last element as the heap now has one element less
    *h = old[:n-1]

    return x
}

func frequencySort(s string) string {
    mp := make(map[rune]int)

    for _, char := range s {
        mp[char]++
    }

    h := &MaxHeap{}
    heap.Init(h)

    for char, freq := range mp {
        pair := Pair {
            Letter: char,
            Frequency: freq,
        }

        heap.Push(h, pair)
    }

    //str := ""
    var str strings.Builder

    for h.Len() != 0 {
        pair := heap.Pop(h).(Pair)
        for i := 0; i < pair.Frequency; i++ {
            //str += (string)(pair.Letter)
            str.WriteRune(pair.Letter)
        }
    }

    //return str
    return str.String()
}