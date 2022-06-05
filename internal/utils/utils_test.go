package utils

import (
	"log"
	"testing"
	"time"
)

func TestContains(t *testing.T) {
	now := time.Now()
	type item struct {
		ID   string
		num  int
		Name string
	}
	type order struct {
		ID        string
		CreatedAt time.Time
		UpdatedAt *time.Time
		Item      item
	}
	type args struct {
		list interface{}
		elem interface{}
	}
	tests := []struct {
		name string
		in   args
		want bool
	}{
		{
			name: "slice of struct",
			in: args{
				list: []item{
					{
						ID:   "1",
						Name: "test1",
					},
					{
						ID:   "2",
						Name: "test2",
					},
					{
						ID:   "3",
						num:  3,
						Name: "test3",
					},
				},
				elem: item{
					ID:   "3",
					num:  3,
					Name: "test3",
				},
			},
			want: true,
		},
		{
			name: "slice of struct 2",
			in: args{
				list: []order{
					{
						ID:        "1",
						CreatedAt: now,
						Item: item{
							ID:   "1",
							Name: "test1",
						},
					},
					{
						ID:        "2",
						CreatedAt: now,
						UpdatedAt: nil,
						Item: item{
							ID:   "3",
							Name: "test3",
						},
					},
				},
				elem: order{
					ID:        "2",
					CreatedAt: now,
					Item: item{
						ID:   "3",
						Name: "test3",
					},
				},
			},
			want: true,
		},
		{
			name: "slice of struct 3",
			in: args{
				list: []order{
					{
						ID:        "1",
						CreatedAt: now,
						Item: item{
							ID:   "1",
							Name: "test1",
						},
					},
					{
						ID:        "2",
						CreatedAt: now,
						UpdatedAt: &now,
						Item: item{
							ID:   "3",
							Name: "test3",
						},
					},
				},
				elem: order{
					ID:        "2",
					CreatedAt: now,
					UpdatedAt: &now,
					Item: item{
						ID:   "3",
						Name: "test3",
					},
				},
			},
			want: true,
		},
		{
			name: "slice of pointer",
			in: args{
				list: []*item{
					{
						ID:   "1",
						Name: "test1",
					},
					{
						ID:   "2",
						Name: "test2",
					},
					{
						ID:   "3",
						num:  3,
						Name: "test3",
					},
				},
				elem: &item{
					ID:   "3",
					num:  3,
					Name: "test3",
				},
			},
			want: true,
		},
		{
			name: "int32",
			in: args{
				list: []int32{1, 2, 3, 4, 5},
				elem: 4,
			},
			want: true,
		},
		{
			name: "int",
			in: args{
				list: []int{1, 2, 3, 4, 5},
				elem: 1,
			},
			want: true,
		},
		{
			name: "float64",
			in: args{
				list: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				elem: 3.3,
			},
			want: true,
		},
		{
			name: "string",
			in: args{
				list: []string{"apple", "orange", "lemon"},
				elem: "lemon",
			},
			want: true,
		},
		{
			name: "recover panic",
			in: args{
				list: []int{1, 2, 3},
				elem: "str",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Contains(tt.in.list, tt.in.elem)
			if tt.want != got {
				t.Errorf("testing %s: faild want: %v, got: %v", tt.name, tt.want, got)
			}
		})
	}
}

func TestExpandIntList(t *testing.T) {
	cases := map[string]struct {
		interfaceIntList []interface{}
		intList          []int
	}{
		"basic": {[]interface{}{1, 2}, []int{1, 2}},
	}
	for name, tc := range cases {
		intList := ExpandIntList(tc.interfaceIntList)

		for i, v := range intList {
			if v != tc.interfaceIntList[i].(int) {
				t.Errorf("testing %s: failed", name)
			}
		}
	}
}

func TestFlattenIntList(t *testing.T) {
	cases := map[string]struct {
		intList          []int
		interfaceIntList []interface{}
	}{
		"basic": {[]int{1, 2}, []interface{}{1, 2}},
	}
	for name, tc := range cases {
		interfaceIntList := FlattenIntList(tc.intList)

		for i, v := range interfaceIntList {
			if v.(int) != tc.intList[i] {
				t.Errorf("testing %s: failed", name)
			}
		}
	}
}

func TestRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		log.Println(RandomString(i))
	}
}

func TestStringToSlice(t *testing.T) {
	cases := map[string]struct {
		str     string
		strList []string
	}{
		"basic": {"[\"monday\", \"sunday\"]", []string{"monday", "sunday"}},
	}
	for name, tc := range cases {
		strList := StringToSlice(tc.str)

		for i, v := range strList {
			if v != tc.strList[i] {
				t.Errorf("testing %s: failed want: %s, got: %s", name, tc.strList[i], v)
			}
		}
	}
}
