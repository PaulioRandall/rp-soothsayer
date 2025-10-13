import Entity from './Entity'

export default class Code extends Entity {
	static fromObject(obj) {
		const t = Entity.populateFromObject(new Code(), obj)

		t._description = obj.description

		return t
	}

	constructor() {
		super()

		this._description = ''
	}

	get description() {
		return this._description
	}

	toObject() {
		const obj = super.toObject()

		obj.description = this._description

		return obj
	}
}

