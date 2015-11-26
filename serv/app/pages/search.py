# -*- coding: utf-8 -*-
"""Module that controls the search view.
"""
from __future__ import absolute_import

from flask import Blueprint, render_template


bp = Blueprint('search', __name__, url_prefix='/search')


@bp.route('', methods=['GET'])
def index():
    return render_template('search/index.html')
