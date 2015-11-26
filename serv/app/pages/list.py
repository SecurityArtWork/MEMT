# -*- coding: utf-8 -*-
"""Module that controls the artifacts list view.
"""
from __future__ import absolute_import

from flask import Blueprint, render_template


bp = Blueprint('list', __name__, url_prefix='/list')


@bp.route('', methods=['GET'])
def index():
    return render_template('list/index.html')
