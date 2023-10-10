[@@@warning "-37"]
[@@@warning "-32"]
[@@@warning "-69"]
[@@@warning "-8"]

open Riot

let port = 3000
let acceptors = 20

type response = {
  message: string;
}
[@@deriving yojson]

let main () =
  Riot.Logger.set_log_level (Some Debug);
  Server.Logger.set_log_level (Some Debug);
  Socket.Logger.set_log_level (Some Debug);

  Riot.Logger.start () |> Result.get_ok;

  Logger.info (fun f -> f "Starting server on port %d" port);

  let (Ok _server) =
    Http_server.start_link ~port ~acceptors @@ fun reqd ->
    let req = Httpaf.Reqd.request reqd in
    Logger.debug (fun f -> f "request: %a" Httpaf.Request.pp_hum req);
    if req.target = "/json" then
      let json = response_to_yojson {message = "Hello from Ocaml" } 
        |> Yojson.Safe.to_string in
      let json_length = String.length(json)
        |> string_of_int in
      let headers = Httpaf.Headers.of_list [("content-type", "application/json"); ("content-length", json_length) ] in
      let res = Httpaf.Response.create ~headers `OK in
      Httpaf.Reqd.respond_with_string reqd res json;
      Logger.debug (fun f -> f "response: %a" Httpaf.Response.pp_hum res)
  in

  let rec loop () =
    yield ();
    loop ()
  in
  loop ()

let () = Riot.run @@ main
