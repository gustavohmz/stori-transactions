
# Stori Transactions API

## Project Description

The Stori Transactions API is a solution designed to process CSV files containing account transactions, calculate a financial summary, and send the summary via email. This project utilizes AWS S3 for file upload and retrieval, DynamoDB for storing transactions, and Go for the API development.

## Key Features
- Process a CSV file uploaded to an S3 bucket.
- Store transactions in DynamoDB.
- Calculate:
    - Total balance.
    - Number of transactions grouped by month.
    - Average debit and credit amounts per month.
- Send a summary via an HTML-styled email that includes a brand logo.

## Requirements
- Docker installed on your machine to run the API in a containerized environment.
- AWS CLI installed and configured on your machine.
    - Install AWS CLI.
    - Configure the credentials for AWS CLI using the following command:
    ```bash
    aws configure
    ```
    Enter the values for the following prompts:

    - AWS Access Key ID: `<your-aws-access-key-id>` (from the `.env` file)
    - AWS Secret Access Key: `<your-aws-secret-access-key>` (from the `.env` file)
    - Default region: `<your-aws-region>` (from the `.env` file, e.g., `us-east-1`)
    - Default output format: `json`
- Note: AWS CLI is required to upload the CSV file to the S3 bucket as part of the API workflow.
- Access to:
    - An S3 bucket for storing the CSV file.
    - A DynamoDB table named Transactions.
    - AWS SES for sending emails (see note below).

## Project Configuration
To properly run the project, create a `.env` file in the root directory with the following variables:

```bash
# AWS Configuration
AWS_ACCESS_KEY_ID=<your-aws-access-key-id>
AWS_SECRET_ACCESS_KEY=<your-aws-secret-access-key>
AWS_REGION=<your-aws-region>
S3_BUCKET=<your-s3-bucket-name>
S3_INPUT_FILE=<path-to-your-csv-file-in-s3>

# DynamoDB Table Name
DYNAMODB_TABLE=Transactions

# Email Service Configuration
SMTP_FROM=<your-email-address>
SMTP_TO=<recipient-email-address>
SMTP_USERNAME=<your-smtp-username>
SMTP_PASSWORD=<your-smtp-password>
SMTP_HOST=<your-smtp-host>
SMTP_PORT=587

# Email Assets
EMAIL_LOGO_URL=<url-of-your-logo>
```
**Note: For security reasons, the `.env` file is not included in the repository and will be provided separately. Make sure to fill in the correct values before running the API.**

## How to Use the Project
### Steps to Run the API
1. Clone the repository to your local machine:
```bash     
git clone https://github.com/gustavohmz/stori-transactions.git
cd stori-transactions

```
2. Place the provided `.env` file in the root directory.

3. Build and run the Docker image:
```bash     
docker build -t stori-transactions-api .
docker run -p 8080:8080 --env-file .env stori-transactions-api
```
4. Upload the CSV file to the S3 bucket:
```bash     
aws s3 cp ./data/sample-transactions.csv s3://stori-transactions-file/data/sample-transactions.csv 
```

5. The API will automatically process the file, calculate the summaries, and send an email with the summary.

## Considerations
+ **CSV File Format:** The CSV file must follow the structure below:
```bash     
ID,Amount,Type,Date,AccountID
1,100.50,credit,2024-08-09,account-123
2,-50.25,debit,2024-11-09,account-123
3,200.00,credit,2024-09-09,account-456
4,-75.00,debit,2024-11-09,account-456
```

+ **AWS SES:** Ensure that the recipient email address is verified in AWS SES to avoid errors when sending emails.
## Technical Notes
+ **Architecture:** This project follows a hexagonal architecture for better maintainability and scalability.
+ **Technologies Used:**
    + **AWS S3:** Used to store the CSV file and the logo included in the email.
    + **DynamoDB:** A NoSQL database used to store the processed transactions.
    + **Go:** Programming language used to develop  the API.
    + **Docker:** Used for containerization and deployment.
