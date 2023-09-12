const server = Bun.serve({
  port: 3000,
  fetch(request) {
    const p = JSON.stringify({
      message: "Hello from Bun"
    })

    const h = new Headers()
    h.set("Content-Type", "application/json")

    return new Response(p, {
      headers: h
    });
  },
});

console.log(`Listening on localhost:${server.port}`);
