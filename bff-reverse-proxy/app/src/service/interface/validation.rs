use hyper::{Request, Body};
use async_trait::async_trait;
use crate::data::interface::model::SessionInfo;
use crate::config::errconfig::ValidationError;

pub struct Validation {}

#[async_trait]
pub trait ValidationInterface{
    async fn is_session_valid(req : &Request<Body>) -> Result<SessionInfo, ValidationError>;
    async fn is_access_token_valid(session: &mut SessionInfo) -> Result<String, ValidationError>;
}