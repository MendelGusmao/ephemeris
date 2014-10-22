'use strict';

angular.module('ephemeris')
  .config(['$stateProvider', function($stateProvider) {
    $stateProvider
      .state('events-list', {
        url: '/events',
        templateUrl: '/assets/views/events/index.html'
      })
      .state('events-new', {
        url: '/events/new',
        templateUrl: '/assets/views/events/new.html'
      })
      .state('events-show', {
        url: '/events/:id',
        templateUrl: '/assets/views/events/show.html'
      })
      .state('events-edit', {
        url: '/events/:id/edit',
        templateUrl: '/assets/views/events/edit.html'
      });
  }]);
