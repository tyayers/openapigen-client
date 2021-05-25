# openapigen-client
This is a simple web client to generate OpenAPI specs from deployed REST endpoints.  The background of this project is that I had a lot of API proxies delivering large, unstructured payloads, and needed a way to automatically generate (and update specs) based on the returned data structures.  Usually OpenAPI tooling generates the specs from source code objects, but since this data is coming from test backend systems (SAP, etc..), I needed an inspector service to generate the specs based on live data.

Currently the project is in a simple form (only handling flat resources), but could be easily expanded to more complex use-cases, if needed.

## Try it out
You can use a deployed version of the client deployed on Firebase here: [https://openapigen.web.app/](https://openapigen.web.app/).

![App preview](img/apppreview.gif)  

## Build

If you have Go installed, you can run the client locally with this command.

```
go run client.go
```

You can also easily build and deploy the containerized version to Google Cloud by running the build.sh command with your GCP project of choice as parameter (this will build, publish and deploy the container to Cloud Run, as long as you have the **gcloud** CLI installed and authenticated).

```
./build.sh $PROJECT
```

Where $PROJECT is an environment variable set to your preferred GCP project.

## Deploy
Deploy with one click to Google Cloud:

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

Or build the container as described above and deploy to your container orchestrator of choice.

## Credits
The library [openapigen](https://github.com/tyayers/openapigen) is used for the spec generation (library was created for this project).

The web client uses the micro [lit.css](https://ajusa.github.io/lit/docs/lit.html?) framework for styling and layout.