#!/bin/bash

accountId=$(aws sts get-caller-identity --query "Account" --output text)
lambdaTopic="stockerize-handle-events"
lambdaName="stockerizeEvents"

echo "Desplegando EVENTS"

# Paso 1: Compilar el archivo principal de Go para la arquitectura especificada
echo "Compilando el código fuente..."
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap ./events/main.go

# Paso 2: Crear un archivo zip que incluya el ejecutable, el archivo de configuración y los activos
echo "Empaquetando los archivos necesarios en un archivo zip..."
zip -r $lambdaName.zip bootstrap .env

# Paso 3: Crear/Actualizar la función de Lambda en AWS usando el CLI de AWS
if aws lambda get-function --function-name $lambdaName 2>&1 | grep -q 'ResourceNotFoundException'
then
  echo "Creando función en AWS Lambda..."
  aws lambda create-function --function-name $lambdaName \
  --runtime provided.al2023 --handler bootstrap \
  --architectures arm64 \
  --memory-size 1024 \
  --timeout 60 \
  --role arn:aws:iam::$accountId:role/lambda-admin \
  --zip-file fileb://$lambdaName.zip
else 
  echo "Actualizando código de la función Lambda..."
  aws lambda update-function-code --function-name $lambdaName \
  --zip-file fileb://$lambdaName.zip
fi

# Paso 4: Crear topic si no existe
if echo "$existingTopics" | grep -q "$lambdaTopic"; then
  echo "El topic $lambdaTopic ya existe."
else
  aws sns create-topic --name $lambdaTopic
  aws lambda add-permission --function-name $lambdaName \
    --statement-id sns-permission \
    --action lambda:InvokeFunction \
    --principal sns.amazonaws.com \
    --source-arn arn:aws:sns:us-east-2:$accountId:$lambdaTopic
  aws sns subscribe \
    --topic-arn arn:aws:sns:us-east-2:$accountId:$lambdaTopic \
    --protocol lambda \
    --notification-endpoint arn:aws:lambda:us-east-2:$accountId:function:$lambdaName
  echo "El topic $lambdaTopic ha sido creado."
fi

# Paso 5: cleanup
rm $lambdaName.zip
rm bootstrap
echo "Despliegue completado."