# -*- coding: utf-8 -*-
"""Extensions module. Each extension is initialized in the application factory pattern
"""
from __future__ import print_function, absolute_import
# pylint: disable=F0401,C0103

import os
from celery import Celery

from config import Config

from flask.ext.babel import Babel
from flask.ext.pymongo import PyMongo
from flask_socketio import SocketIO

from kombu.serialization import register
from kombu import serialization

from .utils import memt_dumps, memt_loads

# Flask extensions
babel = Babel()
mongo = PyMongo()

socketio = SocketIO()

# Celery config
celery = Celery(__name__,
                broker=Config.CELERY_BROKER_URL,
                backend=Config.CELERY_RESULT_BACKEND)

register('memtjson', memt_dumps, memt_loads,
    content_type='application/x-memtjson',
    content_encoding='utf-8')
serialization.registry._decoders.pop("application/x-python-serialize")


if not os.environ.get('PRODUCTION'):
    from flask_debugtoolbar import DebugToolbarExtension
    toolbar = DebugToolbarExtension()

