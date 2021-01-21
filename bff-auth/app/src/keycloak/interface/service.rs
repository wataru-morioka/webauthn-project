use crate::data::entity::SessionInfo;
use super::super::model::KeycloakTokenResponse;

pub struct KeycloakService {}

pub trait KeycloakInterface {
    fn generate_authz_req_uri(state: &str, code_challenge: &str) -> String;
    fn generate_token_req_body<'a>(authorization_code: &'a str, session: &'a SessionInfo) -> Vec<(&'a str, &'a str)>;
    fn set_token_info(session: &mut SessionInfo, token_res: KeycloakTokenResponse);
}