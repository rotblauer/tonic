app.factory('Article', ['$resource', function ($resource) {
	return $resource('/v1/a/:id', {id: '@_id'}, //{ 'id': 'id' }
	    {
        'query': {
		      method: 'GET',
		      isArray: true,
		      // transformResponse: function(res) {
		      //   return angular.fromJson(res).data;
		      // }
		    },
        'update': { 
        	method:'PUT'
        // 	params: {
        // 		id: 'id'
        // 	}
        }
	    });
	//     
	// var url = '/v1/a';
	// return {
	// 	create: function (article) {
	
	// 	},
	// 	update: function (article) {
	
	// 	},
	// 	delete: function (article) {
	
	// 	},
	// 	get: function (article) {
	
	// 	},
	// 	index: function () {
	// 		return $http.get(url).then(function (res) {
	// 			return res;
	// 		}).catch(function (err) { return err; });
	// 	}
	// };
}]);

