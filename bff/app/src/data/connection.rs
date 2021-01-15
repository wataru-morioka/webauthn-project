use rusoto_core::Region;
use rusoto_dynamodb::DynamoDbClient;
use once_cell::sync::Lazy;
use actix_web::client::{Client, Connector};
use openssl::ssl::{SslConnector, SslMethod};

pub static DYNAMODB_CLIENT: Lazy<DynamoDbClient> = Lazy::new(|| DynamoDbClient::new(Region::ApNortheast1));

pub fn generate_api_client() -> Client {
    let builder = SslConnector::builder(SslMethod::tls()).unwrap();
    Client::builder()
    .connector(Connector::new().ssl(builder.build()).finish())
    .finish()
}