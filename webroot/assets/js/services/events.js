'use strict';

angular.module('ephemeris')
  .factory('Events', ['APIRoot', 'ErrorInterceptor', '$resource', function(APIRoot, ErrorInterceptor, $resource) {
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
