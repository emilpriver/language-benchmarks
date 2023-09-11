var t=Bun.serve({port:3000,fetch(c){return new Response("Welcome to Bun!")}});console.log(`Listening on localhost:${t.port}`);
