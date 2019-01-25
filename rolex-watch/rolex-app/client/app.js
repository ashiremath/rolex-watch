// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	
	$scope.querywatch = function(){

		var id = $scope.watch_id;

		appFactory.querywatch(id, function(data){
			$scope.query_watch = data;

			if ($scope.query_watch == "Could not locate watch"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordwatch = function(){

		appFactory.recordwatch($scope.watch, function(data){
			$scope.create_watch = data;
			$("#success_create").show();
		});
	}


});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

	factory.querywatch = function(id, callback){
    	$http.get('/get_watch/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordwatch = function(data, callback){

		var watch = data.id + "-" + data.name + "-" + data.qty + "-" + data.outlet + "-" + data.timestamp;

    	$http.get('/add_watch/'+watch).success(function(output){
			callback(output)
		});
	}

	return factory;
});


