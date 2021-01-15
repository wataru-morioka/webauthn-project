use crate::config::envconfig::ENV;
use super::interface::service::*;
use crate::data::entity::SessionInfo;

impl CognitoInterface for CognitoService {
    fn generate_authz_req_uri(state: &str, code_challenge: &str) -> String {
        format!(
            "{}?response_type={}&client_id={}&redirect_uri={}&state={}&scope={}&code_challenge_method={}&code_challenge={}"
            , ENV.cognito_authorization_endpoint
            , "code"
            , ENV.cognito_clientid
            , ENV.cognito_callback_uri
            , state
            , "email+openid"
            , "S256"
            , code_challenge
        )
    }

    fn generate_token_req_body<'a>(authorization_code: &'a str, session: &'a SessionInfo) -> Vec<(&'a str, &'a str)> {
        let request_body = [
            ("grant_type", "authorization_code"),
            ("client_id", &ENV.cognito_clientid),
            ("redirect_uri", &ENV.cognito_callback_uri),
            ("code", authorization_code),
            ("code_verifier", &session.code_verifier),
        ];
        request_body.to_vec()
    }
}

