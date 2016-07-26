app.controller('articleCtrl', ['$scope', '$log', 'Article', function ($scope, $log, Article) {

  $scope.data = {};
  $scope.data.newArticle = new Article({});
  $scope.data.articles = Article.query();

  $scope.createArticle = function () {
    $scope.data.newArticle.$save(function (article) {
      $log.log(article);
      $scope.data.newArticle = new Article({});
      $scope.data.articles.push(article);
    });
  };

  $scope.editArticle = function (article, $index) {
    article.title = "I'm different now.";
    article.$update({id: article.id}, function (data) {
      $log.log('udpated', data);
      $scope.data.articles.splice($index, 1); // remove old
      $scope.data.articles.splice($index, 0, data); // in new
    });
  };

  $scope.deleteArticle = function (article, $index) {
  	article.$remove({id: article.id}, function (data) {
      $log.log('deleted article', data);
      $scope.data.articles.splice($index, 1);
    }, function (err) {
      $log.log(err);
    });
  };

}]);