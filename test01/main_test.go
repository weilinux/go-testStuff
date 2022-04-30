package main

import "testing"

//multiple input, can run each at a test
func TestCalculate(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		// TODO: Add test cases.
		{
			name: "exp01",
			args: args{x: 5},
			wantResult: 7,
		},
		//2nd test fails with 20001, 20002
		{
			name: "exp02",
			args: args{x: 20000},
			wantResult: 20002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Calculate(tt.args.x); gotResult != tt.wantResult {
				t.Errorf("Calculate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

//single input
func TestCalculate2(t *testing.T) {
	if Calculate(8) != 10 {
		t.Errorf("Expected 10, got %v", Calculate(8))
	}
}

//multiple inputs
func TestTableCalculate(t *testing.T) {
	var tests = []struct{
		input int
		expected int
	}{
		{2,4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received {}", test.input, test.expected, output)
		}
	}

}