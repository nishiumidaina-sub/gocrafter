# IdleBonus
放置ボーナス関連。  
[protobuf](https://github.com/game-core/gocrafter/tree/main/docs/proto/api/game/idleBonus)  

- [GetUser](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/idleBonus.md#GetUser)
- [GetMaster](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/idleBonus.md#GetMaster)
- [Receive](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/idleBonus.md#Receive)

## GetUser
放置ボーナスのユーザーステータスを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| req | *IdleBonusGetUserRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *IdleBonusGetUserResponse | レスポンス |
| err | error | エラー |

## GetMaster
放置ボーナスのマスターデータを取得する。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| req | *IdleBonusGetMasterRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *IdleBonusGetMasterResponse | レスポンス |
| err | error | エラー |

## Receive
放置ボーナスを受けとる。
- request

| Name | Type | Description |
| :--- | :--- | :--- |
| ctx | context.Context | コンテキスト |
| tx | *gorm.DB | トランザクション |
| now | time.Time | 現在時刻 |
| req | *IdleBonusReceiveRequest | リクエスト |

- response

| Name | Type | Description |
| :--- | :--- | :--- |
| res | *IdleBonusReceiveResponse | レスポンス |
| err | error | エラー |
