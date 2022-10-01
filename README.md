# Momo Store aka Пельменная №2

<img width="900" alt="image" src="https://user-images.githubusercontent.com/9394918/167876466-2c530828-d658-4efe-9064-825626cc6db5.png">

## Frontend

```bash
npm install
NODE_ENV=production VUE_APP_API_URL=http://localhost:8081 npm run serve
```

## Backend

```bash
go run ./cmd/api
go test -v ./... 
```

terraform init -backend-config=backend.tfvars
terraform apply -var-file="secret.tfvars"


helm upgrade --install --atomic momo-secrets ./infrastructure/momo-store-helm/charts/secrets/
helm upgrade --install --atomic -f ./infrastructure/momo-store-helm/values.yaml momo-backend ./infrastructure/momo-store-helm/charts/backend/
helm upgrade --install --atomic -f ./infrastructure/momo-store-helm/values.yaml momo-frontend ./infrastructure/momo-store-helm/charts/frontend/