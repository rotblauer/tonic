app.controller('registrationCtrl', ['$scope', '$log', '$http', 'jwtHelper', function ($scope, $log, $http, jwtHelper) {


  $scope.auth = function () {
    return window.localStorage.getItem('token');
  }

  $scope.tokenPayload = function () {
    var token = window.localStorage.getItem('token');
    return {
      payload: jwtHelper.decodeToken(token) || null, 
      expiry: jwtHelper.getTokenExpirationDate(token) || null,
      expired: (Date.parse(jwtHelper.getTokenExpirationDate(token)) >= Date.now() ? false : true) || null
    };
  };

  $scope.submitRegistration = function () {
    var form = {
      email: $scope.registrationForm.email, 
      password: $scope.registrationForm.password
    };
    $http.post('/v1/signin', form)
      .success(function (data, status, headers, config) {
        var token = headers('Authorization').match(/token=([^ ]+) /)[1];
        $log.log(data, status, headers, config);
        $log.log(headers('Authorization'));

        window.localStorage.setItem('token', token);
      })
      .error(function (data, status, headers, config) {
        $log.log(data, status, headers, config);
      });
  };

  $scope.logOut = function () {
    window.localStorage.setItem('token', '');
  };

}]);