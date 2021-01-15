use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
pub struct CognitoTokenResponse {
    pub access_token: String, 
    pub refresh_token:String, 
    pub id_token:String,
    pub token_type: String, 
    pub expires_in: i32,
}

