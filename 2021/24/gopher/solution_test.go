package main

import (
  "testing"
)

func TestPuzzle(t *testing.T){
  tests := []struct{
    input []int
    expected int
  }{
    {
      []int{1},
      1,
    },
  }

  for _, test := range tests {
    t.Run("", func(t *testing.T) {
      result := Puzzle()

      if result != test.expected {
        t.Errorf("expected %6d, result %6d\n", test.expected, result)
      }
    })
  }

