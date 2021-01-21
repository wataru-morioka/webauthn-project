use actix_web::{App, HttpServer};
use log::info;
use app::cognito;
use app::keycloak;
use app::config::envconfig;
use app::reverse_proxy;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    envconfig::initialize_logger();

    let server = HttpServer::new(|| {
        App::new()
            .service(cognito::controller::login)
            .service(cognito::controller::oauth2_callback)
            .service(keycloak::controller::management_login)
            .service(keycloak::controller::management_oauth2_callback)
    })
    // .server_hostname("webauthn.api".to_string())
    .bind("0.0.0.0:8080")?;

    info!("server start!");

    server.run().await
}