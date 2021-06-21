package core

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		collect interface{}
		target  interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		// test cases.
		{"nil", args{nil, "b"}, false, -1},
		{"slice-string", args{[]string{"a", "b"}, "b"}, true, 1},
		{"slice-string", args{[]string{"a", "b"}, "c"}, false, -1},
		{"slice-int", args{[]int{2, 3}, 2}, true, 0},
		{"slice-int", args{[]int{2, 3}, 4}, false, -1},
		{"map-string", args{map[string]string{"a": "", "b": "b"}, "a"}, true, 0},
		{"map-string", args{map[string]string{"a": "", "b": "b"}, "b"}, true, 0},
		{"map-string", args{map[string]string{"a": "", "b": "b"}, "c"}, false, 0},
		{"string", args{"", ""}, true, 0},
		{"string", args{"hello", "he"}, true, 0},
		{"string", args{"hello", "llo"}, true, 2},
		{"string", args{"hello", "hi"}, false, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Contains(tt.args.collect, tt.args.target)
			if got != tt.want {
				t.Errorf("Contains() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Contains() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil}, true},
		{"string", args{""}, true},
		{"string", args{"a"}, false},
		{"slice", args{[]int{}}, true},
		{"slice", args{[]int{3}}, false},
		{"array", args{[2]int{}}, true},
		{"array", args{[2]int{3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.target); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}