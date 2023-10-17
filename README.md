# getsecrets

Access secrets a JSON string from AWS secretsmanager and store them in a file.

## Usage
The environment variables below would permit getsecrets to download the secrets content for MYDB1 to mydb1.json and MYDB2 to mydb2.json. It would also provide some additional information to the application if needed

```bash
```bash
export GETSECRETS_DEBUG="true"
export GETSECRETS_MYDB1_SECRETID="some-secretsmanager-id1"
export GETSECRETS_MYDB1_FILENAME="mydb1.json"
export GETSECRETS_MYDB2_SECRETID="some-secretsmanager-id2"
export GETSECRETS_MYDB2_FILENAME="mydb2.json"
export GETSECRETS_MYDB2_DATABASE="some_db_name"
```