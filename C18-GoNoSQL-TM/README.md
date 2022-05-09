# DynamoDB

## Requisitos

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Hombrew/install/HEAD/install.sh)"
brew -v
brew cask install docker
docker -v
```

### Docker

Plataforma de software que permite crear, probar e implementar aplicaciones rápidamente. Docker empaqueta software en unidades estandarizadas llamadas contenedores que incluyen todo lo snecesario para que el software se ejecute, incluidas bibliotecas, herramientas del sistema, código y tiempo de ejecución.

### VM vs Contenedor

La VM virutaliza el Sistema Operativo, esto implica que cada VM necesita instalar su propio OS que demanda gran cantidad de recursos lo que afecta el performance del servidor. Mientras que en Docker se utiliza el OS del servidor, haciendo uso del nucleo del OS propio del servidor y solamente virtualizando los programas utilizados en dicho contenedor, lo que reduce el consumo de recursos, ademas de facilitar la gestión del network dentro del host. Tambien nos permite aislar procesos con el fin de garantizar seguridad.

## Qué es DynamoDB

Base de datos clave-valor y documentos que ofrece rendimiento en milisegundos de un solo dígito a cualquier escala, puede gestionar más de 10 billones de solicitudes por día y puede admitir picos de más de 20 millones de solicitudes por segundo.

### Docker Compose para DynamoDB

Herraminta desarrollada para yudar a definir y compartir aplicaciones de arios contenedores. Con Compose, puede crear un archivo YAML para definir los servicios y, con un solo comando, ponerlo todo en marcha o eliminarlo.

```yml
services:
    dynamodb-local:
        image: amazon/dynamodb-local:latest
        container_name: dynamodb-local
        ports:
            - 8000:8000
        command: "-jar DynamoDBLocal.jar -sharedDb -dbPath ."
    dynamodb-admin:
        image: aaronshaf/dynamodb-admin
        ports:
            - 8001:8001
        environment:
            DYNAMO_ENDPOINT: "http://dynamodb-local:8000"
        depends_on:
            - dynamodb-local
```

#### Run and Down

```bash
docker-compose up
docker-compose down
```

### Tables

Forma de guardar documentos en DynamoDb, podemos indicar el nombre de la tabla, el Hash Attribute Name, la primary key y el tipo de dato de la primary key por la que vamos a identificar cada documento.

## DynamoDb Repository

### Package was/was-sdk-go

Nos permite interactuar con los servicios de AWS.

```bash
go get github.com/aws/aws-sdk-go/aws
```

```go
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/dynamodb"
)

func initDynamo() (*dynamodb.DynamoDB, error) {
    region := "us-west-2"
    endpoint := "http://localhost:8000"
    cred := credential.NewStaticCredentials("local", "local", "")
    sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))
    if err != nil {
        retutn nil, err
    }
    dynamo := dynamo.New(sess)
    return dynamo, nil
}
```

### Implementation

```go
type User struct {
    Id int
    Name string
    Age int
}

func itemToUser(input map[string]*dynamodb.AttributeValue) (User, error) {
    var item User
    err := dynamodbattribute.UnmarshalMap(input, &item)
    if err!= nil {
        return nil, err
    }
    return &item, nil
}
```

```go
type repository {
    dynamo *dynamodb.DynamoDB
    table string
}

func (receiver *repository) GetOne(ctx context.Context, id string) (User, error) {
    result, err := receiver.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
        TableName: aws,String(receiver.table),
        Key: map[string]*dynamodb.AttributeValue{
            "id": {
                S: aws.String(id)
            }
        }
    })

    if err != nil {
        return User{}, err
    }
    if result.Item == nil {
        return User{}, nil
    }
    return itemToUser(result.Item)
}

func (receiver *repository) Store(ctx context.Context, model *User) error {
    av, err := dynamodbattribute.MarshalMap(model)
    if err != nil {
        return err
    }
    input := &dynamodb.PutItemInput({
        Item: av,
        TableName: aws.String(receiver.table)
    })

    _, err = receiver.dynamo.PutItemWithContext(ctx, input)

    if err != nil {
        return err
    }

    return nil
}
```
