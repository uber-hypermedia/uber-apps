package data

var (
	Emptylist = `
	{ 
		"uber": 
		{ 
			"version": "1.0", 
			"data": 
			[
				{ 
					"id": "links", 
					"data": 
					[ 
						{
							"id": "alps",
							"rel": [ "profile" ],
							"url": "/tasks-alps.xml",
							"action": "read"
						},
						{ 
							"id": "list", 
							"name": "links",
							"rel": [ "collection" ], 
							"url": "/tasks/", 
							"action": "read" 
						},
						{ 
							"id": "search", 
							"name": "links",
							"rel": [ "search" ], 
							"url": "/tasks/search", 
							"action": "read",
							"model": "?text={text}"
						},
						{ 
							"id": "add", 
							"name": "links",
							"rel": [ "add" ], 
							"url": "/tasks/", 
							"action": "append",
							"model": "text={text}"
						} 
					] 
				},
				{
					"id": "tasks"
				}
			] 
		}
	}`
	Singletask = `
	{ 
		"uber": 
		{ 
			"version": "1.0", 
			"data": 
			[
				{ 
					"id": "links", 
					"data": 
					[ 
						{
							"id": "alps",
							"rel": [ "profile" ],
							"url": "/tasks-alps.xml",
							"action": "read"
						},
						{ 
							"id": "list", 
							"name": "links",
							"rel": [ "collection" ], 
							"url": "/tasks/", 
							"action": "read" 
						},
						{ 
							"id": "search", 
							"name": "links",
							"rel": [ "search" ], 
							"url": "/tasks/search", 
							"action": "read",
							"model": "?text={text}"
						},
						{ 
							"id": "add", 
							"name": "links",
							"rel": [ "add" ], 
							"url": "/tasks/", 
							"action": "append",
							"model": "text={text}"
						} 
					] 
				},
				{
					"id": "tasks",
					"data": 
					[
						{
							"id": "task1",
							"name": "tasks",
							"rel": [ "item" ],
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "action": "append", "model": "id={id}"},
								{ "name": "text", "value": "task one" }
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
			"version": "1.0", 
			"data": 
			[
				{ 
					"id": "links", 
					"data": 
					[ 
						{
							"id": "alps",
							"rel": [ "profile" ],
							"url": "/tasks-alps.xml",
							"action": "read"
						},
						{ 
							"id": "list", 
							"name": "links",
							"rel": [ "collection" ], 
							"url": "/tasks/", 
							"action": "read" 
						},
						{ 
							"id": "search", 
							"name": "links",
							"rel": [ "search" ], 
							"url": "/tasks/search", 
							"action": "read",
							"model": "?text={text}"
						},
						{ 
							"id": "add", 
							"name": "links",
							"rel": [ "add" ], 
							"url": "/tasks/", 
							"action": "append",
							"model": "text={text}"
						} 					
					] 
				},
				{
					"id": "tasks",
					"data": 
					[
						{
							"id": "task1",
							"name": "tasks",
							"rel": [ "item" ],
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "action": "append", "model": "id={id}"},
								{ "name": "text", "value": "task one" }
							]
						},
						{
							"id": "task2",
							"name": "tasks",
							"rel": [ "item" ],
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "action": "append", "model": "id={id}"},
								{ "name": "text", "value": "task two" }
							]
						},
 						{
							"id": "task3",
							"name": "tasks",
							"rel": [ "item" ],
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "action": "append", "model": "id={id}"},
								{ "name": "text", "value": "task three" }
							]
						}
					]
				}
			]
		}
	}`
	Tasktwo = `
	{ 
		"uber": 
		{ 
			"version": "1.0", 
			"data": 
			[
				{ 
					"id": "links", 
					"data": 
					[ 
						{
							"id": "alps",
							"rel": [ "profile" ],
							"url": "/tasks-alps.xml",
							"action": "read"
						},
						{ 
							"id": "list", 
							"name": "links",
							"rel": [ "collection" ], 
							"url": "/tasks/", 
							"action": "read" 
						},
						{ 
							"id": "search", 
							"name": "links",
							"rel": [ "search" ], 
							"url": "/tasks/search", 
							"action": "read",
							"model": "?text={text}"
						},
						{ 
							"id": "add", 
							"name": "links",
							"rel": [ "add" ], 
							"url": "/tasks/", 
							"action": "append",
							"model": "text={text}"
						} 
					] 
				},
				{
					"id": "tasks",
					"data": 
					[
						{
							"id": "task1",
							"name": "tasks",
							"rel": [ "item" ],
							"data": 
							[
								{ "rel": [ "complete" ], "url": "/tasks/complete/", "action": "append", "model": "id={id}"},
								{ "name": "text", "value": "task two" }
							]
						}
					]
				}
			]
		}
	}`
)
