use log::info;
// use envconfig::Envconfig;
use std::convert::Infallible;
use hyper::server::conn::AddrStream;
use hyper::{Body, Request, Server};
use hyper::service::{service_fn, make_service_fn};
use app::reverse_proxy;
use app::envconfig;
use app::envconfig::ENV;
use app::validation;

#[tokio::main]
async fn main() -> Result<(), ()> {
    envconfig::initialize_logger();

    let addr = ([0, 0, 0, 0], 8080).into();

    let proxy_service = make_service_fn(|socket: &AddrStream| {
        let remote_addr = socket.remote_addr();
        async move {
            Ok::<_, Infallible>(service_fn(move |req: Request<Body>| {
                // TODO 
                // クッキーに紐づくセッションの存在確認（DynamoDB）
                // なければ、再ログイン(ブラウザで判断)
                // アクセストークンの有効期限確認
                // 有効期限が切れていた場合、リフレッシュトークンにより新規アクセストークイン発行し、dynamoDBを更新
                // リフレッシュトークンが無効な場合、再ログイン(ブラウザで判断)
                reverse_proxy::call(remote_addr.ip(), &ENV.api_endpont, req)
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