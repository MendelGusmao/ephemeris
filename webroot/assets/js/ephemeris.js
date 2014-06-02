'use strict';

angular.module('ephemeris', ['ngResource', 'ui.router', 'ui.bootstrap'])
  .config(['$locationProvider', function($locationProvider) {
    $locationProvider.hashPrefix('!');
  }])
  .config(['$httpProvider', function($httpProvider) {
    $httpProvider.interceptors.push('ErrorInterceptor');
  }])
  .value('APIRoot', '/api');
