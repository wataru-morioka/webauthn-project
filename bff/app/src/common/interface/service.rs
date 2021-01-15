use actix_web::{web, HttpRequest, http::Cookie};
use async_trait::async_trait;
use crate::common::model::CallbackRequestBody;
use crate::data::entity::SessionInfo;

pub struct CommonService {}

#[async_trait(?Send)]
pub trait CommonInterface {
    async fn generate_session(req: HttpRequest, table_name: &str) -> Result<(Cookie<'static>, SessionInfo, String), ()>;
    async fn is_callback_valid(req: &HttpRequest, query: &web::Query<CallbackRequestBody>, table_name: &str) -> Result<SessionInfo, ()>;
}