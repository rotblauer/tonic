app.controller('articleCtrl', ['$scope', '$log', 'Article', function ($scope, $log, Article) {

  $scope.newArticle = new Article({});
  $scope.articles = Article.query();

  $scope.createArticle = function () {
    $scope.newArticle.$save().then(function () {
      $scope.newArticle = new Article({});
      $scope.articles = Article.query();
    });
  };

  $scope.editArticle = function (article) {
    article.title = "I'm different now.";
    article.$update({id: article.id}, function (data) {
      $log.log('udpated', data);
    });
  };

  $scope.deleteArticle = function (article) {
  	article.$remove(function (data) {
      $log.log('deleted article', data);
    }, function (err) {
      $log.log(err);
    });
  };

}]);