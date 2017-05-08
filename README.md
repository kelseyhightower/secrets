# Secrets 

An example application that prints secrets then exits.

## Usage

```
kubectl create -f secrets.yaml
``` 

## Build Container

```
gcloud container builds submit --config cloudbuild.yaml .
```
