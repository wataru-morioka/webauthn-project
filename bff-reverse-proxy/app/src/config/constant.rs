pub struct Const;

impl Const {
    pub const AUTH_COOKIE: &'static str = "auth_session";
    pub const COOKIE: &'static str = "Cookie";
    pub const COGNITO_SESSION_MANAGE: &'static str = "CognitoSessionManage";
    pub const KEYCLOAK_SESSION_MANAGE: &'static str = "KeycloakSessionManage";
}


pub enum ValidationError {
    SessionError,
    AccessTokenExpireError,
}

