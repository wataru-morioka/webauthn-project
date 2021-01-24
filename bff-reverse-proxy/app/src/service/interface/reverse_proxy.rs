use async_trait::async_trait;
use hyper::*;
use std::net::IpAddr;

pub struct ReverseProxy {}

#[async_trait]
pub trait ReverseProxyInterface{
    async fn call(client_ip: IpAddr, forward_uri: &str, request: Request<Body>) -> Result<Response<Body>>;
}