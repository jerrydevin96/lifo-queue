# LIFO Queue

#### *LIFO Queue implemented in Golang using postgres/maria data backend*

This project implements a simple LIFO queue in *Golang* with *postgresql/mariadb* acting as the databackend. The project can be complied and packaged into a *Docker* image and deployed in *Kubernetes*. Two rest API end points are exposed for pushing elements into the queue or popping elements from the queue. The element/value pushed into the queue is stored in a table either in a Postgresql DB or a MariaDB (DB provider can be selected during deployment). The final application endpoints are exposed via nginx ingress controller. Tests were performed on a single node *Rancher* cluster.

------



##### Application Specs:

- The Application is deployed in Kubernetes (Rancher used for testing)
- Databases are setup on Docker
- Application is exposed via Nginx ingress controller

###### API Endpoints:

| **API method** | **API Endpoint** |            **Function**             |
| :------------: | :--------------: | :---------------------------------: |
|      GET       |     /v1/pop      | Pops the last element in the queue  |
|      POST      |     /v1/push     | Pushes a new element into the queue |

###### DB Table Design:

| **Column Name** |     **Data Type**      | **Constraints** |
| :-------------: | :--------------------: | :-------------: |
|      entry      |        integer         |     unique      |
|      value      | character varying(255) |                 |

##### Deployment Instructions:

###### Setup Database:

The Application allows connections to Postgresql DB or a Maria DB. The databases can be setup on docker using the below commands.

***Postgres Setup:***

```shell
docker run -d -p 5432:5432 -e POSTGRES_USER=<postgres root user> \
	-e POSTGRES_PASSWORD=<postgres root password> \
	-e POSTGRES_DB=<database name> \
	--name postgres-poc postgres
```

***Maria Setup:***

```shell
docker run -d -p 3306:3306 --name mariadb-poc \
	-e MYSQL_ROOT_PASSWORD=root \
	-e MYSQL_DATABASE=<database name>
	-e MYSQL_USER=<DB user> \
	-e MYSQL_PASSWORD=<DB password> \
	mariadb
```

###### Application Deployment:

The application can be deployed into any kubernetes cluster(Rancher used for testing). Before deployment, the cluster should be ready and kubeconfig should be pointing to the cluster. The following are the deployment steps:

- Clone the repository (https://github.com/jerrydevin96/lifo-queue.git) on the machine, where kubectl is installed and configured.

- Navigate to deployment folder

- Edit *config.yml* and provide appropriate values

  ```yaml
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: appconfig
  data:
    DBHOST: "<db host>"
    DBPORT: "<db port>"
    DBUSER: "<db user>"
    DBNAME: "<db name>"
    TABLENAME: "<db table name>"
    DBPROVIDER: "<db provider (postgres/maria)>"
    REINITIALIZE_TABLE: "<Y/N>" #will drop existing table and create a new one if 'Y'
  ```

- Edit *secrets.yml* and provide appropriate values

  ```yaml
  apiVersion: v1
  kind: Secret
  metadata:
    name: appsecrets
  type: Opaque
  stringData:
    DBPASSWORD: "<db password>"
  ```

- Apply all resources inside deployments folder into the cluster

  ```shell
  kubectl apply -f deployment/ -n <namespace>
  ```

- The application will be deployed and the API endpoints can be accessed via the ingress endpoint.