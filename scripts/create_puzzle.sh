#!/bin/bash

function createDay() {
    local year
    local day
    local dir
    local session

    year="$1"
    day="$2"
    session="$1"

    dir="./${year}/$(printf "%02d" "${day}")"

    if [ -d "${dir}" ]; then
      echo "directory already exist"
      return 1
    fi

    mkdir -p "${dir}"

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

  touch "${dir}/INPUT"
  curl --cookie "session=${session}" "https://adventofcode.com/${year}/day/${day}/input" | cat >> "${dir}/INPUT"

  pwd
  mv "$WORK_DIR"/temp/description.md "${dir}/README.md"
}

createDay "$1" "$2" "$3"
