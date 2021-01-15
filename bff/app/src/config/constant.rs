pub struct Const;

impl Const {
    pub const AUTH_COOKIE: &'static str = "auth_session";
    pub const COGNITO_SESSION_MANAGE: &'static str = "CognitoSessionManage";
    pub const KEYCLOAK_SESSION_MANAGE: &'static str = "KeycloakSessionManage";
    pub const COOKIE_EXPIRES_IN: i64 = 100;
}