package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestUnique(t *testing.T) {
	asrt := assert.New(t)

	type args struct {
		Name  []Name
		Place []Place
	}
	tests := []struct {
		name string
		args args
		want []Character
	}{
		{
			name: "",
			args: args{
				Name: []Name{
					{
						Id:   1111,
						Name: "Leoric",
					},
					{
						Id:   2222,
						Name: "Cain",
					},
					{
						Id:   3333,
						Name: "Azmodan",
					},
					{
						Id:   4444,
						Name: "Tyrael",
					},
					{
						Id:   6666,
						Name: "Ba`al",
					},
				},
				Place: []Place{
					{
						Id:    1111,
						Place: "Act 1",
					},
					{
						Id:    2222,
						Place: "Act 2",
					},
					{
						Id:    3333,
						Place: "Act 3",
					},
					{
						Id:    5555,
						Place: "Act 5",
					},
					{
						Id:    7777,
						Place: "Act 7",
					},
					{
						Id:    8888,
						Place: "Act 8",
					},
				},
			},
			want: []Character{
				{
					Id:    1111,
					Name:  "Leoric",
					Place: "Act 1",
				},
				{
					Id:    2222,
					Name:  "Cain",
					Place: "Act 2",
				},
				{
					Id:    3333,
					Name:  "Azmodan",
					Place: "Act 3",
				},
				{
					Id:    4444,
					Name:  "Tyrael",
					Place: "",
				},
				{
					Id:    5555,
					Name:  "",
					Place: "Act 5",
				},
				{
					Id:    6666,
					Name:  "Ba`al",
					Place: "",
				},
				{
					Id:    7777,
					Name:  "",
					Place: "Act 7",
				},
				{
					Id:    8888,
					Name:  "",
					Place: "Act 8",
				},
			},
		},
	}

	for key, tt := range tests {
		name := fmt.Sprintln(key, tt.name)

		got := listSortedChars(tt.args.Name, tt.args.Place)

		asrt.Equal(tt.want, got, name)
	}
}

func listSortedChars(name []Name, place []Place) []Character {
	result := make([]Character, 0)

	for i := 0; i < len(name); i++ {
		char := Character{Id: name[i].Id, Name: name[i].Name}

		for j := 0; j < len(place); j++ {
			if place[j].Id == name[i].Id {
				char.Place = place[j].Place
			}
		}

		result = append(result, char)
	}


	for i := 0; i < len(place); i++ {
		var exist bool
		for j := 0; j < len(result); j++ {
			if result[j].Id == place[i].Id {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, Character{Id: place[i].Id, Place: place[i].Place})
		}
	}

	result = mergeSort(result)

	return result
}

func mergeSort(items []Character) []Character {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := num / 2
	var (
		left = make([]Character, middle)
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
	result = make([]Character, len(left) + len(right))

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