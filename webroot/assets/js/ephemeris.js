'use strict';

angular.module('ephemeris', ['ngResource', 'ui.router', 'ui.bootstrap'])
  .config(['$locationProvider', function($locationProvider) {
    $locationProvider.hashPrefix('!');
  }]);
