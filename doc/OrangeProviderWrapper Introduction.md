# Orange Provider Wrapper Introduction

## 1. Background

Orange V1 对外部的provider的数据源或算法的支持比较复杂，而且只支持Ontology NeoVM的DID，需要定制开发的部分比较多。

## 2. Wrapper的介绍

Orange provider wrapper 提供下列功能：
1. 生成新的EVM钱包文件
2. 注册DID（钱包对应地址需要支付gas）
3. 通过自动对request内容验证签名
4. 自动对响应response进行加密（仅Data provider）
5. 通过配置文件简单的配置API，无需二次开发

## 3. config.json
wrapper 的配置文件，一般在项目的根目录下，可以在启动参数中指定其他的路径

```
{
    "orange_did":"did:etho:1ae43df6f4c5621e2b156162e958c80a67ee4f5f",
    "keystore":"./keystore",
    "wallet_pwd":"123456",
    "chain_id":5851,
    "chain_rpc":"http://polaris1.ont.io:20339",
    "contract_address":"0x18d3dB10B18369691c86e7EF99cBd9B290BaD87A",
    "api_configs":[
        {
            "provider_type":"dp",
            "verify_request":true,
            "server_path":"/balance1",
            "has_api_key":true,
            "api_key_location":"header",
            "api_key_name":"x-api-key",
            "api_key":"test",
            "api_url":"http://localhost:8088/sampleGetUrlDP",
            "api_method":"GET",
            "param_type":"url",
            "failed_keywords":[]
        },
        {
            "provider_type":"dp",
            "verify_request":false,
            "server_path":"/balance2",
            "has_api_key":true,
            "api_key_location":"header",
            "api_key_name":"x-api-key",
            "api_key":"test",
            "api_url":"http://localhost:8088/sampleGetBodyDP",
            "api_method":"POST",
            "param_type":"body",
            "failed_keywords":[]
        },{
            "provider_type":"ap",
            "verify_request":true,
            "server_path":"/score",
            "has_api_key":true,
            "api_key_location":"header",
            "api_key_name":"x-api-key",
            "api_key":"test",
            "api_url":"http://localhost:8088/sampleAP",
            "api_method":"POST",
            "param_type":"body",
            "failed_keywords":[]
        }
    ]
}
```

1. orange_did: Orange 系统的EVM did,用于对系统请求签名的验证，可以通过官网获得。
2. keystore: 本地的evm钱包文件的路径
3. wallet_pwd: 钱包文件的密码
4. chain_id: evm did合约所在链的chain id (https://chainlist.org/)
5. chain_rpc: evm did合约所在链的rpc (https://chainlist.org/)
6. contract_address: evm did合约地址
7. api_configs: api的配置信息
   1. provider_type: dp(数据提供方) | ap（算法提供方）
   2. verify_request: true | false  是否验证请求方签名
   3. server_path: 对外提供的wrapper服务的路径 http(s)://ip:port/path
   4. has_api_key: true | false 源api是否有api key
   5. api_key_location: header 源apikey所在位置
   6. api_key_name: 源apikey的名称
   7. api_key:源apikey的内容（req.Header.Set(ApiKeyName, ApiKey)）
   8. api_method:GET | POST 源api的method
   9. param_type:BODY | URL 源api的参数形式
   10. failed_keywords: 当源api返回包含这些关键字时，可以认为是失败

## 4. Tools
1. 生成新的钱包
   ```
   ./orange-provider-wrapper --operation new-wallet
   ```
   在config.json中指定的目录下以指定的password生成新的evm钱包

2.  注册evm did
    ```   
    ./orange-provider-wrapper --operation register-did
    ```