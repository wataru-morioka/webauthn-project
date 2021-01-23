use async_trait::async_trait;

#[async_trait]
pub struct ReverseProxy {};

pub trait ReverseProxyInterface{
    async fn call(client_ip: IpAddr, forward_uri: &str, request: Request<Body>) -> Result<Response<Body>>;
}