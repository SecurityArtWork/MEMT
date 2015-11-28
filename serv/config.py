# -*- coding: utf-8 -*-
"""Configuration module for the Faask application
"""
from __future__ import absolute_import, print_function
# pylint: disable=R0903,W0232,W0221,C0103,C0112,C0111

import os
from kombu import Queue, Exchange

BASEDIR = os.path.abspath(os.path.dirname(__file__))

class Config(object):
    """Global configuration
    """
    SECRET_KEY = os.environ.get('SECRET_KEY') or '{}'.format(os.urandom(69))
    WTF_CSRF_ENABLED = True
    SSL_DISABLE = True

    SQLALCHEMY_COMMIT_ON_TEARDOWN = True
    SQLALCHEMY_RECORD_QUERIES = True

    CELERY_BROKER_URL = 'amqp://localhost/'
    CELERY_RESULT_BACKEND = "mongodb://localhost:27017/memtJobs"
    CELERY_TASK_RESULT_EXPIRES = 18000  # 5 hours.
    CELERY_DEFAULT_QUEUE = 'memtQueue'
    CELERY_QUEUES = (
        Queue('memtQueue', routing_key='memt.#'),
        Queue('feedTasks', routing_key='feed.#'),
    )
    CELERY_MONGODB_BACKEND_SETTINGS = {
        "taskmeta_collection": "memtResults",
    }
    CELERY_TASK_SERIALIZER = 'json'
    CELERY_RESULT_SERIALIZER = 'json'
    CELERY_ACCEPT_CONTENT = ['json']
    CELERY_TIMEZONE = 'Europe/Madrid'
    CELERY_ENABLE_UTC = True
    CELERY_CREATE_MISSING_QUEUES = True

    MONGO_HOST = 'localhost'
    MONGO_PORT = '27017'

    MAXMAIN_DB_COUNTRIES = "/opt/dbs/GeoLite2-Country.mmdb"
    MAXMAIN_DB_CITIES = "/opt/dbs/GeoLite2-City.mmdb"

    RT_LAST_COUNTRIES = 1000
    FEED_LAST_NEWS = 5

    TMP_UPLOAD_FOLDER = os.path.join(BASEDIR, "..", "aux", "uploads")
    BIN_UPLOAD_FOLDER = os.path.join(BASEDIR, "..", "aux", "malware", "artifacts")
    IMG_UPLOAD_FOLDER = os.path.join(BASEDIR, "..", "aux", "malware", "images")

    MEMT_SLOW_DB_QUERY_TIME = 0.5

    LANGUAGES = {
        'en': 'English',
        'es': 'Español'
    }

    @staticmethod
    def init_app(app):
        pass


class DevelopmentConfig(Config):
    DEBUG = True
    SQLALCHEMY_DATABASE_URI = os.environ.get('DEV_DATABASE_URL') or \
        'sqlite:///' + os.path.join(BASEDIR, 'data-dev.sqlite')
    MONGO_DBNAME = 'memt'
    WTF_CSRF_ENABLED = False


class TestingConfig(Config):
    TESTING = True
    SQLALCHEMY_DATABASE_URI = os.environ.get('TEST_DATABASE_URL') or \
        'sqlite:///' + os.path.join(BASEDIR, 'data-test.sqlite')
    MONGO_DBNAME = 'memt'
    WTF_CSRF_ENABLED = False


class ProductionConfig(Config):
    SQLALCHEMY_DATABASE_URI = os.environ.get('DATABASE_URL') or \
        'sqlite:///' + os.path.join(BASEDIR, 'data.sqlite')
    MONGO_DBNAME = 'memt'

    @classmethod
    def init_app(cls, app):
        pass


class UnixConfig(ProductionConfig):
    """
    """
    @classmethod
    def init_app(cls, app):
        """
        """
        ProductionConfig.init_app(app)

        # log to syslog
        import logging
        from logging.handlers import SysLogHandler
        syslog_handler = SysLogHandler()
        syslog_handler.setLevel(logging.WARNING)
        app.logger.addHandler(syslog_handler)


config = {
    'development': DevelopmentConfig,
    'testing': TestingConfig,
    'production': ProductionConfig,
    'unix': UnixConfig,
    'default': DevelopmentConfig
}
