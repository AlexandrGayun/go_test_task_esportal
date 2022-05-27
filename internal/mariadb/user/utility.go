package user

import (
	"github.com/AlexandrGayun/go_test_task_esportal/internal/mariadb/user/db"
	"sort"
)

//AtLeastTwice is possible in Two variants:
// 1)enumerate all values, comparing each other(with double loop)
// 2)sort array ascending/descending and compare nested elements in one loop
func AtLeastTwice(users []db.User) bool {
	sort.Slice(users, func(i int, j int) bool {
		return users[i].Age < users[j].Age
	})
	min := users[0].Age
	for i := 0; i < len(users)-1; i++ {
		if (min * 2) <= users[i+1].Age {
			return true
		} else if users[i+1].Age < min {
			min = users[i+1].Age
		}
	}
	return false
}

//ExactlyTwice will require double loop, but we could skip some iterations if the slice is sorted
//Possibly on some datasets sorting will give more overhead
func ExactlyTwice(users []db.User) bool {
	sort.Slice(users, func(i int, j int) bool {
		return users[i].Age < users[j].Age
	})
	found := false
	for i := 0; i < len(users)-1; i++ {
		if found {
			break
		}
		currElem := users[i].Age
		for j := 0; j < len(users)-1; j++ {
			if (currElem * 2) == users[j+1].Age {
				found = true
				break
			} else if (currElem * 2) < users[j].Age {
				//no sense to check remained elements, they`re all bigger than double current element
				break
			}
		}
	}
	return found
}

//ConstrainedExactlyTwice input age restricted within the range 18 to 80
//The idea is to use map as hash table with O(1) - O(n) search/insertion
//which is faster than iterating over two loops with O(n*n) in the worst scenario
//Not sure how to use restriction to make it faster because according to benchmarks
//range over map is still faster than for loop over the permitted age values
//Benchmark shows that ConstrainedExactlyTwice at least 30% faster than ExactlyTwice
//even on the big sets
//
//ExactlyTwice/Low_size_set#05-20     490ns ± 0%     346ns ± 0%   -29.37%
//ExactlyTwice/10k_size_set-20        545µs ± 0%     371µs ± 0%   -31.85%
//ExactlyTwice/100k_size_set-20      6.02ms ± 0%    3.83ms ± 0%   -36.48%
//
//but in counterpart it loses in allocation bytes size
//there is a strange peak on middle size sets where strict solution lose 26% in size of alloc bytes
//
//ExactlyTwice/Low_size_set#05-20      168B ± 0%       48B ± 0%   -71.43%
//ExactlyTwice/10k_size_set-20       3.09kB ± 0%    3.91kB ± 0%   +26.36%
//ExactlyTwice/100k_size_set-20       471kB ± 0%     270kB ± 0%   -42.63%
//
//and significantly loses in allocations per ops
//
//ExactlyTwice/Low_size_set#05-20      3.00 ± 0%      2.00 ± 0%   -33.33%
//ExactlyTwice/10k_size_set-20         3.00 ± 0%     13.00 ± 0%  +333.33%
//ExactlyTwice/100k_size_set-20        3.00 ± 0%     13.00 ± 0%  +333.33%
//
//we can conclude that strict realization wins in performance but loses in memory usage

func ConstrainedExactlyTwice(users []db.User) bool {
	m := make(map[int32]byte)
	for i := 0; i < len(users); i++ {
		age := users[i].Age
		if _, ok := m[age]; !ok {
			m[age] = 1
		}
	}
	for k := range m {
		if _, ok := m[k*2]; ok {
			return true
		}
	}
	return false
}
