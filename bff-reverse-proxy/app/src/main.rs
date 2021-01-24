use log::info;
// use envconfig::Envconfig;
use std::convert::Infallible;
use hyper::server::conn::AddrStream;
use hyper::{Body, Request, Server};
use hyper::service::{service_fn, make_service_fn};
use app::config::envconfig;
use app::config::envconfig::ENV;
use app::service::interface::reverse_proxy::*;

type GenericError = Box<dyn std::error::Error + Send + Sync>;
type Result<T> = std::result::Result<T, GenericError>;

#[tokio::main]
async fn main() -> Result<()> {
    envconfig::initialize_logger();

    let addr = ([0, 0, 0, 0], 8080).into();

    let proxy_service = make_service_fn(|socket: &AddrStream| {
        let remote_addr = socket.remote_addr();
        async move {
            Ok::<_, Infallible>(service_fn(move |req: Request<Body>| {
                ReverseProxy::call(remote_addr.ip(), &ENV.api_endpont, req)
            }))
        }
    });

    let reverse_proxy_server = Server::bind(&addr).serve(proxy_service);

    info!("reverse proxy server start!");

    if let Err(e) = reverse_proxy_server.await {
        eprintln!("server error: {}", e);
    }

    Ok(())
}
