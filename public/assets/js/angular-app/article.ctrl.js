app.controller('articleCtrl', ['$scope', '$log', 'Article', function ($scope, $log, Article) {

  $scope.newArticle = new Article({});
  $scope.articles = Article.query();

  $scope.createArticle = function () {
    $scope.newArticle.$save().then(function () {
      $scope.newArticle = new Article({});
      $scope.articles = Article.query();
    });
  };

  $scope.deleteArticle = function (article) {
  	$log.log(article);
  	article.$delete(function () {
  		$log.log('deleted');
  	});
  };

}]);