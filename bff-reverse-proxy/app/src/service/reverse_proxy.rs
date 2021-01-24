use hyper::header::{HeaderMap, HeaderValue};
use hyper::{Body, Client, Request, Response, Uri, StatusCode};
use hyper::*;
use once_cell::sync::Lazy;
use std::net::IpAddr;
use std::str::FromStr;
use log::info;
use async_trait::async_trait;
use crate::service::interface::reverse_proxy::*;
use crate::service::interface::validation::*;
use crate::data::interface::model::SessionInfo;
use super::util;

#[async_trait]
impl ReverseProxyInterface for ReverseProxy {
    async fn call(
        client_ip: IpAddr,
        forward_uri: &str,
        request: Request<Body>,
    ) -> Result<Response<Body>> {
        info!("セッションチェック");
        let mut session: SessionInfo = match Validation::is_session_valid(&request).await {
            Ok(session) => session,
            Err(_) => {
                return Ok(util::create_response(StatusCode::FORBIDDEN).unwrap())
            }
        };

        info!("トークンチェック");
        let access_token: String = match Validation::is_access_token_valid(&mut session).await {
            Ok(access_token) => access_token,
            Err(_) => {
                return Ok(util::create_response(StatusCode::FORBIDDEN).unwrap())
            }
        };

        let proxied_request = create_proxied_request(client_ip, &forward_uri, request, access_token).unwrap();
        let client = Client::new();
        let response = client.request(proxied_request).await.unwrap();
        let proxied_response = create_proxied_response(response);
        Ok(proxied_response)
    }
}

fn is_hop_header(name: &str) -> bool {
    use unicase::Ascii;

    static HOP_HEADERS: Lazy<Vec<Ascii<&'static str>>> = Lazy::new(|| vec![
        Ascii::new("Connection"),
        Ascii::new("Keep-Alive"),
        Ascii::new("Proxy-Authenticate"),
        Ascii::new("Proxy-Authorization"),
        Ascii::new("Te"),
        Ascii::new("Trailers"),
        Ascii::new("Transfer-Encoding"),
        Ascii::new("Upgrade"),
    ]);

    HOP_HEADERS.iter().any(|h| h == &name)
}

fn remove_hop_headers(headers: &HeaderMap<HeaderValue>) -> HeaderMap<HeaderValue> {
    let mut result = HeaderMap::new();
    for (k, v) in headers.iter() {
        if !is_hop_header(k.as_str()) {
            result.insert(k.clone(), v.clone());
        }
    }
    result
}

fn create_proxied_response<B>(mut response: Response<B>) -> Response<B> {
    *response.headers_mut() = remove_hop_headers(response.headers());
    response
}

fn forward_uri<B>(forward_url: &str, req: &Request<B>) -> Uri {
    let forward_uri = match req.uri().query() {
        Some(query) => format!("{}{}?{}", forward_url, req.uri().path(), query),
        None => format!("{}{}", forward_url, req.uri().path()),
    };

    hyper::Uri::from_str(forward_uri.as_str()).unwrap()
}

fn create_proxied_request<B>(
    client_ip: IpAddr,
    forward_url: &str,
    mut request: Request<B>,
    access_token: String
) -> Result<Request<B>> {
    *request.headers_mut() = remove_hop_headers(request.headers());
    let bearer_header = format!("Bearer: {}", access_token);
    request.headers_mut().insert("Authorization", HeaderValue::from_str(&bearer_header).unwrap());
    *request.uri_mut() = forward_uri(forward_url, &request);

    info!("proxy to: {}", forward_url);

    let x_forwarded_for_header_name = "x-forwarded-for";

    match request.headers_mut().entry(x_forwarded_for_header_name) {
        hyper::header::Entry::Vacant(entry) => {
            entry.insert(client_ip.to_string().parse().unwrap());
        }

        hyper::header::Entry::Occupied(mut entry) => {
            let addr = format!("{}, {}", entry.get().to_str().unwrap(), client_ip);
            entry.insert(addr.parse().unwrap());
        }
    }

    Ok(request)
}



