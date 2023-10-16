![UJBlockchain](https://blockchain.uj.ac.za/static/images/main-logo.png)


# SARCHI
South Africa-Switzerland Bilateral Research Chair in Blockchain Technology aims to explore blockchain integrations with real-world applications and development in Agric food.

## Database
Set environment variable for postgres database using Viper. Database connection is needed for Gorm. Replace "..." with your credentials.

```
DNS: "host=localhost user=postgres password=... dbname=... port=... sslmode=disable TimeZone=Africa/Johannesburg"
```

## Create Testnet account
Create Hedera testnet account from https://portal.hedera.com/register. Replace "..." in development.yaml with your own details. 

```
ACCOUNT_ID: ...
DER_ENCODED_PUBLIC_KEY: ...
DER_ENCODED_PRIVATE_KEY: ...
HEX_ENCODED_PRIVATE_KEY: ...
HEDERA_NETWORK: ...

```

## Deployment
Follow step by step gist for deployment using: [Deployment](https://gist.github.com/ujblockchain/9152d51a574791ed95b7e4a39ae83a18)