
export default class Study {
	constructor(data) {
		this._data = data /* = {
			name: "",
			filepath: "",
			media: [
				{
					id: "",
					name: "",
					tagIds: [""],
					filepath: "",
					sections: [
						{
							id: "",
							name: "",
							tagIds: [""],
							startTime: 0, // in seconds
							endTime: 0, // in seconds
						}
					]
				}
			],
			sessions: [
				{
					id: "",
					name: "",
					tagIds: [""],
					observation: [
						{
							id: "",
							name: "",
							tagIds: [""],
							startTime: 0, // in seconds
							endTime: -1, // in seconds
							quote: "",
							description: "",
						}
					]
				}
			],
			codes: [
				{
					id: "",
					name: "",
					tagIds: [""],
					description: "",
				}
			],
			tags: [
				{
					id: "",
					name: "",
					description: "",
				}
			]
		} */
	}

	get data() {
		return structuredClone(this._data)
	}

	get name() {
		return this._data.name
	}

	get filepath() {
		return this._data.filepath
	}

	get media() {
		return this.toObject().media
	}

	get sessions() {
		return this.toObject().sessions
	}

	get codes() {
		return this.toObject().codes
	}

	get tags() {
		return this.toObject().tags
	}

	// Creates a new tag.
	newTag(name, description="") {
		// TODO
	}

	// Updates an existing tag.
	updateTag(id, { name=null, description=null }) {
		// TODO
	}

	// Deletes an existing tag.
	deleteTag(id) {
		// TODO
	}

	// Searches all data to find the thing (entity, object, entry) with the
	// specified ID, else returns null.
	findThingById(id) {
		// TODO

		return null
	}

	findThingsWithTag(tag) {
		// TODO
	}

	findThingsWithTagId(tagId) {
		// TODO
	}
}

