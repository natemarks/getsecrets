# getsecrets

Get AWS secrets from AWS Secrets Manager

## Usage
The environment variables below would permit getsecrets to download the secrets content for ONESIGNRO to onesignro.json and BIOMETRICRO to biometricro.json. It would also provide some additional information to the application if needed. In this example, the application might use the secret to  connect to the instance, then connect to the DATABASE.
```bash

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
{"level":"info","version":"a99ca69db5bd9ffbf424438d7c7c565b8bafe48e","time":"2023-10-18T06:24:12-04:00","message":"starting"}
{"level":"debug","version":"a99ca69db5bd9ffbf424438d7c7c565b8bafe48e","time":"2023-10-18T06:24:12-04:00","message":"will lookup DNASecretsSandboxBiometricb-Ez2NzslYaGgI for BIOMETRICRO and write it to biometricro.json"}
{"level":"debug","version":"a99ca69db5bd9ffbf424438d7c7c565b8bafe48e","time":"2023-10-18T06:24:12-04:00","message":"will lookup DNASecretsSandboxApplexones-zZcO0lSih1TE for ONESIGNRO and write it to onesignro.json"}
{"level":"info","version":"a99ca69db5bd9ffbf424438d7c7c565b8bafe48e","time":"2023-10-18T06:24:12-04:00","message":"wrote BIOMETRICRO - DNASecretsSandboxBiometricb-Ez2NzslYaGgI to biometricro.json"}
{"level":"info","version":"a99ca69db5bd9ffbf424438d7c7c565b8bafe48e","time":"2023-10-18T06:24:13-04:00","message":"wrote ONESIGNRO - DNASecretsSandboxApplexones-zZcO0lSih1TE to onesignro.json"}
```