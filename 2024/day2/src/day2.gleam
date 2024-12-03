import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile

type Direction {
  Increasing
  Decreasing
  Establishing
  Unsafe
}

pub fn main() {
  io.println("Hello from day2!")
  let assert Ok(input) = simplifile.read("day2.txt")
  let safe_reports =
    parse_reports(input)
    |> parse_levels
    |> int.to_string
  let fault_tolerant_safe_reports =
    parse_reports(input)
    |> parse_levels_tolerant
    |> int.to_string
  io.println("Safe reports are " <> safe_reports)
  io.println("Fault tolerant safe reports are " <> fault_tolerant_safe_reports)
}

fn parse_reports(input: String) -> List(List(Int)) {
  input
  |> string.trim
  |> string.split("\n")
  |> list.map(fn(s) {
    s
    |> string.split(" ")
    |> list.map(fn(i) {
      let assert Ok(level) = int.parse(i)
      level
    })
  })
}

fn parse_levels(report: List(List(Int))) -> Int {
  report
  |> list.map(fn(levels) {
    levels
    |> list.fold_until(#(0, Establishing), fn(prev, level) {
      case is_safe(level, prev) {
        #(True, dir) -> list.Continue(#(level, dir))
        #(False, _) -> list.Stop(#(-1, Unsafe))
      }
    })
  })
  |> list.count(fn(i) { i.0 != -1 })
}

fn parse_levels_tolerant(report: List(List(Int))) -> Int {
  report
  |> list.map(fn(levels) {
    levels
    |> list.fold_until(#(0, Establishing, False), fn(prev, level) {
      case is_safe_tolerant(level, prev) {
        #(True, dir, retried) -> list.Continue(#(level, dir, retried))
        #(False, _, True) -> list.Continue(#(prev.0, prev.1, True))
        #(False, _, False) -> list.Stop(#(-1, Unsafe, True))
      }
    })
  })
  |> list.count(fn(i) { i.0 != -1 })
}

fn is_safe_tolerant(
  curr: Int,
  prev: #(Int, Direction, Bool),
) -> #(Bool, Direction, Bool) {
  let #(prev_val, direction, retry) = prev
  let diff = curr - prev_val
  let new_dir = case diff > 0 {
    True -> Increasing
    False -> Decreasing
  }
  let abs_diff = int.absolute_value(diff)
  let safe =
    abs_diff <= 3
    && abs_diff >= 1
    && { direction == new_dir || direction == Establishing }
  case prev {
    #(0, _, _) -> #(True, Establishing, retry)
    _ ->
      case safe {
        True -> #(safe, new_dir, retry)
        False -> #(safe, new_dir, !retry)
      }
  }
}

fn is_safe(curr: Int, prev: #(Int, Direction)) -> #(Bool, Direction) {
  let #(prev_val, direction) = prev
  let diff = curr - prev_val
  let new_dir = case diff > 0 {
    True -> Increasing
    False -> Decreasing
  }
  let abs_diff = int.absolute_value(diff)
  let safe =
    abs_diff <= 3
    && abs_diff >= 1
    && { direction == new_dir || direction == Establishing }
  case prev {
    #(0, _) -> #(True, Establishing)
    _ -> #(safe, new_dir)
  }
}
