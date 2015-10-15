package test1

import "testing"


func sum(a, b int) int {
    return a + b;
}

func TestSum(t *testing.T) {
    if sum(1,23) != 24 {
        t.Error("sum() error")
    }
}

func TestHello(t *testing.T) {
    t.Error("sum() error", 123)
}

