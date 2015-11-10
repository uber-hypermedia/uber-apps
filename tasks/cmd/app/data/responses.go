package data

var (
	Emptylist = `
	{ 
		"uber": 
		{ 
			"version": 1.0, 
			"data": 
			[
				{ 
					"id": "links", 
					"data": 
					[ 
						{ 
							"id": "list", 
							"rel": [ "collection" ], 
							"name": "links",
							"url": "/tasks/", 
							"action": "read" 
						},
						{ 
							"id": "search", 
							"rel": [ "search" ], 
							"name": "links",
							"url": "/tasks/search", 
							"action": "read",
							"model": "?text={text}",
						} 
					] 
				},
				{
					"id": "tasks",
					"data": []
				}
			] 
		}
	}`
	Singletask = `
	{ 
		"uber": 
		{ 
			"version": 1.0, 
			"data": 
			[
				{ 
					"id": "links", 
					"data": 
					[ 
						{ 
							"id": "list", 
							"rel": [ "collection" ], 
							"name": "links",
							"url": "/tasks/", 
							"action": "read" 
						},
						{ 
							"id": "search", 
							"rel": [ "search" ], 
							"name": "links",
							"url": "/tasks/search", 
							"action": "read",
							"model": "?text={text}",
						} 
					] 
				},
				{
					"id": "tasks",
					"data": 
					[
						{
							"id": "task1",
							"rel": [ "item" ],
							"name": "tasks",
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "model": "id={id}", "action": "append"},
								{ "name": "text", "value": "this is a task" },
							]
						}
					]
				}
			]
		}
	}`
	Multipletasks = `
	{ 
		"uber": 
		{ 
			"version": 1.0, 
			"data": 
			[
				{ 
					"id": "links", 
					"data": 
					[ 
						{ 
							"id": "list", 
							"rel": [ "collection" ], 
							"name": "links",
							"url": "/tasks/", 
							"action": "read" 
						},
						{ 
							"id": "search", 
							"rel": [ "search" ], 
							"name": "links",
							"url": "/tasks/search", 
							"action": "read",
							"model": "?text={text}",
						} 
					] 
				},
				{
					"id": "tasks",
					"data": 
					[
						{
							"id": "task1",
							"rel": [ "item" ],
							"name": "tasks",
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "model": "id={id}", "action": "append"},
								{ "name": "text", "value": "this is task one" },
							]
						},
						{
							"id": "task2",
							"rel": [ "item" ],
							"name": "tasks",
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "model": "id={id}", "action": "append"},
								{ "name": "text", "value": "this is task two" },
							]
						},
 						{
							"id": "task3",
							"rel": [ "item" ],
							"name": "tasks",
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "model": "id={id}", "action": "append"},
								{ "name": "text", "value": "this is task three" },
							]
						}
					]
				}
			]
		}
	}`
)
