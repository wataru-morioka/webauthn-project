use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct SessionId {
    #[serde(rename = "SessionId")]
    pub session_id: String,
}

#[derive(Deserialize)]
pub struct CallbackRequestBody {
    pub code: Option<String>,
    pub state: Option<String>
}