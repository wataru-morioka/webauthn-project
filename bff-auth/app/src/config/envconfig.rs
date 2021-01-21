use envconfig::Envconfig;
use once_cell::sync::Lazy;
use env_logger;
use std::io::Write;
use crate::common::util;

#[derive(Envconfig, Debug)]
pub struct Env {
    #[envconfig(from = "COGNITO_CLIENTID")]
    pub cognito_clientid: String,
    #[envconfig(from = "COGNITO_CLIENTSECRET")]
    pub cognito_clientsecret: String,
    #[envconfig(from = "KEYCLOAK_CLIENTID")]
    pub keycloak_clientid: String,
    #[envconfig(from = "KEYCLOAK_CLIENTSECRET")]
    pub keycloak_clientsecret: String,
    #[envconfig(from = "RUST_LOG")]
    pub rust_log: String,
    #[envconfig(from = "FRONTEND_ENDPOINT")]
    pub frontend_endpoint: String,
    #[envconfig(from = "COGNITO_AUTHORIZATION_ENDPOINT")]
    pub cognito_authorization_endpoint: String,
    #[envconfig(from = "COGNITO_CALLBACK_URI")]
    pub cognito_callback_uri: String,
    #[envconfig(from = "COGNITO_TOKEN_ENDPOINT")]
    pub cognito_token_endpoint: String,
    #[envconfig(from = "KEYCLOAK_AUTHORIZATION_ENDPOINT")]
    pub keycloak_authorization_endpoint: String,
    #[envconfig(from = "KEYCLOAK_CALLBACK_URI")]
    pub keycloak_callback_uri: String,
    #[envconfig(from = "KEYCLOAK_TOKEN_ENDPOINT")]
    pub keycloak_token_endpoint: String,
}

pub static ENV: Lazy<Env> = Lazy::new(|| util::get_env());

pub fn initialize_logger() {
    env_logger::Builder::from_default_env()
    .format(|buf, record| {
        let ts = buf.timestamp();
        writeln!(
            buf,
            "[{} {} {}] {} {}:{}",
            ts,
            record.level(),
            record.target(),
            record.args(),
            record.file().unwrap_or("unknown"),
            record.line().unwrap_or(0),
        )
    })
    .init();
}