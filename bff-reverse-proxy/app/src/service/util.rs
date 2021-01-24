use time::*;
use base64;
use std::time::SystemTime;
use hyper::{Body, Response, StatusCode};
use crate::config::envconfig::ENV;
use crate::data::interface::model::SessionInfo;
use crate::config::constant::Const;

pub fn create_response(status_code: StatusCode) -> Result<Response<Body>> {
    Ok(Response::builder().status(status_code).body(Body::empty()).unwrap())
}

pub fn get_current_unix_timestamp() -> u64 {
    SystemTime::now()
    .duration_since(SystemTime::UNIX_EPOCH)
    .unwrap()
    .as_secs()
}

pub fn generate_basic_header() -> String {
    let basic = format!("Basic {}:{}", &ENV.cognito_clientid, &ENV.cognito_clientsecret);
    base64::encode(basic)
}

pub fn create_token_params<'a>(session: &'a SessionInfo) -> Vec<(&'a str, &'a str)> {
    let request_body = [
        ("grant_type", "authorization_code"),
        ("client_id", &ENV.cognito_clientid),
        ("refresh_token", &session.refresh_token.as_ref().unwrap()),
    ];
    request_body.to_vec()
}

pub fn get_cookie_value(cookie: String) -> String {
    let mut prefix = String::from(Const::AUTH_COOKIE);
    prefix.push('=');
    cookie.replacen(&prefix, "", 1)
}