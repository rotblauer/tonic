'use strict';

var app = angular.module('main', [
  'ui.router',
  // 'ng-token-auth',
  'angular-jwt',
  'ngResource'
  ]);

app.config(function($stateProvider, $urlRouterProvider, $httpProvider, jwtInterceptorProvider, $resourceProvider) {

  // $resourceProvider.defaults.stripTrailingSlashes = false;
  
  jwtInterceptorProvider.tokenGetter = [function() {
      return localStorage.getItem('token');
    }];

  $httpProvider.interceptors.push('jwtInterceptor');

  // Routing.
  // 
  $urlRouterProvider.otherwise("/articles");
  //
  // Now set up the states
  $stateProvider
    .state('articles', {
      url: "/articles",
      views: {
        'articles': {
          templateUrl: "/public/assets/templates/article/articles.html",
          controller: 'articleCtrl'
        },
        'registration': {
          templateUrl: "/public/assets/templates/registration.html",
          controller: 'registrationCtrl'
        },
        'createArticle': {
          templateUrl: "/public/assets/templates/article/create.html",
          controller: 'articleCtrl'
        }
      }
    })
    .state('articles.list', {
      url: "/list",
      templateUrl: "/public/assets/templates/articles.list.html",
      controller: "articleCtrl"
    });
  });











