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

kubectl create namespace momo-store

cd infrastructure/momo-store-helm/
helm dependency build
helm upgrade --install --atomic momo-secrets charts/secrets/
helm upgrade --install --atomic -f values.yaml momo-backend charts/backend/
helm upgrade --install --atomic -f values.yaml momo-frontend charts/frontend/
helm upgrade --install --atomic ingress-nginx ingress-nginx/ingress-nginx
helm upgrade --install --atomic cert-manager jetstack/cert-manager --namespace momo-store --set installCRDs=true
helm upgrade --install --atomic prometheus -f  values.yaml prometheus-community/kube-prometheus-stack
helm upgrade --install --atomic loki loki/loki-stack
helm upgrade --install --atomic ingress charts/ingress/