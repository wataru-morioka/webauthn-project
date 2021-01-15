use actix_web::{web, HttpRequest, HttpMessage, http::Cookie};
use async_trait::async_trait;
use log::{info, error};
use super::model::CallbackRequestBody;
use super::interface::service::*;
use crate::data::entity::SessionInfo;
use crate::config::constant::Const;
use crate::data::interface::repository::*;
use crate::common::util;

#[async_trait(?Send)]
impl CommonInterface for CommonService {
    async fn generate_session(req: HttpRequest, table_name: &str)
    -> Result<(Cookie<'static>, SessionInfo, String), ()> {
        if let Some(cookie) = req.cookie(Const::AUTH_COOKIE) {
            println!("ログイン時クッキーあり: {}", cookie);
            if let Err(()) = DynamoDbRepository::delete_session(
                cookie.value().to_string(),
                table_name
            ).await {
                return Err(());
            };
        }
    
        let cookie: Cookie = util::generate_cookie();
        info!("generate cookie: {}", cookie);
        
        let state: String = util::generate_base64_url();
        info!("generate state: {}", state);
    
        let code_verifier : String = util::generate_base64_url();
        let code_challenge: String = util::generate_code_challenge(&code_verifier);
        info!("generate code_verifier: {}", code_verifier);
        info!("generate code_challenge: {}", code_challenge);
    
        let session = SessionInfo::new(
            cookie.value().to_string(),
            state,
            code_verifier,
            util::get_current_unix_timestamp(),
        );
        if let Err(()) = DynamoDbRepository::put_session(&session, table_name).await {
            return Err(());
        }
        Ok((cookie, session, code_challenge))
    }
    
    async fn is_callback_valid(req: &HttpRequest, query: &web::Query<CallbackRequestBody>, table_name: &str)
    -> Result<SessionInfo, ()> {
        let session_id: String = match req.cookie(Const::AUTH_COOKIE) {
            Some(cookie) => {
                info!("クッキー受信 cookie: {}", cookie.value());
                cookie.value().to_string()
            },
            None => {
                error!("クッキーなし");
                return Err(());
            }
        };
        info!("get session_id! : {}", &session_id);
    
        let authorization_code: String = match query.code.clone() {
            Some(code) => code,
            None => {
                error!("認可コード不正");
                return Err(());
            }
        };
        info!("get grantcode! : {}", &authorization_code);
    
        let session: SessionInfo = match DynamoDbRepository::get_session(session_id, table_name).await {
            Ok(session) => session,
            Err(()) => {
                return Err(());
            }
        };
    
        info!("get server state!! : {}", session.state);
        if query.state.is_none() || query.state != Some(session.state.clone()) {
            println!("state不正");
            return Err(());
        }
        info!("get state!! : {}", query.state.clone().unwrap());
    
        Ok(session)
    }
}

