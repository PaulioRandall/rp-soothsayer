import Media from './Media'
import Session from './Session'
import Code from './Code'
import Tag from './Tag'

export default class Study {
	static fromObject(obj) {
		const study = new Study()

		study._name = obj.name
		g_populateArrayFromObjectArray(study._media, obj.media, Media.fromObject)
		g_populateArrayFromObjectArray(study._sessions, obj.sessions, Media.fromObject)
		g_populateArrayFromObjectArray(study._codes, obj.codes, Code.fromObject)
		g_populateArrayFromObjectArray(study._tags, obj.tags, Tag.fromObject)

		return study
	}

	constructor() {
		this._name = ""
		this._media = []
		this._sessions = []
		this._codes = []
		this._tags = []
	}

	get name() {
		return this._name
	}

	get media() {
		return this._media
	}

	get sessions() {
		return this._sessions
	}

	get codes() {
		return this._codes
	}

	get tags() {
		return this._tags
	}

	toObject() {
		return {
			name: this._name,
			media: g_toObjects(this._media),
			sessions: g_toObjects(this._sessions),
			codes: g_toObjects(this._codes),
			tags: g_toObjects(this._tags),
		}
	}
}

