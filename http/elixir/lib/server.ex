defmodule Server do
  import Plug.Conn

  def init(opts), do: opts

  def call(%{path_info: ["json"]} = conn, _opts) do
    conn
    |> put_resp_header("content-type", "application/json")
    |> send_resp(200, Jason.encode!(%{message: "Hello from Elixir"}))
  end

  def call(%{path_info: []} = conn, _opts) do
    send_resp(conn, 200, "Hello from Elixir")
  end

  def call(conn, _opts) do
    send_resp(conn, 204, "")
  end
end
