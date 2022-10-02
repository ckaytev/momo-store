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
helm upgrade --install --atomic momo-secrets ./infrastructure/momo-store-helm/charts/secrets/
helm upgrade --install --atomic -f ./infrastructure/momo-store-helm/values.yaml momo-backend ./infrastructure/momo-store-helm/charts/backend/
helm upgrade --install --atomic -f ./infrastructure/momo-store-helm/values.yaml momo-frontend ./infrastructure/momo-store-helm/charts/frontend/

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm upgrade --install --atomic ingress-nginx ingress-nginx/ingress-nginx

helm repo add jetstack https://charts.jetstack.io
helm repo update
helm upgrade --install --atomic cert-manager jetstack/cert-manager --namespace momo-store --version v1.9.1 --set installCRDs=true

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade --install --atomic prometheus --set serviceMonitorsSelector.app=prometheus --set ruleSelector.app=prometheus prometheus-community/kube-prometheus-stack

helm repo add loki https://grafana.github.io/loki/charts
helm repo update
helm upgrade --install --atomic loki loki/loki-stack


helm upgrade --install --atomic ingress ./infrastructure/momo-store-helm/charts/ingress/
