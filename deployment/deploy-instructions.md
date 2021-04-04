### Deployment Instructions (minikube)

- Create two directories to store DB data

  ```shell
  mkdir -p /data/postgres
  mkdir -p /data/maria
  ```

   

- Deploy postgres and Maria using below command:

  ```shell
  kubectl apply -f database/
  ```

- Update config.yml, secrets.yml with appropriate values. 
  example config.yml

  ```yaml
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: appconfig
  data:
    DBHOST: "postgres-0.postgres"
    DBPORT: "5432"
    DBUSER: "user"
    DBNAME: "lifo"
    TABLENAME: "lifo"
    DBPROVIDER: "postgres"
    REINITIALIZE_TABLE: "N"
    SWAGGER_HOST: "10.15.20.158"
  ```

  example secrets.yml

  ```yaml
  apiVersion: v1
  kind: Secret
  metadata:
    name: appsecrets
  type: Opaque
  stringData:
    DBPASSWORD: "password"
  ```

- Modify ingress.yml and provide hostname for ingress

  ```yaml
  host: ec2-xx-xxx-xxx-xx.compute-1.amazonaws.com/
  ```

- Apply all remaining ymls

  ```shell
  kubectl apply -f config.yml
  kubectl apply -f secret.yml
  kubectl apply -f deployment.yml
  kubectl apply -f service.yml
  kubectl apply -f ingress.yml
  ```

  