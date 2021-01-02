<p align="center">
  <img alt="kontrolio logo" src="../assets/logo.png" height="300" />
  <h3 align="center">kontrolio-api</h3>
  <p align="center">Time clock, clock card machine, punch clock, or time recorder API.</p>
</p>

## 🖥️ Hosting

Kontrolio's API is ready to be served on AWS Lambda. To host it yourself, you just need to have a running Postgres database, clone this repo, and configure the following secrets ([as seen on the `env.example` file](../.env.example)):

| Environment variable  | Description                     |
| --------------------- | ------------------------------- |
| DB_USER               | Databse user                    |
| DB_PASSWORD           | Database user's password        |
| DB_NAME               | Database name                   |
| DB_HOST               | Database host                   |
| AWS_ACCESS_KEY_ID     | Your AWS role access key ID     |
| AWS_SECRET_ACCESS_KEY | Your AWS role secret access key |

Your AWS role must have the following permissions granted:

- IAMFullAccess
- AmazonS3FullAccess
- CloudWatchLogsFullAccess
- AmazonAPIGatewayAdministrator
- AWSCloudFormationFullAccess
- AWSLambda_FullAccess
