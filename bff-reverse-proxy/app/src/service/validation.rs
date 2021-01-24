use hyper::{Request, Body};
use hyper::header::{HeaderValue};
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
        let cookie: &HeaderValue = req.headers()
            .get_all(Const::COOKIE)
            .iter()
            .find(|x| x.to_str().unwrap().starts_with(Const::AUTH_COOKIE))
            .ok_or(ValidationError::SessionError)?;

        info!("get cookie!!: {}", cookie.to_str().unwrap());

        DynamoDbRepository::get_session(
            cookie.to_str().unwrap().to_string(), 
            Const::COGNITO_SESSION_MANAGE
        ).await
        .map_err(|_| ValidationError::SessionError)
        .and_then(|session| Ok(session))
    }

    async fn is_access_token_valid(session: &mut SessionInfo) -> Result<String, ValidationError> {
        let expire: u64 = session.created_timestamp + session.token_expires_in.unwrap();
        let now: u64 = util::get_current_unix_timestamp();
        if now < expire {
            return Ok(session.access_token.unwrap())
        }

        match ApiRepository::token_request().await {
            Ok(res) => {
                session.access_token = res.access_token;
                session.refresh_token = res.refresh_token;
                session.created_timestamp = now;
                match DynamoDbRepository::put_session(
                    session, 
                    Const::COGNITO_SESSION_MANAGE
                ).await {
                    Ok(_) => {
                        return Ok(res.access_token)
                    },
                    Err(()) => {
                        return Err(())
                    }
                }
            },
            Err(()) => {
                return Err(())
            }
        };
    }
}