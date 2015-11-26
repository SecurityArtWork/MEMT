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


# Flask extensions
babel = Babel()
mongo_assets = PyMongo()
socketio = SocketIO()

# Celery config
celery = Celery(__name__,
                broker=Config.CELERY_BROKER_URL,
                backend=Config.CELERY_RESULT_BACKEND)

if not os.environ.get('PRODUCTION'):
    from flask_debugtoolbar import DebugToolbarExtension
    toolbar = DebugToolbarExtension()
