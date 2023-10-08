type response = {
  message: string;
}
[@@deriving yojson]

let () =
  Dream.run
  @@ Dream.logger
  @@ Dream.router [
    Dream.get "/" (fun _ ->
      let json = response_to_yojson {message = "Hello from Ocaml" } 
      |> Yojson.Safe.to_string in
      Dream.json json);
  ]
