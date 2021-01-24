use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use crate::data::interface::model::SessionInfo;

pub struct DynamoDbRepository {}

#[async_trait]
pub trait DynamoDbInterface {
    async fn put_session(session: &SessionInfo, table_name: &str) -> Result<(), ()>;
    async fn delete_session(session_id: String, table_name: &str) -> Result<(), ()>;
    async fn get_session(session_id: String, table_name: &str) -> Result<SessionInfo, ()>;
}

pub struct ApiRepository {}

#[async_trait(?Send)]
pub trait ApiInterface {
    async fn token_request<T>(
        endpoint: String, 
        request_body: &Vec<(&str, &str)>,
        client_id: String,
        client_secret: String
    ) -> Result<T, ()> 
    where T : for<'a> Deserialize<'a>;
}

