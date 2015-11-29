# MEMT Server

The server is composed by three main components, those components are [Gunicorn](http://gunicorn.org/), [Celery](http://www.celeryproject.org/), and [MongoDB](https://www.mongodb.org/).

- Gunicorn 'Green Unicorn' is a Python WSGI HTTP Server for UNIX. It's a pre-fork worker model ported from Ruby's Unicorn project. The Gunicorn server is broadly compatible with various web frameworks, simply implemented, light on server resources, and fairly speedy.

- Celery is an asynchronous task queue/job queue based on distributed message passing. It is focused on real-time operation, but supports scheduling as well.
The execution units, called tasks, are executed concurrently on a single or more worker servers using multiprocessing, Eventlet, or gevent. Tasks can execute asynchronously (in the background) or synchronously (wait until ready).

- MongoDB (from humongous) is a cross-platform document-oriented database. Classified as a NoSQL database, MongoDB eschews the traditional table-based relational database structure in favor of JSON-like documents with dynamic schema (MongoDB calls the format BSON), making the integration of data in certain types of applications easier and faster.

Those packages have some dependencies on other technologies, for example, Celery needs `amqp` database to store temporal messages, this is for example [RabbitMQ](https://www.rabbitmq.com/), Gunicorn can run as standalone, but it will increase its performance by adding a Nginx in front of Gunicorn.

Starting the server is not simple at first sight, because MongoDB has to be up and running before Celery and Gunicorn, and RabbitMQ (ampq database) needs to be up and running as well so the next instructions explains how you should start the project.

When working with Python you need to set up a virtual environment just to avoid messing up your python library folder in the system. When the environment is ready and you are bind to it, you should run the `pip install -r requirements.txt` and it will install all the dependencies needed for booting up the server. Take in mind that you have to install the correct requirements, since there are several inside requirements folder. Once you have finished with pip, you will run the following commands to start the server up correctly.

In our case you should run `pip install` as follows: From the root of the project type in a shell `pip install -r serv/requirements/prod.txt`, **be sure you are inside a virtual environment jail**.

The order to start all services is as follows, first the databases and then the Celery workers and the Gunicorn server.
- MongoDB, you will have to execute the follwing command line:
    ```sh
    mongod --dbpath <path to db folder>
    ```

- RabbitMQ, you will have to execute the follwing command line:
    ```sh
    rabbitmq-server
    ```

- Celery, you will have to be in the project folder, like `$VIRTUALENV/Web/app/` to be able to execute the following command line successfull:
    ```sh
    celery worker -A celery_worker.celery
    ```

- Gunicorn, you will have to be in the project folder, like $VIRTUALENV/Web/app/ to be able to execute the following command line successfull.
    ```sh
    gunicorn --worker-class eventlet --workers X --bind 127.0.0.1:5000 wsgi
    ```
    You have to adapt the `X` on `--workers` to match the formula `2*N+1`, where `N` is the number of cores in your CPU. Usually for development `X=1` is enough.
    Optionaly you can get the log messages adding these params
    ```sh
    --log-file=- --error-logfile=- --log-level=error
    ```

When you have everything running, you can acceses to some web sites to see what's going on. For example:

- To see what RabbitMQ is doing just visit [http://localhost:15672](http://localhost:15672)
- To see and test the webpage visit [http://localhost:5000/](http://localhost:5000/)
- To see what's on in MongoDB just run the mongo client, `mongo`.
