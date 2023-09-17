defmodule Server.Application do
  use Application

  def start(_type, _args) do
    Supervisor.start_link(
      [{Bandit, plug: Server, scheme: :http, port: 3000}],
      strategy: :one_for_one,
      name: Server.Supervisor
    )
  end
end
