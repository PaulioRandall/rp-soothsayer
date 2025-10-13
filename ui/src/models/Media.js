import Entity from './Entity'
import MediaSection from './MediaSection'

export default class Media extends Entity {
	static fromObject(obj) {
		const media = Entity.populateFromObject(new Media(), obj)
	
		g_populateArrayFromObjectArray(media._mediaSections, obj.mediaSections, MediaSection.fromObject)
		
		return media
	}

	constructor() {
		super()

		this._mediaSections = []
	}

	toObject() {
		const obj = super.toObject()

		obj.mediaSections = g_toObjects(this._mediaSections)

		return obj
	}
}
