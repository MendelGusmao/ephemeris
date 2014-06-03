'use strict';

angular.module('ephemeris')
  .factory('Events', ['APIRoot', '$resource', function(APIRoot, $resource) {
    return $resource(APIRoot + '/events/:id',
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
