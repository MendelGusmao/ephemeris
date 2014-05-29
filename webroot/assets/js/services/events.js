'use strict';

angular.module('ephemeris')
  .factory('Events', ['$resource', function($resource) {
    return $resource('/api/events/:id',
      {
        id: '@id'
      },
      {
        all: {
          method: 'GET',
          isArray: true
        },
        update: {
          method: 'PUT'
        }
      }
    );
  }]);
