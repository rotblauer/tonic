'use strict';

var app = angular.module('main', [
  'ui.router', 
  'ng-token-auth',
  'angular-jwt'
  ]);

app.config(function($authProvider) {

    // the following shows the default values. values passed to this method
    // will extend the defaults using angular.extend

    $authProvider.configure({
      apiUrl:                  '/v1',
      // tokenValidationPath:     '/auth/validate_token',
      // signOutUrl:              '/auth/sign_out',
      emailRegistrationPath:   '/auth/signup',
      // accountUpdatePath:       '/auth',
      // accountDeletePath:       '/auth',
      confirmationSuccessUrl:  window.location.href,
      // passwordResetPath:       '/auth/password',
      // passwordUpdatePath:      '/auth/password',
      // passwordResetSuccessUrl: window.location.href,
      emailSignInPath:         '/auth/signin',
      storage:                 'localStorage', // 'cookies'
      // forceValidateToken:      false,
      validateOnPageLoad:      false, //true, // FIXME
      proxyIf:                 function() { return false; },
      proxyUrl:                '/proxy',
      // omniauthWindowType:      'sameWindow',
      // authProviderPaths: {
      //   github:   '/auth/github',
      //   facebook: '/auth/facebook',
      //   google:   '/auth/google'
      // },
      tokenFormat: {
        // "access-token": "{{ token }}",
        // "token-type":   "Bearer",
        // "client":       "{{ clientId }}",
        // "expiry":       "{{ expiry }}",
        // "uid":          "{{ uid }}"
        "Authorization": "Bearer {{ token }}"
      },
      cookieOps: {
        path: "/",
        expires: 9999,
        expirationUnit: 'days',
        secure: false,
        domain: 'domain.com'
      },
      createPopup: function(url) {
        return window.open(url, '_blank', 'closebuttoncaption=Cancel');
      },
      parseExpiry: function(headers) {
        // convert from UTC ruby (seconds) to UTC js (milliseconds)
        // return (parseInt(headers['expiry']) * 1000) || null;
        return (Date.parse(headers['Authorization'].match(/expiry=([^ ]+) /)[1])) || null;
      },
      handleLoginResponse: function(response) {
        return response.data;
      },
      handleAccountUpdateResponse: function(response) {
        return response.data;
      },
      handleTokenValidationResponse: function(response) {
        return response.data;
      }
    });
  });


app.config(function($stateProvider, $urlRouterProvider) {
  //
  // For any unmatched url, redirect to /state1
  $urlRouterProvider.otherwise("/state1");
  //
  // Now set up the states
  $stateProvider
    .state('state1', {
      url: "/state1",
      templateUrl: "/public/assets/templates/state1.html"
    })
    .state('state1.list', {
      url: "/list",
      templateUrl: "/public/assets/templates/state1.list.html",
      controller: "testCtrl"
    });
    // .state('state2', {
    //   url: "/state2",
    //   templateUrl: "partials/state2.html"
    // })
    // .state('state2.list', {
    //   url: "/list",
    //   templateUrl: "partials/state2.list.html",
    //   controller: function($scope) {
    //     $scope.things = ["A", "Set", "Of", "Things"];
    //   }
    // });
});


app.controller('testCtrl', ['$scope', function ($scope) {
	$scope.testes = 'check check';
  $scope.items = ['why', 'hello', 'there'];
}]);