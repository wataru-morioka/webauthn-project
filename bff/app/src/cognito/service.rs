use crate::config::envconfig::ENV;
use super::interface::service::*;
use crate::data::entity::SessionInfo;
use super::model::CognitoTokenResponse;

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

    fn set_token_info(session: &mut SessionInfo, token_res: CognitoTokenResponse) {
        session.access_token = Some(token_res.access_token);
        session.refresh_token = Some(token_res.refresh_token);
        session.token_expires_in = Some(token_res.expires_in);
    }
}

