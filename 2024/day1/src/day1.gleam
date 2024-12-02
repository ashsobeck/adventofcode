import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub fn main() {
  let file_name = "day1.txt"
  let assert Ok(input) = simplifile.read(from: file_name)
  let input_list =
    input
    |> string.split("\n")
    |> list.filter(fn(s) { s != "" })
  let #(map_one, map_two) = get_lists(input_list)
  io.println(parse_lists(map_one, map_two))
  io.println(parse_similarity(map_one, map_two))
}

fn parse_similarity(map_one: List(Int), map_two: List(Int)) -> String {
  let similarity =
    map_one
    |> list.map(fn(x) { x * list.count(map_two, fn(y) { y == x }) })
    |> list.reduce(fn(acc, x) { acc + x })

  case similarity {
    Ok(s) -> "Similarity is " <> int.to_string(s)
    _ -> "Error"
  }
}

fn parse_lists(map_one: List(Int), map_two: List(Int)) -> String {
  let distance =
    list.map2(map_one, map_two, fn(x, y) { int.absolute_value(x - y) })
    |> list.reduce(fn(acc, x) { acc + x })

  case distance {
    Ok(d) -> "Distance is " <> int.to_string(d)
    _ -> "Error"
  }
}

fn get_lists(input_list: List(String)) -> #(List(Int), List(Int)) {
  let separated_maps =
    input_list
    |> list.map(fn(e) { string.split(e, "   ") })
    |> list.flatten
    |> list.index_map(fn(e, i) {
      case int.is_odd(i) {
        True -> Error(int.parse(e) |> result.unwrap(0))
        _ -> Ok(int.parse(e) |> result.unwrap(0))
      }
    })

  let map_a =
    separated_maps
    |> list.filter(fn(e) {
      case e {
        Ok(_) -> True
        _ -> False
      }
    })
    |> list.map(fn(e) { result.unwrap(e, 0) })
    |> list.sort(int.compare)

  let map_b =
    separated_maps
    |> list.filter(fn(e) {
      case e {
        Error(_) -> True
        _ -> False
      }
    })
    |> list.map(fn(e) {
      case e {
        Error(e) -> e
        _ -> -1
      }
    })
    |> list.sort(int.compare)
  #(map_a, map_b)
}
