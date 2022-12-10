# App Template
This [template](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template) can be used to create an app stub for an [Eliona](https://www.eliona.io/) enviroment.


## Configuration

The app needs environment variables and database tables for configuration. To edit the database tables the app provides an own API access.


### Registration in Eliona ###

To start and initialize an app in an Eliona environment, the app have to registered in Eliona. For this, an entry in the database table `public.eliona_app` is necessary.


### Environment variables

~~`APPNAME`: must be set to `template`. Some resources use this name to identify the app inside an Eliona environment.~~-
Any name is possible... feel free to name your application everything you want.

for example `eliona-app`


```
export APPNAME=eliona-app
```

- `CONNECTION_STRING`: configures the [Eliona database](https://github.com/eliona-smart-building-assistant/go-eliona/tree/main/db). Otherwise, the app can't be initialized and started. (e.g. `postgres://postgress:secret@localhost:5432/postgres`)

```
export CONNECTION_STRING=postgres://user:pass@localhost:5432/iot
```

- `API_ENDPOINT`:  configures the endpoint to access the [Eliona API v2](https://github.com/eliona-smart-building-assistant/eliona-api). Otherwise, the app can't be initialized and started. (e.g. `http://api-v2:3000/v2`)

```
export API_ENDPOINT=http://localhost:3000/v2
```

- `API_TOKEN`: defines the secret to authenticate the app and access the API.

```
export API_TOKEN=secret
```

- `API_SERVER_PORT`(optional): define the port the API server listens. The default value is Port `3000`.

- `DEBUG_LEVEL`(optional): defines the minimum level that should be [logged](https://github.com/eliona-smart-building-assistant/go-eliona/tree/main/log). Not defined the default level is `info`.


Whole env configuration
```
APPNAME='eliona-app'';CONNECTION_STRING=postgres://postgres:secret@localhost:5432/iot;API_SERVER_PORT=2999;API_ENDPOINT=http://127.0.0.1:3000/v2;API_TOKEN=secret;DEBUG_LEVEL=debug
```

### Database tables ###

All tables from eliona-mock repository plus one additional table "fiat_currency".

- `public.fiat_currency`: fiat currencies (EUR, USD) that are supported by Eliona system.

**Generation**: to generate access method to database see Generation section below.
'
## References

### App API ###

The app provides its own API to access configuration data and other functions. The full description of the API is defined in the `openapi.yaml` OpenAPI definition file.

- [API Reference](https://eliona-smart-building-assistant.github.io/open-api-docs/?https://raw.githubusercontent.com/eliona-smart-building-assistant/app-template/develop/openapi.yaml) shows Details of the API

**Generation**: to generate api server stub see Generation section below.


### Eliona ###

No action involved.

## Tools

### Generate API server stub ###

For the API server the [OpenAPI Generator](https://openapi-generator.tech/docs/generators/openapi-yaml) for go-server is used to generate a server stub. The easiest way to generate the server files is to use one of the predefined generation script which use the OpenAPI Generator Docker image.

#### Windows
```
.\generate-api-server.cmd
```

#### Linux

Permit access to the script file

```
chmod 777 generate-api-server.sh
```

Re-generate API

```
./generate-api-server.sh
```


### Generate Database access ###

For the database access [SQLBoiler](https://github.com/volatiletech/sqlboiler) is used. The easiest way to generate the database files is to use one of the predefined generation script which use the SQLBoiler implementation. Please note that the database connection in the `sqlboiler.toml` file have to be configured.


#### Windows
```
.\generate-db.cmd
```

#### Linux

Permit access to the script file

```
chmod 777 generate-db.sh
```

Re-generate API

```
./generate-db.sh
```

### Dockerization

```
docker-compose up --build
```

#### Issues

Docker cleanup
```
docker system prune -a
```

### Arbitrary commands

go mod tidy

### Endpoint description

#### BTC endpoint

<i>We only care about USD price, so in the Eliona system there is entity "fiat_currency" inside public schema, which holds information about allowed currencies.</i>


```
http://{domain}:{appport}/v1/btc
```

### Congecko API response

#### API Response

```
{"time":{"updated":"Dec 4, 2022 09:16:00 UTC","updatedISO":"2022-12-04T09:16:00+00:00","updateduk":"Dec 4, 2022 at 09:16 GMT"},"disclaimer":"This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org","chartName":"Bitcoin","bpi":{"USD":{"code":"USD","symbol":"&#36;","rate":"17,028.3491","description":"United States Dollar","rate_float":17028.3491},"GBP":{"code":"GBP","symbol":"&pound;","rate":"14,228.7522","description":"British Pound Sterling","rate_float":14228.7522},"EUR":{"code":"EUR","symbol":"&euro;","rate":"16,588.0981","description":"Euro","rate_float":16588.0981}}}
```

