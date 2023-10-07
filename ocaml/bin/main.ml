let json_string = {|
  {"message" : "Hello from OCaml"}|}

let () =
  Dream.run
  @@ Dream.logger
  @@ Dream.router [
    Dream.get "/" (fun _ ->
      let json = Yojson.Safe.from_string json_string in
      let html_string = Yojson.Safe.to_string json in
      Dream.html html_string);
  ]
