# psqltimestampz

PostgreSQL timestampz 調査

## snipets

```
docker exec -it pg bash
/# cat /var/lib/postgresql/data/postgresql.conf | grep timezone
log_timezone = 'Asia/Tokyo'                                                    
timezone = 'Asia/Tokyo'                                                        
#timezone_abbreviations = 'Default'     # Select the set of available time zone
# share/timezonesets/.
```

```
PGPASSWORD=pass psql -h localhost -p 5432 -U postgres
PGTZ=UTC PGPASSWORD=pass psql -h localhost -p 5432 -U postgres

select current_setting('timezone');

set timezone to 'UTC';
set timezone to 'Asia/Tokyo';
```

```
docker-compose down --rmi all --volumes
```