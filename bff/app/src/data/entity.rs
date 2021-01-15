use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct SessionInfo {
    #[serde(rename = "SessionId")]
    pub session_id: String,
    #[serde(rename = "AccessToken")]
    pub access_token: Option<String>,
    #[serde(rename = "ExpiresIn")]
    pub expires_in: Option<i32>,
    #[serde(rename = "RefreshToken")]
    pub refresh_token: Option<String>,
    #[serde(rename = "State")]
    pub state: String,
    #[serde(rename = "Nonce")]
    pub nonce: Option<String>,
    #[serde(rename = "CodeVerifier")]
    pub code_verifier: String,
    #[serde(rename = "CreatedTimestamp")]
    pub created_timestamp: u64,
}

impl SessionInfo {
    pub fn new(session_id: String, state: String, code_verifier: String, created_timestamp: u64) -> SessionInfo {
        SessionInfo {
            session_id,
            access_token: None,
            expires_in: None,
            refresh_token: None,
            state,
            nonce: None,
            code_verifier,
            created_timestamp,
        }
    }
}