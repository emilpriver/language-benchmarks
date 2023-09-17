import { Elysia } from "elysia";

const app = new Elysia()
  .get("/json", () => {
    const p = JSON.stringify({
      message: "Hello from Bun"
    })

    return new Response(p, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
  })
  .listen(3000);


console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
