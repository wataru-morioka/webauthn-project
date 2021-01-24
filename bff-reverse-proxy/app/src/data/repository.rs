use std::collections::HashMap;
use rusoto_dynamodb::{DynamoDb, GetItemInput, PutItemInput, DeleteItemInput, AttributeValue};
use serde_dynamodb;
use async_trait::async_trait;
use log::{info, error};
use reqwest::{Client, Response};
use crate::data::interface::model::SessionId;
use super::connection::DYNAMODB_CLIENT;
use super::interface::repository::*;
use crate::data::interface::model::SessionInfo;
use crate::config::errconfig::ValidationError;
use crate::config::envconfig::ENV;

#[async_trait]
impl DynamoDbInterface for DynamoDbRepository {
    async fn put_session(session: &SessionInfo, table_name: &str) -> Result<(), ()> {
        let put_item: HashMap<String, AttributeValue> = serde_dynamodb::to_hashmap(session)
        .expect("sessionオブジェクト変換エラー");

        let create_serials = PutItemInput {
            item: put_item,
            table_name: String::from(table_name),
            ..Default::default()
        };

        DYNAMODB_CLIENT.put_item(create_serials).await
            .map_err(|error| {
                error!("セッション更新エラー: {:?}", error);
                ()
            })
            .and_then(|_| {
                info!("セッション更新");
                Ok(())
            })   
    }

    async fn delete_session(session_id: String, table_name: &str)  -> Result<(), ()> {
        let key: HashMap<String, AttributeValue> = serde_dynamodb::to_hashmap(
            &SessionId {
                session_id,
            }
        ).unwrap();
    
        let delete_serials = DeleteItemInput {
            key: key,
            table_name: String::from(table_name),
            ..Default::default()
        };
    
        DYNAMODB_CLIENT.delete_item(delete_serials).await
            .map_err(|error| {
                error!("セッション削除エラー: {:?}", error);
                ()
            })
            .and_then(|_| {
                info!("セッション削除");
                Ok(())
            }) 
    }

    async fn get_session(session_id: String, table_name: &str) -> Result<SessionInfo, ()> {
        let key: HashMap<String, AttributeValue> = serde_dynamodb::to_hashmap(
            &SessionId {
                session_id,
            }
        ).unwrap();

        let get_serials = GetItemInput {
            key: key,
            table_name: String::from(table_name),
            ..Default::default()
        };

        DYNAMODB_CLIENT.get_item(get_serials).await
            .map_err(|error| {
                error!("セッション取得エラー: {:?}", error);
                ()
            })
            .and_then(|result| {
                result.item
                    .ok_or({
                        error!("no session");
                        ()
                    })
                    .and_then(|item| {
                        let session: SessionInfo = serde_dynamodb::from_hashmap(item).unwrap();
                        info!("get session!!: {}", session.session_id);
                        info!("get code_verifier!!: {}", session.code_verifier);
                        Ok(session)
                    })
            })
    }
}

#[async_trait]
impl ApiInterface for ApiRepository {
    async fn token_request(params: &Vec<(&str, &str)>) -> Result<Response, ValidationError> {
        Client::new()
            .post(&ENV.cognito_token_endpoint)
            .basic_auth(ENV.cognito_clientid.clone(), Some(ENV.cognito_clientsecret.clone()))
            .form(params)
            .send()
            .await
            .map_err(|_| ValidationError::AccessTokenExpireError)
    }
}
