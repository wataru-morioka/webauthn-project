pub struct Validation {};

pub trait ValidationInterface{
    fn is_session_valid() -> Result<(), ()>;
    fn is_access_token_valid() -> Result<(), ()>;
}