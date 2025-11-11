package validate

var studySchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"name": JsonObject{
			"type": "string",
		},
		"filepath": JsonObject{
			"type": "string",
		},
		"media": JsonObject{
			"type":  "array",
			"items": mediaSchema,
		},
		"codes": JsonObject{
			"type":  "array",
			"items": codeSchema,
		},
		"tags": JsonObject{
			"type":  "array",
			"items": tagSchema,
		},
	},
}

var mediaSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
			"type": "string",
		},
		"name": JsonObject{
			"type": "string",
		},
		"tags": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
		"filepath": JsonObject{
			"type": "string",
		},
		"sections": JsonObject{
			"type":  "array",
			"items": mediaSectionSchema,
		},
		"observations": JsonObject{
			"type":  "array",
			"items": observationSchema,
		},
	},
}

var mediaSectionSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
			"type": "string",
		},
		"name": JsonObject{
			"type": "string",
		},
		"tags": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
		"start": JsonObject{
			"type": "number",
		},
		"end": JsonObject{
			"type": "number",
		},
	},
}

var observationSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
			"type": "string",
		},
		"tags": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
		"start": JsonObject{
			"type": "number",
		},
		"quote": JsonObject{
			"type": "string",
		},
		"description": JsonObject{
			"type": "string",
		},
		"codes": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
	},
}

var codeSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
			"type": "string",
		},
		"name": JsonObject{
			"type": "string",
		},
		"tags": JsonObject{
			"type": "array",
			"items": JsonObject{
				"type": "string",
			},
		},
		"description": JsonObject{
			"type": "string",
		},
	},
}

var tagSchema = JsonObject{
	"type": "object",
	"fields": SchemaPropFields{
		"id": JsonObject{
			"type": "string",
		},
		"name": JsonObject{
			"type": "string",
		},
		"description": JsonObject{
			"type": "string",
		},
	},
}
