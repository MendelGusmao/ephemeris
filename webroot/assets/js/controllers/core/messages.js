'use strict';

angular.module('ephemeris')
  .controller('MessagesController', ['$scope', '$timeout', function($scope, $timeout) {
      $scope.$on('messages', function(scope, messages) {
        $scope.messages = messages;
        console.info("Received broadcast:", messages);

        $timeout(function() {
          $scope.messages = [];
        }, 3000);
      });
    }
  ]);
