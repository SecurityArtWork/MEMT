# -*- coding: utf-8 -*-
"""Module that controls the FAQ view.
"""
from __future__ import absolute_import

from flask import Blueprint, render_template


bp = Blueprint('faq', __name__, url_prefix='/faq')


@bp.route('', methods=['GET'])
def index():
    return render_template('faq/index.html')
