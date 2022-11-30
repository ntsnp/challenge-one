# Submission for NTSNP Challenge One

I have used MongoDB as database so, before running the project, you need to add mongodb connection string to `.env` and `frontend/.env.local` files.

```
MONOGODB_URI=<your_mongodb_connection_string>
```

- Scrape data and add to database and add to a folder
```
./get-sentry-blogs folderName
```
> works on linux machine, for windows add .py extension to the file and run it
> If you face permission error use `sudo chmod +x get-sentry-blogs`

- Frontend
```
cd frontend
npm install
npm run dev
```

- Cron Jobs
```
crontab -e
```

This will open a file, add the following line to it and the end.
```
* */2 * * * /home/$(USER)/projectashik/get-sentry-blogs
```