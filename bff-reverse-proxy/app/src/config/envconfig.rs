use envconfig::Envconfig;
use once_cell::sync::Lazy;
use env_logger;
use std::io::Write;
use log::{info, error};

#[derive(Envconfig, Debug)]
pub struct Env {
    #[envconfig(from = "API_ENDPOINT")]
    pub api_endpont: String,
    #[envconfig(from = "COGNITO_CLIENTID")]
    pub cognito_clientid: String,
    #[envconfig(from = "COGNITO_CLIENTSECRET")]
    pub cognito_clientsecret: String,
    #[envconfig(from = "COGNITO_TOKEN_ENDPOINT")]
    pub cognito_token_endpoint: String,
}

pub static ENV: Lazy<Env> = Lazy::new(|| get_env());

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

pub fn get_env() -> Env {
    let env = match Env::init_from_env() {
        Ok(mut env) => {
            env.api_endpont.retain(|c| c != '\n');
            env.cognito_clientid.retain(|c| c != '\n');
            env.cognito_clientsecret.retain(|c| c != '\n');
            env.cognito_token_endpoint.retain(|c| c != '\n');
            env
        },
        Err(err) => {
            error!("{}",err);
            panic!("failed to get env");
        }
    };
    env
}