# -*- coding: utf-8 -*-
"""Module that controls the details view.
"""
from __future__ import absolute_import

from flask import Blueprint, render_template


bp = Blueprint('detail', __name__, url_prefix='/details')


@bp.route('', methods=['GET'])
def index():
    return render_template('detail/index.html')
