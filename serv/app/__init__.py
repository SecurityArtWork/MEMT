# -*- coding: utf-8 -*-
"""The main web server application.
"""
from __future__ import print_function, absolute_import
# pylint: disable=C0111,E1101

import os

from flask import Flask
from flask import render_template
from flask import request

from config import config, Config
from .extensions import babel
from .extensions import mongo_assets
from .extensions import socketio
from .extensions import celery

from .api.v0.graph import GraphView
from .api.v0.malware import MalwareView

from .pages.index import bp as index
from .pages.search import bp as search
from .pages.list import bp as listbp
from .pages.detail import bp as detail
from .pages.faq import bp as faq
from .pages.download import bp as download


def create_app(config_name):
    """This function creates a Factory application patters, according to the Flask web page.

    :param config_name: This is the configuration object that will be used to configure the application.
    """
    app = Flask(__name__)

    app.config.from_object(config[config_name])
    config[config_name].init_app(app)

    register_extensions(app)
    register_blueprints(app)

    register_socketio_tasks()

    register_errorhandlers(app)
    return app


def register_extensions(app):
    babel.init_app(app)
    configure_babel(app)
    mongo_assets.init_app(app, config_prefix='MONGO')

    # Setting socketIO in async mode
    socketio.init_app(app, async_mode='eventlet', engineio_logger=True)
    celery.conf.update(app.config)

    # Add development helper for the UI
    if not os.environ.get('PRODUCTION'):
        from .extensions import toolbar
        toolbar.init_app(app)

    return None


def register_blueprints(app):
    app.register_blueprint(index)
    app.register_blueprint(search)
    app.register_blueprint(listbp)
    app.register_blueprint(detail)
    app.register_blueprint(faq)
    app.register_blueprint(download)
    GraphView.register(app)
    MalwareView.register(app)
    return None


def register_errorhandlers(app):
    def render_error(error):
        # If a HTTPException, pull the `code` attribute; default to 500
        error_code = getattr(error, 'code', 500)
        return render_template("errors/{0}.html".format(error_code)),\
            error_code
    for errcode in [401, 404, 500]:
        app.errorhandler(errcode)(render_error)
    return None


def register_before_request(app):
    return None


def register_after_request(app):
    return None


def register_socketio_tasks():
    pass


def configure_babel(app):
    @babel.localeselector
    def get_locale():
        languages = app.config['LANGUAGES']
        return request.accept_languages.best_match(languages.keys())
