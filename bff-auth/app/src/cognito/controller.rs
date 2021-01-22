use actix_web::{web, get, post, delete, put, App, HttpRequest, HttpResponse, Responder, http::header};
use log::{info, error};
use crate::config::constant::Const;
use crate::config::envconfig::ENV;
use super::model::CognitoTokenResponse;
use crate::data::interface::repository::*;
use crate::data::entity::SessionInfo;
use super::interface::service::*;
use crate::common::interface::service::*;
use crate::common::model::CallbackRequestBody;

#[get("/login")]
pub async fn login(req: HttpRequest) -> impl Responder {
    let (cookie, session, code_challenge) = match CommonService::generate_session(
        req, 
        &Const::COGNITO_SESSION_MANAGE
    ).await {
        Ok(session) => session,
        Err(()) => {
            return HttpResponse::InternalServerError().finish();
        }  
    };

    // TODO nonce
    let request_uri = CognitoService::generate_authz_req_uri(&session.state, &code_challenge);

    HttpResponse::TemporaryRedirect()
    .header(header::LOCATION, request_uri)
    .cookie(cookie)
    .finish()
}

#[get("/oauth2/callback")]
pub async fn oauth2_callback(req: HttpRequest, query: web::Query<CallbackRequestBody>) -> impl Responder {
    let mut session: SessionInfo = match CommonService::is_callback_valid(
        &req,
        &query, 
        Const::COGNITO_SESSION_MANAGE
    ).await {
        Ok(session) => session,
        Err(()) => {
            return HttpResponse::Forbidden().finish();
        }  
    };
 
    let request_body = CognitoService::generate_token_req_body(&query.code.as_ref().unwrap(), &session);
    let token_res: CognitoTokenResponse = match ApiRepository::token_request::<CognitoTokenResponse>(
        ENV.cognito_token_endpoint.clone(),
        &request_body.to_vec(),
        ENV.cognito_clientid.clone(),
        ENV.cognito_clientsecret.clone(),
    ).await {
        Ok(res) => res,
        Err(()) => {
            return HttpResponse::InternalServerError().finish();
        }
    };

    info!("get access_token!!: {}", token_res.access_token);
    CognitoService::set_token_info(&mut session, token_res);

    if let Err(()) = DynamoDbRepository::put_session(&session, Const::COGNITO_SESSION_MANAGE).await {
        return HttpResponse::InternalServerError().finish();
    }

    HttpResponse::TemporaryRedirect()
    .header(header::LOCATION, format!("{}/home", ENV.frontend_endpoint))
    .finish()
}

