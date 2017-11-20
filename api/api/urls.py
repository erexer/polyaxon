# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

from django.conf import settings
from django.conf.urls import include, url
from django.contrib import admin

from rest_framework.authtoken.views import obtain_auth_token

from api import views


API_V1 = 'api/v1'

urlpatterns = [
    url(r'^users/', include('users.urls', namespace='users')),
    url(r'^admin/', include(admin.site.urls)),
    url(r'^api/token/', obtain_auth_token),
    url(r'^{}/'.format(API_V1), include('clusters.urls', namespace='v1')),
    url(r'^{}/'.format(API_V1), include('experiments.urls', namespace='v1')),
    url(r'^{}/'.format(API_V1), include('projects.urls', namespace='v1')),
    url(r'^_health/?$', views.HealthView.as_view(), name='health_check'),
    url(r'^$', views.IndexView.as_view(), name='index'),
]

if settings.DEBUG:
    import debug_toolbar

    urlpatterns += [
        url(r'^__debug__/', include(debug_toolbar.urls)),
    ]
