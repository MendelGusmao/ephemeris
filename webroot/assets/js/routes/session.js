'use strict';

angular.module('ephemeris')
  .config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
    $stateProvider
      .state('session-signin', {
        url: '/session/signin',
        templateUrl: '/assets/views/session/signin.html'
      });
  }]);
