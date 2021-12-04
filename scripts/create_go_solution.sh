#!/bin/bash

function create_go_solution() {
  local day
  local year
  local dir

  year="$1"
  day="$2"

  dir="./${year}/$(printf "%02d" "${day}")/gopher"

  if [ -d "${dir}" ]; then
      echo "directory already exist"
      return 1
  fi

  mkdir -p "${dir}"

  cat <<EOF > "${dir}/go.mod"
module aoc${year}_$(printf "%02d" "${day}")

go 1.17

require (
	github.com/nozgurozturk/aoc/util/gopher v0.0.0
)
replace github.com/nozgurozturk/aoc/util/gopher v0.0.0 => ../../../util/gopher

EOF

  cat <<EOF > "${dir}/solution.go"
package main

const FileName = "INPUT"

func main(){
  Puzzle()
}

func Puzzle(){
  panic("not implemented")
}

EOF

  cat <<EOF > "${dir}/solution_test.go"
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

EOF
}

create_go_solution "$1" "$2"
