open Lwt.Infix

let calculate_sum_of_squares chunk =
  Lwt_list.fold_left_s (fun acc x -> Lwt.return (acc + x * x)) 0 chunk

let main () =
  let numbers = List.init 1_000_000 (fun x -> x + 1) in
  let chunk_size = 10_000 in
  let chunks = ListUtils.chunk_list chunk_size numbers in

  Lwt_list.map_p calculate_sum_of_squares chunks
  |> Lwt.map List.fold_left (+) 0
  |> Lwt_main.run
  |> Printf.printf "Sum of squares: %d\n"

let () = main ()

