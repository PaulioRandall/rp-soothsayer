import Entity from './Entity'

export default class PointOfInterest extends Entity {
	static fromObject(obj) {
		const poi = Entity.populateFromObject(new PointOfInterest(), obj)
	
		// TODO

		return poi
	}

	constructor() {
		super()
	}

	toObject() {
		const obj = super.toObject()

		// TODO

		return obj
	}
}
