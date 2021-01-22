use envconfig::Envconfig;
use once_cell::sync::Lazy;
use env_logger;
use std::io::Write;
use crate::util;

#[derive(Envconfig, Debug)]
pub struct Env {
    #[envconfig(from = "API_ENDPOINT")]
    pub api_endpont: String,
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