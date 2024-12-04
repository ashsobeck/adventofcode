import gleam/int
import gleam/io
import gleam/list
import gleam/option
import gleam/regex
import gleam/result
import gleam/string
import simplifile

pub fn main() {
  io.println("Hello from day3!")
  let assert Ok(input) = simplifile.read("day3.txt")
  let mults =
    find_mults(input)
    |> parse_mults
    |> int.to_string
  io.println("Mults found and added: " <> mults)
}

fn find_mults(input: String) -> List(List(regex.Match)) {
  input
  |> string.trim
  |> string.split("\n")
  |> list.map(fn(s) {
    let assert Ok(re) = regex.from_string("mul\\((\\d+),(\\d+)\\)")
    regex.scan(re, s)
  })
}

fn parse_mults(mults: List(List(regex.Match))) -> Int {
  mults
  |> list.fold(0, fn(acc, mem) {
    mem
    |> list.fold(0, fn(a, m) {
      case m.submatches |> option.values {
        [l, r] ->
          a + { [int.parse(l), int.parse(r)] |> result.values |> int.product }
        _ -> 0
      }
    })
    |> int.add(acc)
  })
}
