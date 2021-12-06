#!/bin/bash

function create_rust_solution() {
  local day
  local year
  local dir

  year="$1"
  day="$2"

  dir="./${year}/$(printf "%02d" "${day}")/ferris"

  if [ -d "${dir}" ]; then
      echo "directory already exist"
      return 1
  fi

  mkdir -p "${dir}/src"

  cat <<EOF > "${dir}/Cargo.toml"
[package]
name = "aoc${year}_$(printf "%02d" "${day}")"
version = "0.1.0"
authors = ["n.ozgurozturk <ozgur@nozgurozturk.com>"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]

EOF

  cat <<EOF > "${dir}/src/main.rs"
fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}

EOF

}

create_rust_solution "$1" "$2"
