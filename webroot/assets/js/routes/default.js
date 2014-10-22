'use strict';

angular.module('ephemeris')
  .config(['$urlRouterProvider', function($urlRouterProvider) {
    $urlRouterProvider.otherwise('/events');
  }]);
