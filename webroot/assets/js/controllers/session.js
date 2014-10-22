'use strict';
angular.module('ephemeris')
  .controller('SessionCtrl', ['$scope', '$stateParams', '$location', 'Session', function($scope, $stateParams, $location, Session) {
    $scope.check = function() {
      Session.get({}, function() {
        console.log(arguments);
      });
    }

    $scope.signin = function(credentials) {
      new Session(credentials).$save(function(response, headers) {
        console.log(response, headers);
      });
    };

    $scope.destroy = function() {
      new Session($scope).$delete(function() {
        console.log(arguments);
      });
    };
  }]);
