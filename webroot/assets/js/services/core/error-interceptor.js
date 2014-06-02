'use strict';

angular.module('ephemeris')
  .factory('ErrorInterceptor', ['APIRoot', function(APIRoot) {
    return {
      responseError: function(response) {
        if (response.config.url.substr(0, APIRoot.length) != APIRoot) {
          return;
        }

        var errors = response.data || [];

        for (var i in errors) {
          var error = errors[i];
          console.error(
            '[ephemeris]',
            error.classification, '@', error.fieldNames.join(', '),
            '->', error.message
          );
        }
      }
    }
  }]);
