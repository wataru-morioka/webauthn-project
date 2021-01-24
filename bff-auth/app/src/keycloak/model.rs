use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
pub struct KeycloakTokenResponse {
    pub access_token: String, 
    pub refresh_token:String, 
    pub id_token:String,
    pub token_type: String, 
    pub expires_in: u64,
    pub refresh_expires_in: u64,
    #[serde(rename = "not-before-policy")]
    pub not_before_policy: i32,
    pub session_state: String
}