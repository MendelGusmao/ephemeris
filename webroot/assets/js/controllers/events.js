'use strict';
angular.module('ephemeris')
  .controller('EventsCtrl', ['$scope', '$stateParams', '$location', 'Events', function($scope, $stateParams, $location, Events) {
    $scope.all = function() {
      $scope.events = Events.all();
    }

    $scope.one = function() {
      $scope.event = Events.get({ id: $stateParams.id });
    }

    $scope.create = function() {
      var event = new Events({
        id: this.id,
        name: this.name,
        place: this.place,
        description: this.description,
        URL: this.URL,
        logoURL: this.logoURL,
        beginning: this.beginning,
        end: this.end,
        registrationBeginning: this.registrationBeginning,
        registrationEnd: this.registrationEnd,
        visibility: this.visibility,
        status: this.status
      });

      event.$save(function(response, headers) {
        $location.path(headers('Location'));
      });
    };

    $scope.remove = function(event) {
      if (event) {
        event.$remove();

        for (var i in $scope.events) {
          if ($scope.events[i] === event) {
            $scope.events.splice(i, 1);
          }
        }
      } else {
        $scope.event.$remove(function() {
          $location.path('events');
        });
      }
    };

    $scope.update = function() {
      var event = $scope.event;

      event.$update(function(response, headers) {
        $location.path(headers('Location'));
      });
    };
  }]);
