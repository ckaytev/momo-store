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

## [Momo-store](https://momo-store.mooo.com/)

## CI/CD

- используется единый [репозиторий](https://gitlab.praktikum-services.ru/antona-zateyev/momo-store)
- развертывание приложение осуществляется с использованием [Downstream pipeline](https://docs.gitlab.com/ee/ci/pipelines/downstream_pipelines.html#parent-child-pipelines) 
- при изменениях в соответствующих директориях триггерятся pipeline для backend, frontend и infrastructure (momo-store-helm)
- backend и frontend проходят этапы сборки, тестирования, релиза, деплоя в dev-окружение (docker-compose) и prod-окружение (k8s)
- momo-store-helm проходит этапы релиза и деплоя в prod-окружение (k8s)
- trunk-based development

## Versioning

- [SemVer 2.0.0](https://semver.org/lang/ru/)
- мажорные и минорные версии приложения изменяются вручную в файлах `backend/.gitlab-ci.yaml` и `frontend/.gitlab-ci.yaml` в переменной `VERSION`
- патч-версии изменяются автоматически на основе переменной `CI_PIPELINE_ID`
- для инфраструктуры версия приложения изменяется вручную в чарте `infrastructure/momo-store-helm/Chart.yaml`

## Infrastructure

- код ---> [Gitlab](https://gitlab.praktikum-services.ru/)
- helm-charts ---> [Nexus](https://nexus.praktikum-services.ru/)
- анализ кода ---> [SonarQube](https://sonarqube.praktikum-services.ru/)
- docker-images ---> [Gitlab Container Registry](https://gitlab.praktikum-services.ru/antona-zateyev/momo-store/container_registry)
- терраформ бэкэнд и статика ---> [Yandex Object Storage](https://cloud.yandex.ru/services/storage)
- продакшн ---> [Yandex Managed Service for Kubernetes](https://cloud.yandex.ru/services/managed-kubernetes)

## Init kubernetes

- клонировать репозиторий на машину с установленным terraform
- через консоль Yandex Cloud создать сервисный аккаунт с ролью `editor`, получить статический ключ доступа, сохранить секретный ключ в файле `infrastructure/terraform/backend.tfvars`
- получить [iam-token](https://cloud.yandex.ru/docs/iam/operations/iam-token/create), сохранить в файле `infrastructure/terraform/secret.tfvars`
- через консоль Yandex Cloud создать Object Storage, внести параметры подключения в файл `infrastructure/terraform/provider.tf`
- выполнить следующие комманды:

```
cd infrastructure/terraform
terraform init -backend-config=backend.tfvars
terraform apply -var-file="secret.tfvars"
```

## Init production

```
# создаем базовый namespace
kubectl create namespace momo-store

# устанавливаем cert-manager
cd infrastructure/momo-store-helm/
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm upgrade --install --atomic -n momo-store cert-manager jetstack/cert-manager --version v1.9.1 --set installCRDs=true

# сохраняем креды для docker-registry
kubectl create secret generic -n momo-store docker-config-secret --from-file=.dockerconfigjson="/home/user/.docker/config.json" --type=kubernetes.io/dockerconfigjson 
# устанавливаем приложение, указав версии backend и frontend
helm dependency build
helm upgrade --install --atomic -n momo-store momo-store . --set backend.image.tag=latest --set frontend.image.tag=latest

# смотрим IP load balancer, прописываем А-записи для приложения и мониторинга
kubectl get svc
```

## [Monitoring](https://grafana.momo-store.mooo.com/)

- admin / uhfafyfltdjgc
- включен в состав helm-chart приложения, зависимости прописаны в `infrastructure/momo-store-helm/Chart.yaml`
- [infrastructure](https://grafana.momo-store.mooo.com/d/CgCw8jKZz3/golang-metrics-for-prometheus-operator?orgId=1&refresh=5s)

<img width="500" alt="image" src="https://storage.yandexcloud.net/momo-store-zateevant/monitoring/infrastructure.JPG">

- [frontend](https://grafana.momo-store.mooo.com/d/MsjffzSZz/nginx-exporter?orgId=1&refresh=5s)

<img width="500" alt="image" src="https://storage.yandexcloud.net/momo-store-zateevant/monitoring/frontend.JPG">

- [backend](https://grafana.momo-store.mooo.com/d/5i5VYXVVk/momo-store?orgId=1&from=now-1h&to=now)

<img width="500" alt="image" src="https://storage.yandexcloud.net/momo-store-zateevant/monitoring/backend.JPG">

- [logs](https://grafana.momo-store.mooo.com/d/liz0yRCZz/loki-logs-dashboard?var-namespace=momo-store&var-pod=All&var-search=)

<img width="500" alt="image" src="https://storage.yandexcloud.net/momo-store-zateevant/monitoring/logs.JPG">


## Backlog

- добавить тестовое окружение (отдельный кластер или отдельный namespace)
- вывести мониторинг из чарта самого приложения (ускорить деплой)
- поднять Vault для хранения секретов