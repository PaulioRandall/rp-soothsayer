export default class Study {
	constructor() {
		this._id = ''
		this._name = ''

		this._media = []
		this._sessions = []
		this._observationCodes = []
	}

	set id(v) {
		checkId(v)
		this._id = v
	}

	get id() {
		return this._id
	}

	set name(v) {
		checkName(v)
		this._name = v
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

	fromObject(obj) {
		if (obj.id) {
			this.id = obj.id
		}

		if (obj.name) {
			this.name = obj.name
		}

		if (obj.media) {
			this.addMediaFromObject(obj.media)
		}

		if (obj.sessions) {
			this.addSessionsFromObject(obj.sessions)
		}

		if (obj.observationCodes) {
			this.addObservationCodesFromObject(obj.observationCodes)
		}
	}

	toObject(obj) {
		// TODO
	}

	addMediaFromObject(obj) {
		// TODO: could be single media object or array of media objects
	}

	addSessionsFromObject(obj) {
		// TODO: could be single media object or array of media objects
	}

	addObservationCodesFromObject(obj) {
		// TODO: could be single media object or array of media objects
	}
}

function checkId(id) {
	if (typeof id !== 'string') {
		throw new Error('Study ID must be a non-empty string')
	}

	if (id.trim() === '') {
		throw new Error('Study ID must be a non-empty string')
	}
}

function checkName(name) {
	if (typeof name !== 'string') {
		throw new Error('Study name must be a non-empty string')
	}

	if (name.trim() === '') {
		throw new Error('Study name must be a non-empty string')
	}
}
