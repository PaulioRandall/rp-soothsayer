export default class Entity {
	static populateFromObject(entity, obj) {
		entity._parentId = obj.parentId
		entity._id = obj.id
		entity._name = obj.name

		return entity
	}

	constructor() {
		this._parentId = ''
		this._id = ''
		this._name = ''
	}

	get parentId() {
		return this._parentId
	}

	get id() {
		return this._id
	}

	get name() {
		return this._name
	}

	toObject() {
		return {
			parentId: this._parentId,
			id: this._id,
			name: this._name,
		}
	}
}
