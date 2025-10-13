import { validateTagObject } from './validateObject'
import Entity from './Entity'

export default class Tag extends Entity {
	static fromObject(obj) {
		validateTagObject(obj)

		if (obj.description) {
			this.description = obj.description
		}
	}

	constructor() {
		super()
		this._description = ''
	}

	set description(v) {
		this._description = v
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
