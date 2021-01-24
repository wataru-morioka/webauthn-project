use time::*;
use std::time::SystemTime;
use hyper::{Body, Response, StatusCode};

pub fn create_response(status_code: StatusCode) -> Result<Response<Body>> {
    Ok(Response::builder().status(status_code).body(Body::empty()).unwrap())
}

pub fn get_current_unix_timestamp() -> u64 {
    SystemTime::now()
    .duration_since(SystemTime::UNIX_EPOCH)
    .unwrap()
    .as_secs()
}