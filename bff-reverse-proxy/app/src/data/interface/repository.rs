use async_trait::async_trait;
use crate::data::interface::model::SessionInfo;
use reqwest::Response;
use crate::config::errconfig::ValidationError;
use super::model::CognitoTokenResponse;

pub struct DynamoDbRepository {}

#[async_trait]
pub trait DynamoDbInterface {
    async fn put_session(session: &SessionInfo, table_name: &str) -> Result<(), ()>;
    async fn delete_session(session_id: String, table_name: &str) -> Result<(), ()>;
    async fn get_session(session_id: String, table_name: &str) -> Result<SessionInfo, ()>;
}

pub struct ApiRepository {}

#[async_trait]
pub trait ApiInterface {
    async fn refresh_token_request(params: &Vec<(&str, &str)>) -> Result<CognitoTokenResponse, ValidationError>;
}

