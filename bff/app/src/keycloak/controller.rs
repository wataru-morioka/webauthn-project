use actix_web::{web, get, post, delete, put, App, HttpRequest, HttpResponse, Responder, http::header};
use log::{info, error};
use crate::config::constant::Const;
use crate::config::envconfig::ENV;
use super::model::KeycloakTokenResponse;
use crate::data::interface::repository::*;
use crate::data::entity::SessionInfo;
use super::interface::service::*;
use crate::common::interface::service::*;
use crate::common::model::CallbackRequestBody;

#[get("/management/login")]
async fn management_login(req: HttpRequest) -> impl Responder {
    // TODO logger
    let (cookie, session, code_challenge) = match CommonService::generate_session(
        req, 
        &Const::KEYCLOAK_SESSION_MANAGE
    ).await {
        Ok(session) => session,
        Err(()) => {
            return HttpResponse::InternalServerError().finish();
        }  
    };

    // TODO nonce
    let request_uri = KeycloakService::generate_authz_req_uri(&session.state, &code_challenge);

    HttpResponse::TemporaryRedirect()
    .header(header::LOCATION, request_uri)
    .cookie(cookie)
    .finish()
}

#[get("/management/oauth2/callback")]
async fn management_oauth2_callback(req: HttpRequest, query: web::Query<CallbackRequestBody>) -> impl Responder {
    let mut session: SessionInfo = match CommonService::is_callback_valid(
        &req, 
        &query,
        Const::KEYCLOAK_SESSION_MANAGE
    ).await {
        Ok(session) => session,
        Err(()) => {
            return HttpResponse::Forbidden().finish();
        }  
    };
 
    let request_body = KeycloakService::generate_token_req_body(&query.code.as_ref().unwrap(), &session);
    let token_res: KeycloakTokenResponse = match ApiRepository::token_request::<KeycloakTokenResponse>(
        ENV.keycloak_token_endpoint.clone(),
        &request_body.to_vec(),
        ENV.keycloak_clientid.clone(),
        ENV.keycloak_clientsecret.clone(),
    ).await {
        Ok(res) => res,
        Err(()) => {
            return HttpResponse::InternalServerError().finish();
        }
    };

    info!("get access_token!!: {}", token_res.access_token);
    KeycloakService::set_token_info(&mut session, token_res);

    if let Err(()) = DynamoDbRepository::put_session(&session, Const::KEYCLOAK_SESSION_MANAGE).await {
        return HttpResponse::InternalServerError().finish();
    }

    HttpResponse::TemporaryRedirect()
    .header(header::LOCATION, format!("{}/home", ENV.frontend_endpoint))
    .finish()
}