use rusoto_core::Region;
use rusoto_dynamodb::DynamoDbClient;
use once_cell::sync::Lazy;

pub static DYNAMODB_CLIENT: Lazy<DynamoDbClient> = Lazy::new(|| DynamoDbClient::new(Region::ApNortheast1));

