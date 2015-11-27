# -*- coding: utf-8 -*-
"""Module that controls the index view.
"""
from __future__ import absolute_import

from flask import Blueprint
from flask import render_template


from .common import get_common_info

bp = Blueprint('index', __name__)


@bp.route('/', methods=['GET'])
def index():
    info = get_common_info()
    return render_template('index/index.html',
                           totalAssets=info["total_assets"],
                           totalStrain=info["total_strains"],
                           lastNews=info["last_news"],
                           lastCountries=info["last_countries"])
