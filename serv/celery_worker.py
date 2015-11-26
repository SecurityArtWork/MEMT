# -*- coding: utf-8 -*-
"""This module sets the Celery into the Flask application context.
"""
from __future__ import print_function, absolute_import

import os
from app import celery, create_app

app = create_app(os.getenv('FLASK_CONFIG') or 'default')
app.app_context().push()
