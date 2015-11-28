```
 __  __ _____ __  __ _____
|  \/  | ____|  \/  |_   _|
| |\/| |  _| | |\/| | | |
| |  | | |___| |  | | | |
|_|  |_|_____|_|  |_| |_|
```

# MEMT (Massive Early Malware Triage)

## What is MEMT?
MEMT is not a simple tool or command line interface, MEMT is a complete whole platform which get focused on Bid Data. The main idea behind MEMT is to catch new malware that other tools are not able to identify, like anti-virus, for that purposes MEMT bet for technologies like MongoDB, Celery, Go and Python. Technologies open source that helps MEMT to categorize malware using Big Data techniques and algorithms.

MEMT is able to identify malware, it has a great dashboard which displays in real time from where is the malware been identified around the globe, moreover the platform offers a really good detail of each malware, for example some static analysis is shown as well as a picture of the malware, yes you have read well, the picture of the malware. That picture helps to identify in one sight the binary sections and how it is spited up internally, actually it is a really nice feature.

## Requirements
There several technologies that the MEMT project uses. They are as follows:

| Back-end | Front-end  |
| :------: | :--------: |
|    Go    | JavaScript |
|  Python  |   jQuery   |
|  Flask   | Bootstrap  |
| MongoDB  |  SocketIO  |
|  Celery  |            |
| RabbitMQ |            |
| SocketIO |            |
| -------- | ---------- |

## Components
As said before MENT is a big platform so MEMT is built upon different modules or applications, following you have the main parts.

- Web server: It is build using Python and Flask as a web server, moreover it needs to be run by Gunicorn which is a WSGI compatible and allows good performance. On the other hand, there is a Celery and RabbitMQ. RabbitMQ is like a buffer database and it helps Celery to process all the background tasks.

- Database: We use MongoDB as database. It is well known how good performance and capacity has this database. MongoDB has such amount of functionalities that makes that make the development really easy.

- Cli: It is a command line interface that is built using Go. The Cli will help the users of MEMT to load malware in a easy way. It can run in two modes, daemon and standalone.

- Categorizer: This tool helps to whom wants to use the platform for their own purposes to create the first set of data from a set of malware. This initial set is needed to populate the database and be able to interact with the front-end correctly.

- Analyzer: It will perform the background task that has been sent by Celery. This tool analyze the binary that has been uploaded into the platform and find for patters among different parameters. Finally it saves update the whole system to be more accurate next time.

## Installation
You can find the installation files (README.md) inside root each component.

- [Web server](srev/README.md)
- [Cli](cli/README.md)
- [Categorizer](cat/README.md)
- [Analyzer](anal/README.md)

