import Entity from './Entity'
import PointOfInterest from './PointOfInterest'

export default class Session extends Entity {
	static fromObject(obj) {
		const session = Entity.populateFromObject(new Session(), obj)
	
		g_populateArrayFromObjectArray(session._pointsOfInterest, obj.pointsOfInterest, PointOfInterest.fromObject)

		return session
	}

	constructor() {
		super()

		this._pointsOfInterest = []
	}

	toObject() {
		const obj = super.toObject()

		obj.pointsOfInterest = g_toObjects(this._pointsOfInterest)

		return obj
	}
}
