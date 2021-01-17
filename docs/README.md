![Go](https://github.com/marcelovicentegc/kontrolio-api/workflows/Go/badge.svg) [![serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com)

<p align="center">
  <img alt="kontrolio logo" src="../assets/logo.png" height="300" />
  <h3 align="center">kontrolio-api</h3>
  <p align="center">Time clock, clock card machine, punch clock, or time recorder API.</p>
</p>

### Development

- Make sure you have Docker installed and running on your machine, as the `serverless-offline` plugin needs it to spin its container.
- Set the following secrets ([as seen on the `env.example` file](../.env.example)):

| Environment variable  | Type            | Description                                                           |
| --------------------- | --------------- | --------------------------------------------------------------------- |
| DB_USER               | `string`        | Databse user                                                          |
| DB_PASSWORD           | `string`        | Database user's password                                              |
| DB_NAME               | `string`        | Database name                                                         |
| DB_HOST               | `string`        | Database host                                                         |
| AWS_ACCESS_KEY_ID     | `string`        | Your AWS role access key ID                                           |
| AWS_SECRET_ACCESS_KEY | `string`        | Your AWS role secret access key                                       |
| CLIENT_URL            | `string`        | The client URL, needed because of CORS policy                         |
| SENDER_EMAIL          | `string`        | Sender email (it's only used when ENABLE_EMAIL_AUTH is set to `true`) |
| ENABLE_EMAIL_AUTH     | `bool` (0 or 1) | Email authentication feature flag                                     |
| JWT_SECRET            | `string`        | JWT secret key                                                        |

- Finally, start the application: `yarn start`

## 🖥️ Hosting

Kontrolio's API is ready to be served on AWS Lambda. To host it yourself, you need to have a running Postgres database and access to the AWS console to configure the SSM and SES services.

Your AWS role must have the following permissions granted:

- IAMFullAccess
- AmazonS3FullAccess
- CloudWatchLogsFullAccess
- AmazonAPIGatewayAdministrator
- AWSCloudFormationFullAccess
- AWSLambda_FullAccess
- AmazonSSMReadOnlyAccess

### Configuration

On the AWS console, navigate to `Systems Manager > Application Management > Parameter Store` and set the following secrets with your production credentials, just as you did while developing it:

| Variable name          | Type           | Description                                                           |
| ---------------------- | -------------- | --------------------------------------------------------------------- |
| kontrolio-db-user      | `SecureString` | Databse user                                                          |
| kontrolio-db-password  | `SecureString` | Database user's password                                              |
| kontrolio-db-name      | `SecureString` | Database name                                                         |
| kontrolio-db-host      | `SecureString` | Database host                                                         |
| kontrolio-sender-email | `SecureString` | Sender email (it's only used when ENABLE_EMAIL_AUTH is set to `true`) |
| kontrolio-jwt-secret   | `SecureString` | JWT secret key                                                        |

On your pipeline, you'll need the AWS keys to deploy, thus, set:

| Environment variable  | Type     | Description                     |
| --------------------- | -------- | ------------------------------- |
| AWS_ACCESS_KEY_ID     | `string` | Your AWS role access key ID     |
| AWS_SECRET_ACCESS_KEY | `string` | Your AWS role secret access key |
