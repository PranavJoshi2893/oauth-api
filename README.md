## How to Start the Server

Follow these steps to set up and run the application.

---

## **Update the `.env` File**

Create a `.env` file in the project root (or copy it from `.env.example`) and fill in the required values:

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=mydatabase
DB_USER=myuser
DB_PASSWORD=mypassword
DB_SSLMODE=disable
SRV_PORT=8080
```

> Make sure all values are correct before starting the server or database.

---

## **Start the Database**

Run the provided **start database script**, which:

* loads environment variables
* starts the PostgreSQL container
* waits until PostgreSQL accepts connections

Example:

```bash
./start_db.sh
```

Wait until you see:

```
PostgreSQL is ready to accept connections.
```

---

## **Start the Server**

Once the database is running, start your server:

```bash
./start_server.sh
```

Your application should now be running on the port specified in `SRV_PORT` (default: `8080`).