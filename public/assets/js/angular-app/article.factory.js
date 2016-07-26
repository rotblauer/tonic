app.factory('Article', ['$resource', function ($resource) {
	return $resource('/v1/a/:id', { id: '@id' },
	    {
        'query': {
		      method: 'GET',
		      isArray: true,
		      transformResponse: function(res) {
		        return angular.fromJson(res).data;
		      }
		    },
        'update': { method:'PUT' }
	    });
}]);

