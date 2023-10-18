# getsecrets

Access secrets a JSON string from AWS secretsmanager and store them in a file.

## Usage
The environment variables below would permit getsecrets to download the secrets content for MYDB1 to mydb1.json and MYDB2 to mydb2.json. It would also provide some additional information to the application if needed

example execution with debug logging
```bash
GETSECRETS_DEBUG="true" \
GETSECRETS_ONESIGNRO_SECRETID="DNASecretsSandboxApplexones-zZcO0lSih1TE" \
GETSECRETS_ONESIGNRO_FILENAME="onesignro.json" \
GETSECRETS_ONESIGNRO_DATABASE="onesign" \
GETSECRETS_BIOMETRICRO_SECRETID="DNASecretsSandboxBiometricb-Ez2NzslYaGgI" \
GETSECRETS_BIOMETRICRO_FILENAME="biometricro.json" \
GETSECRETS_BIOMETRICRO_DATABASE="biometric" \
./build/current/linux/amd64/getsecrets
{"level":"info","version":"","time":"2023-10-18T06:14:50-04:00","message":"starting"}
{"level":"debug","version":"","time":"2023-10-18T06:14:50-04:00","message":"will lookup DNASecretsSandboxBiometricb-Ez2NzslYaGgI for BIOMETRICRO and write it to biometricro.json"}
{"level":"debug","version":"","time":"2023-10-18T06:14:50-04:00","message":"will lookup DNASecretsSandboxApplexones-zZcO0lSih1TE for ONESIGNRO and write it to onesignro.json"}
{"level":"info","version":"","time":"2023-10-18T06:14:50-04:00","message":"wrote ONESIGNRO - DNASecretsSandboxApplexones-zZcO0lSih1TE to onesignro.json"}
{"level":"info","version":"","time":"2023-10-18T06:14:50-04:00","message":"wrote BIOMETRICRO - DNASecretsSandboxBiometricb-Ez2NzslYaGgI to biometricro.json"}
```