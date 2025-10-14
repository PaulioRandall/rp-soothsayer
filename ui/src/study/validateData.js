
export function validateData(data) {
	const errors = []

	// TODO

	return errors
}

export function validateTagObject(obj) {
	validateStringField('Tag', 'id', obj.id)
	validateStringField('Tag', 'name', obj.name)

	if (obj.description) {
		validateStringField('Tag', 'description', obj.description)
	}
}

export function validateStringField(objType, fieldName, value) {
	if (typeof value !== 'string') {
		throw new Error(`${objType} ${fieldName} must be a string`)
	}

	if (value.trim() === '') {
		throw new Error(`${objType} ${fieldName} must be a non-empty`)
	}
}

export default validateData
