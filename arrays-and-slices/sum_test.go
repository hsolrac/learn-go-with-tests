package main

import "testing"

func TestSum(t *testing.T) {
    numerbs := [5]int{1,2,3,4,5}

    got := Sum(numerbs)
    want := 15

    if got != want {
        t.Errorf("got %d want %d given, %v", got, want, numerbs)
    }
}
