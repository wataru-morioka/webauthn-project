use log::info;
use std::convert::Infallible;
use hyper::server::conn::AddrStream;
use hyper::{Body, Request, Response, Server};
use hyper::service::{service_fn, make_service_fn};
use app::reverse_proxy;

#[tokio::main]
async fn main() -> Result<(), ()> {
    // envconfig::initialize_logger();

    let addr = ([0, 0, 0, 0], 8080).into();

    info!("server start!");

    let make_svc = make_service_fn(|socket: &AddrStream| {
        let remote_addr = socket.remote_addr();
        async move {
            Ok::<_, Infallible>(service_fn(move |req: Request<Body>| {
                reverse_proxy::call(remote_addr.ip(), "http://backend-app-service:8000", req)
            }))
        }
    });

    info!("server start!");

    let reverse_proxy_server = Server::bind(&addr).serve(make_svc);

    info!("server start!");

    if let Err(e) = reverse_proxy_server.await {
        eprintln!("server error: {}", e);
    }

    Ok(())
}