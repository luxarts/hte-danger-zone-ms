# hte-danger-zone-ms

### Setup
1. Build image
    ```
    docker build -t danger-zone-ms .
    ```
2. Run image
   ```
   docker run -d --name danger-zone-ms \
     -e MONGO_HOST='mongo:27017' \
     -e MONGO_DB='core' \
     -e MONGO_DANGER_ZONE_COLLECTION='dangerzones' \
     -e REDIS_HOST='redis:6379' \
     -e REDIS_PASSWORD='' \
     -e REDIS_CHANNEL_CREATE_ZONE='c-dangerzone:create' \
     --network hte \
     -p 8080:8080 \
     danger-zone-ms
   ```
   
### Create danger zone
Method: `POST`  
URL: `/:companyID/dangerzone`  
Body:  
```json
{
    "device_id": "someDevice",
    "latitude": 123.4567,
    "longitude": -123.4567,
    "radius": 2.12,
    "ttl": 3600
}
```