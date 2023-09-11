use actix_web::{get, App, HttpServer, Responder};

#[get("/")]
async fn root() -> impl Responder {
    format!("Hello from Rust")
}

#[actix_web::main] // or #[tokio::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(root))
        .bind(("127.0.0.1", 3000))?
        .run()
        .await
}
