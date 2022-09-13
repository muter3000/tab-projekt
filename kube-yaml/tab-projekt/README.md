# TAB-Projekt Helm Chart

## Wymagania

Użytkownik posiada na urządzeniu z którego będzie przebiegała instalacja **kubectl** oraz **helm**

## Architektura projektu

Projekt zakłada że użytkownik posiada klaster Kubernetes na którym może zainstalować projekt. Do celów testowych używać
można **k3d** lub **minikube**.

## Deployment backendu

Aby zainstalować projekt na klastrze należy

1. Zainstalować deployment postgreSQL przy pomocy **helm**
   ```shell
   helm repo add bitnami https://charts.bitnami.com/bitnami
   helm repo update
   helm install postgresql bitnami/postgresql
   ```
2. Zainstalować nginx controller w celu dostępu z zewnątrz
   ```shell
   helm install ngnix nginx-stable/nginx-ingress
   ```
3. Zainstalować helm deployment projektu
   ```shell
   helm repo add nginx-stable https://helm.nginx.com/stable
   helm repo update
   helm install tab-projekt .
   ```
