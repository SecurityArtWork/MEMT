# -*- coding: utf-8 -*-
"""Graph API module.
This module define every API related to maps.
"""
from __future__ import print_function, absolute_import

from . import __api_prefix__

from flask import jsonify

from flask.ext.classy import FlaskView, route



class GraphView(FlaskView):
    route_prefix = __api_prefix__

    @route('/rt/<int:qty>')
    def rt(self, qty=100):
        pass

    @route('/geo/<sha256:hash>')
    def geo(self, hash):
        pass

    @route('/spread/<sha256:hash>')
    def spread(self, hash):
        pass
