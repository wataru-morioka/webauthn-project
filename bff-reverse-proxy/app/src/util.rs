// use rand::prelude::*;
// use rand_chacha::{ChaCha20Rng, rand_core::{RngCore, SeedableRng}};
// use crypto::digest::Digest;
// use crypto::sha2::Sha256;
use envconfig::Envconfig;
use crate::envconfig::Env;
// use actix_web::http::Cookie;
// use time::*;
// use std::time::SystemTime;
use log::{info, error};
// use crate::config::constant::Const;

// pub fn generate_base64_url() -> String {
//     let mut seed = Default::default();
//     thread_rng().fill(&mut seed);
//     let mut rng = ChaCha20Rng::from_seed(seed);
//     let mut bytes = [0u8; 32];
//     rng.fill_bytes(&mut bytes);
//     let base64_url: String = base64_url::encode(&bytes);
//     base64_url
// }

// pub fn generate_code_challenge(verifier: &str) -> String {
//     let mut hasher = Sha256::new();
//     hasher.input(verifier.as_bytes());
//     let mut bytes = [0u8; 32];
//     hasher.result(&mut bytes);
//     base64_url::encode(&bytes)
// }

// pub fn generate_cookie() -> Cookie<'static> {
//     Cookie::build(Const::AUTH_COOKIE, generate_base64_url())
//     // .domain("www.rust-lang.org")
//     .path("/")
//     .max_age(Duration::seconds(Const::COOKIE_EXPIRES_IN))
//     .secure(false)
//     .http_only(true)
//     .finish()
// }

pub fn get_env() -> Env {
    let env = match Env::init_from_env() {
        Ok(mut env) => {
            env.api_endpont.retain(|c| c != '\n');
            env
        },
        Err(err) => {
            error!("{}",err);
            panic!("failed to get env");
        }
    };
    env
}