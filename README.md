## Setup
1. Build image
    ```
    docker build -t danger-zone-ms .
    ```
2. Run image
   ```
   docker run -d --name danger-zone-ms \
     -e POSTGRES_HOST='localhost:5432' \
     -e POSTGRES_USER='postgres' \
     -e POSTGRES_PASSWORD='' \
     -e REDIS_HOST='redis:6379' \
     -e REDIS_PASSWORD='' \
     -e REDIS_CHANNEL_CREATE_ZONE='c-dangerzone:create' \
     -e REDIS_CHANNEL_DELETE_ZONE='c-dangerzone:delete' \
     --network hte \
     -p 8080:8080 \
     danger-zone-ms
   ```
---

## [Sequence diagram](https://sequencediagram.org/index.html#initialData=C4S2BsFMAIBEEMB2BzSAnaAtA9omBZEAYzWwGV0A3YyAKFoAd41QiQnFhoAifbAIxBRoAQQYNu0eAGcp4xs1bskXbghTpoAL1wFipCmmpFIkmdAAmSVGgD6AW2m0rwePxkxuAJUgWQ0s1k0X39neFd3aU8+FGxYACFA6HtcZGx6AF4MgGFg8JgrDQwdPCzaeHEAWgA+QpsHaQAuaAAFAHkyABVoAHo69BLIWQBvC0hjSABJWAAaaHBwufBcObR4PwBXaTngCABfMKKGmpTY5tzIfOgwSHtD+sca4L8m1oBVeIAZSbIACWgiJV+mhBk0SJdgDBRuMaNMlot5itoGtNttoLtwAcKgwADyVIHWdANRoAJgADABGaDQiZw+YI5aIVbrEBbHb7TIZADikC48HA4EshOKumkZWxNWBxOgXIAot0+sLQfciY9qqc0s0fOtrpC7ti8QSjo5SWSydAANo02GzenAJZIlGstEYvZzAB0noAupyeVxgdpdNB+ABPSwwkzQabiqq1YXSuUK4GggD8YwmthAFgy6ZtKrsao12C1lwsutu5XEhqlJvJ5utJjpC3tiKZyJZbPRHNoWVgkCgkKFRUDpQylYYkvjJrgss+8tlvWTorTEcgmezucbsHzx3VqWLcH7vJgNzuNekTxCTRaH2+fwBRpsoMaYwHBVX03H1anTXJABZ6CAA)
```
title Danger Zone MicroService

participant "Mobile App" as app
participant "Danger zone MicroService" as danger_ms
database "Redis" as redis
database "MongoDB" as mongo

==Create danger zone==
app->danger_ms: POST /dangerzones {deviceID, lat, lon, radius, ttl}
danger_ms->mongo: Create item
danger_ms->redis: PUBLISH c-dangerzones:create {deviceID, lat, lon, radius, ttl}
app<--danger_ms:201 {deviceID, lat, lon, radius, ttl}

==Get all danger zones==
app->danger_ms: GET /dangerzones
danger_ms->mongo: Read item
app<--danger_ms:200 [{deviceID, lat, lon, radius, ttl}, ...]

==Get danger zone by device ID==
app->danger_ms: GET /dangerzones?device_id=deviceID
danger_ms->mongo: Read item
app<--danger_ms:200 {deviceID, lat, lon, radius, ttl}

==Delete danger zone==
app->danger_ms: DELETE /dangerzones?device_id=deviceID
danger_ms->mongo: Delete item
danger_ms->redis:PUBLISH c-dangerzones:delete deviceID
app<--danger_ms:204
```