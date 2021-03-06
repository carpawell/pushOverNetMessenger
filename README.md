# **pushOverNetMessenger**


## Run everything in Docker:
*App and PostgreSQL start in Docker containers:*
```
$ make port=[yourPortForApp] app=[yourAppId] user=[yourUserDeviceId] run
```

- `app`, `port` and `user` **required**

*After such demo-start everything can be deleted with:*
```
$ make clean
```


## Run app in Docker:
*Run app:* 
```
$ make DB_DSN=[...] \
       DBPort=[...] \
       port=[...] \
       app=[...] \
       user=[...] startMessenger
```

- All arguments **required**

*Theoretically it will work.*


## Build app:
1. Set `DB_DSN` env variable of your DB
2. Build app: ```$ make build```
3. Run app with required(**_app_, _user_**) arguments: 
```
$ ./dist/pushOverNetMessenger -app=[...] -user=[...] -port=[...]
```


## API
When App is running and passed `user` and `app` arguments are correct there are two options:
1. Send notification to `user` from `app`: `POST` request to `localhost:[port]/messages` with JSON:
```
{
    "message": "your message"
}
```
2. Get statistic of your notifications from certain period:
`GET` request to `localhost:[port]/messages/statistics` with URL parameter `from=MM-DD-YYYYTHH:MM:SSZ`.
Time format is **important**.
