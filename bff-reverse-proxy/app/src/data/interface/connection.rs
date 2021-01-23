pub struct Connection {};

pub trait ConnectionInterface {
    static DYNAMODB_CLIENT: Lazy<DynamoDbClient>;
    pub fn generate_api_client() -> Client;
}