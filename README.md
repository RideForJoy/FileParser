

## Deploy from local machine:
```bash
gcloud functions deploy pc-file-reader --gen2 --runtime=go119 --memory=16Gi --region=us-central1 --source=. --entry-point=ProcessFile --trigger-event-filters="type=google.cloud.storage.object.v1.finalized" --trigger-event-filters="bucket=pc-file"
```
