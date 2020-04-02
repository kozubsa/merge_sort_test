package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type Name struct {
	Id   int32
	Name string
}

type Place struct {
	Id    int32
	Place string
}

type Character struct {
	Id    int32
	Name  string
	Place string
}

type args struct {
	Names  []Name
	Places []Place
}

type TestCase = struct {
	args args
	want []Character
}

func FillTestCase() TestCase {
	names := make([]Name, 0, 100)
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			continue
		}
		names = append(names, Name{
			Id:   int32(i * 1111),
			Name: faker.WORD,
		})
	}
	places := make([]Place, 0, 100)
	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			continue
		}
		places = append(places, Place{
			Id:    int32(i * 1111),
			Place: faker.Word(),
		})
	}
	return TestCase{
		args: args{
			Names:  names,
			Places: places,
		},
		want: []Character{
			{Id: 1111, Name: "Leoric", Place: "Act 1"},
			{Id: 2222, Name: "Cain", Place: "Act 2"},
			{Id: 3333, Name: "Azmodan", Place: "Act 3"},
			{Id: 4444, Name: "Tyrael", Place: ""},
			{Id: 5555, Name: "", Place: "Act 5"},
			{Id: 6666, Name: "Ba`al", Place: ""},
			{Id: 7777, Name: "", Place: "Act 7"},
			{Id: 8888, Name: "", Place: "Act 8"},
		},
	}
}

func BenchmarkMergeSort(b *testing.B) {
	// asrt := assert.New(b)
	testCase := FillTestCase()
	result := make([]Character, 0)

	for i := 0; i < len(testCase.args.Names); i++ {
		char := Character{Id: testCase.args.Names[i].Id, Name: testCase.args.Names[i].Name}

		for j := 0; j < len(testCase.args.Places); j++ {
			if testCase.args.Places[j].Id == testCase.args.Names[i].Id {
				char.Place = testCase.args.Places[j].Place
			}
		}

		result = append(result, char)
	}

	for i := 0; i < len(testCase.args.Places); i++ {
		var exist bool
		for j := 0; j < len(result); j++ {
			if result[j].Id == testCase.args.Places[i].Id {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, Character{Id: testCase.args.Places[i].Id, Place: testCase.args.Places[i].Place})
		}
	}

	result = mergeSort(result)
	// asrt.Equal(testCase.want, result)
}

func TestMergeSort(t *testing.T) {
	asrt := assert.New(t)
	testCase := FillTestCase()
	result := make([]Character, 0)

	for i := 0; i < len(testCase.args.Names); i++ {
		char := Character{Id: testCase.args.Names[i].Id, Name: testCase.args.Names[i].Name}

		for j := 0; j < len(testCase.args.Places); j++ {
			if testCase.args.Places[j].Id == testCase.args.Names[i].Id {
				char.Place = testCase.args.Places[j].Place
			}
		}

		result = append(result, char)
	}

	for i := 0; i < len(testCase.args.Places); i++ {
		var exist bool
		for j := 0; j < len(result); j++ {
			if result[j].Id == testCase.args.Places[i].Id {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, Character{Id: testCase.args.Places[i].Id, Place: testCase.args.Places[i].Place})
		}
	}

	resultMergeSort := mergeSort(result)

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})

	asrt.Equal(resultMergeSort, result)
}

func BenchmarkSort(b *testing.B) {
	// asrt := assert.New(b)
	testCase := FillTestCase()
	result := make([]Character, 0)

	for i := 0; i < len(testCase.args.Names); i++ {
		char := Character{Id: testCase.args.Names[i].Id, Name: testCase.args.Names[i].Name}

		for j := 0; j < len(testCase.args.Places); j++ {
			if testCase.args.Places[j].Id == testCase.args.Names[i].Id {
				char.Place = testCase.args.Places[j].Place
			}
		}

		result = append(result, char)
	}

	for i := 0; i < len(testCase.args.Places); i++ {
		var exist bool
		for j := 0; j < len(result); j++ {
			if result[j].Id == testCase.args.Places[i].Id {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, Character{Id: testCase.args.Places[i].Id, Place: testCase.args.Places[i].Place})
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})
}

func mergeSort(items []Character) []Character {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := num / 2
	var (
		left  = make([]Character, middle)
		right = make([]Character, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []Character) (result []Character) {
	result = make([]Character, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].Id < right[0].Id {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
