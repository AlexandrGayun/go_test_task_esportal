package user

import (
	"fmt"
	"github.com/AlexandrGayun/go_test_task_esportal/internal/mariadb/user/db"
	"math/rand"
	"testing"
	"time"
)

type testDataInstance struct {
	users          []db.User
	expectedResult bool
}

var dataset = []testDataInstance{
	{[]db.User{{Age: 10}, {Age: 20}, {Age: 12}}, true},
	{[]db.User{{Age: 10}, {Age: 12}, {Age: 19}, {Age: 24}}, true},
	{[]db.User{{Age: 10}, {Age: 20}, {Age: 20}, {Age: 1}, {Age: 2}}, true},
	{[]db.User{{Age: 11}, {Age: 12}, {Age: 19}, {Age: 19}}, false},
	{[]db.User{{Age: 6}, {Age: 9}, {Age: 7}, {Age: 11}}, false},
	{[]db.User{{Age: 7}, {Age: 7}, {Age: 4}, {Age: 4}}, false},
}
var datasetConstrained = []testDataInstance{
	{[]db.User{{Age: 18}, {Age: 36}, {Age: 19}}, true},
	{[]db.User{{Age: 19}, {Age: 40}, {Age: 80}, {Age: 35}}, true},
	{[]db.User{{Age: 18}, {Age: 36}, {Age: 36}, {Age: 40}, {Age: 41}}, true},
	{[]db.User{{Age: 18}, {Age: 19}, {Age: 29}, {Age: 29}}, false},
	{[]db.User{{Age: 23}, {Age: 26}, {Age: 34}, {Age: 35}}, false},
	{[]db.User{{Age: 30}, {Age: 30}, {Age: 40}, {Age: 40}}, false},
}

func TestAtLeastTwice(t *testing.T) {
	for k, tc := range dataset {
		t.Run(fmt.Sprintf("Example #%v", k), func(t *testing.T) {
			actualResult := AtLeastTwice(tc.users)
			if tc.expectedResult != actualResult {
				t.Errorf("Expected %t, got %t on the input set %v", tc.expectedResult, actualResult, tc.users)
			}
		})
	}
}

func TestExactlyTwice(t *testing.T) {
	for k, tc := range dataset {
		t.Run(fmt.Sprintf("Example #%v", k), func(t *testing.T) {
			actualResult := ExactlyTwice(tc.users)
			if tc.expectedResult != actualResult {
				t.Errorf("Expected %t, got %t on the input set %v", tc.expectedResult, actualResult, tc.users)
			}
		})
	}
}

func TestConstrainedExactlyTwice(t *testing.T) {
	for k, tc := range datasetConstrained {
		t.Run(fmt.Sprintf("Example #%v", k), func(t *testing.T) {
			actualResult := ConstrainedExactlyTwice(tc.users)
			if tc.expectedResult != actualResult {
				t.Errorf("Expected %t, got %t on the input set %v", tc.expectedResult, actualResult, tc.users)
			}
		})
	}
}

func BenchmarkExactlyTwice(b *testing.B) {
	for _, tc := range dataset {
		b.Run("Low size sets", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ExactlyTwice(tc.users)
			}
		})
	}
	b.Run("10k size set", func(b *testing.B) {
		users := GenerateUsers(10000, false)
		for i := 0; i < b.N; i++ {
			ExactlyTwice(users)
		}
	})
	b.Run("100k size set", func(b *testing.B) {
		users := GenerateUsers(100000, false)
		for i := 0; i < b.N; i++ {
			ExactlyTwice(users)
		}
	})
}

func BenchmarkConstrainedExactlyTwice(b *testing.B) {
	for _, tc := range datasetConstrained {
		b.Run("Low size sets", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConstrainedExactlyTwice(tc.users)
			}
		})
	}
	b.Run("10k size set", func(b *testing.B) {
		users := GenerateConstrainedUsers(10000, false)
		for i := 0; i < b.N; i++ {
			ConstrainedExactlyTwice(users)
		}
	})
	b.Run("100k size set", func(b *testing.B) {
		users := GenerateConstrainedUsers(100000, false)
		for i := 0; i < b.N; i++ {
			ConstrainedExactlyTwice(users)
		}
	})
}

func GenerateUsers(count int, withRandSeed bool) []db.User {
	if withRandSeed {
		rand.Seed(time.Now().UnixNano())
	}
	sl := make([]db.User, count)
	for i := 0; i < count; i++ {
		rndAge := rand.Int31n(125)
		u := db.User{Age: rndAge}
		sl = append(sl, u)
	}
	return sl
}

func GenerateConstrainedUsers(count int, withRandSeed bool) []db.User {
	if withRandSeed {
		rand.Seed(time.Now().UnixNano())
	}
	sl := make([]db.User, count)
	for i := 0; i < count; i++ {
		rndAge := 18 + rand.Int31n(63)
		u := db.User{Age: rndAge}
		sl = append(sl, u)
	}
	return sl
}
