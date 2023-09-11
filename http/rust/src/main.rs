use actix_web::{get, App, HttpServer, Responder};

#[get("/")]
async fn root() -> impl Responder {
    format!("Hello from Rust")
}

#[actix_web::main] // or #[tokio::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(root))
        .bind(("0.0.0.0", 3000))?
        .run()
        .await
}
