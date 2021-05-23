docker build -t local/openapigen-client .
docker tag local/openapigen-client eu.gcr.io/$1/openapigen-client
docker push eu.gcr.io/$1/openapigen-client

gcloud run deploy openapigen-client --image eu.gcr.io/$1/openapigen-client --platform managed --project $1 --region europe-west1 --allow-unauthenticated
