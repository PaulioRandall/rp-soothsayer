export default class Entity {
	constructor() {
		this._id = ''
		this._name = ''
	}

	set id(v) {
		this._id = v
	}

	get id() {
		return this._id
	}

	set name(v) {
		this._name = v
	}

	get name() {
		return this._name
	}

	toObject() {
		return {
			id: this._id,
			name: this._name,
		}
	}
}
