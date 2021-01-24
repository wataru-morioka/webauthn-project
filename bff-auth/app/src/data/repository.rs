use std::collections::HashMap;
use rusoto_dynamodb::{DynamoDb, GetItemInput, PutItemInput, DeleteItemInput, AttributeValue};
use serde_dynamodb;
use serde::{Deserialize, Serialize};
use async_trait::async_trait;
use log::{info, error};
use crate::common::model::SessionId;
use super::connection::DYNAMODB_CLIENT;
use super::interface::repository::*;
use crate::data::entity::SessionInfo;
use super::connection;

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

        if let Err(error) = DYNAMODB_CLIENT.put_item(create_serials).await {
            println!("Error: {:?}", error);
            return Err(());
        };
        info!("セッション更新");
        Ok(())
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
    
        if let Err(error) = DYNAMODB_CLIENT.delete_item(delete_serials).await {
            error!("セッション削除エラー: {:?}", error);
            return Err(());
        };
        info!("セッション削除");
        Ok(())
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

        match DYNAMODB_CLIENT.get_item(get_serials).await {
            Ok(result) => {
                match result.item {
                    Some(item) => {
                        let session: SessionInfo = serde_dynamodb::from_hashmap(item).unwrap(); 
                        info!("get session!!: {}", session.session_id);
                        info!("get code_verifier!!: {}", session.code_verifier);
                        return Ok(session);
                    },
                    None => {
                        error!("no session");
                        return Err(());
                    }
                }
            },
            Err(error) => {
                error!("セッション取得エラー: {:?}", error);
                return Err(());
            },
        };
    }
}

#[async_trait(?Send)]
impl ApiInterface for ApiRepository {
    async fn token_request<T>(
        endpoint: String, 
        request_body: &Vec<(&str, &str)>,
        client_id: String,
        client_secret: String
    ) -> Result<T, ()>
    where T: for<'a> Deserialize<'a> {
        let result = connection::generate_api_client()
            .post(endpoint)
            .basic_auth(client_id, Some(&client_secret))
            .header("Content-Type", "application/x-www-form-urlencoded")
            .send_form(request_body)
            .await
            .expect("トークンリクエストエラー")
            .body()
            .limit(20_000_000) 
            .await;

        match result {
            Ok(bytes) => {
                let body: String = bytes.iter().map(|&s| s as char).collect::<String>();
                match serde_json::from_str(&body) {
                    Ok(token_res) => {
                        return Ok(token_res);
                    },
                    Err(err) => {
                        error!("トーンオブジェクト変換エラー {:?}", err);
                        return Err(());
                    }
                }
            },
            Err(_err) => {
                error!("トークン取得エラー {:?}", _err);
                return Err(());
            }
        };
    }
}
