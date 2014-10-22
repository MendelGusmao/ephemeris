'use strict';

angular.module('ephemeris')
  .factory('Session', ['APIRoot', '$resource', function(APIRoot, $resource) {
    return $resource(APIRoot + '/session');
  }]);
