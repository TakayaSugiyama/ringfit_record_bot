### Deploy Cloud Functions

```bash
gcloud functions deploy GetIfttt --runtime=go116 --trigger-http --allow-unauthenticated --build-env-vars-file .env.yaml
```
