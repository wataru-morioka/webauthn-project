use hyper::{Request, Body};
use log::info;
use async_trait::async_trait;
use crate::config::constant::Const;
use crate::data::interface::model::SessionInfo;
use crate::service::interface::validation::*;
use crate::data::interface::repository::*;
use super::util;
use crate::config::errconfig::ValidationError;

#[async_trait]
impl ValidationInterface for Validation {
    async fn is_session_valid(req : &Request<Body>) -> Result<SessionInfo, ValidationError> {
        let cookie: String = req.headers()
            .get_all(Const::COOKIE)
            .iter()
            .find(|x| x.to_str().unwrap().starts_with(Const::AUTH_COOKIE))
            .ok_or(ValidationError::SessionError)?
            .to_str()
            .unwrap()
            .to_string();

        info!("get cookie!!: {}", cookie);

        DynamoDbRepository::get_session(
            util::get_cookie_value(cookie), 
            Const::COGNITO_SESSION_MANAGE
        ).await
        .map_err(|_| ValidationError::SessionError)
    }

    async fn is_access_token_valid(session: &mut SessionInfo) -> Result<String, ValidationError> {
        let access_token = is_expiratioin_valid(session);
        if access_token.is_ok() {
            // return access_token;
        }

        info!("リフレッシュトークン使用");

        let params = util::create_token_params(&session);
        let res = ApiRepository::token_request(&params).await?;

        info!("res: {:?}", res);

        access_token
    }
}

fn is_expiratioin_valid(session: &mut SessionInfo) -> Result<String, ValidationError> {
    let expire: u64 = session.created_timestamp + session.token_expires_in.unwrap();
    let now: u64 = util::get_current_unix_timestamp();
    if now < expire {
        Ok(session.access_token.clone().unwrap())
    } else {
        Err(ValidationError::AccessTokenExpireError)
    }
}
