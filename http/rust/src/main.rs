use axum::{response::Json, routing::get, Router};
use serde_json::{json, Value};

async fn json() -> Json<Value> {
    Json(json!({ "message": "Hello from Rust" }))
}

#[tokio::main]
async fn main() {
    // build our application with a single route
    let app = Router::new()
        .route("/", get(|| async { "Hello from Rust" }))
        .route("/json", get(json));

    // run it with hyper on localhost:3000
    axum::Server::bind(&"0.0.0.0:3000".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}
