export default class MediaSection {
	static fromObject() {}

	constructor(id, name) {
		this._id = id
		this._name = name

		this._media = []
		this._sessions = []
		this._observationCodes = []
	}

	get id() {
		return this._id
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

	get observationCodes() {
		return this._observationCodes
	}

	toObject(obj) {}
}
