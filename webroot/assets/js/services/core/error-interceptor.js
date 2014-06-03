'use strict';

angular.module('ephemeris')
  .factory('ErrorInterceptor', ['APIRoot', '$q', '$rootScope', function(APIRoot, $q, $rootScope) {
    return {
      responseError: function(rejection) {
        if (rejection.config.url.substr(0, APIRoot.length) == APIRoot) {
          var error = [{
            message: [rejection.status, 'for', rejection.config.url].join(' ')
          }];

          $rootScope.$broadcast("messages", rejection.data || error);
        }

        return $q.reject(rejection);
      }
    }
  }]);
